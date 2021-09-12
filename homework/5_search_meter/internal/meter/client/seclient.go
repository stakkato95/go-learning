package client

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

//response of seClient after metering is finished
type SearchEngineResponse struct {
	StatusCode  int
	RequestTime int64
	Error       error
	Name        string
	Ordinal     int
}

type SEClient interface {
	Do(ctx context.Context, wgStart *sync.WaitGroup, wgWait *sync.WaitGroup, wgAllDone *sync.WaitGroup,
		e SearchEngine, i int, result chan<- SearchEngineResponse)
	DoAll(ctx context.Context, wg *sync.WaitGroup, e SearchEngine, i int, result chan<- SearchEngineResponse)
}

//private implementation of SEClient interface
type seClient struct {
	client http.Client
}

//since seClient is private, return interface
func NewClient() SEClient {
	//create transport and http client
	t := http.Transport{
		DialContext: (&net.Dialer{
			//Timeout and KeepAlive are equal, so no keep-alive messages
			//will be sent during connection
			Timeout:   20 * time.Second,
			KeepAlive: 20 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	c := http.Client{
		Timeout:   10 * time.Second,
		Transport: &t,
	}

	return &seClient{client: c}
}

//do metering for a single search engine
func (s *seClient) Do(ctx context.Context, wgStart *sync.WaitGroup, wgWait *sync.WaitGroup, wgAllDone *sync.WaitGroup,
	e SearchEngine, i int, result chan<- SearchEngineResponse) {
	//at the end indicate caller that it can close 'result' chan, if all other goroutines also called wgAllDone.Done()
	defer func() {
		wgAllDone.Done()
	}()

	//create request object
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.SearchUrl, nil)
	if err != nil {
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		return
	}

	wgStart.Done() //indicate that a goroutine is ready to start processing
	wgWait.Wait()  //wait till caller is ready to receive data on 'result' chan

	start := time.Now()
	resp, err := s.client.Do(req)
	requestTime := time.Now().UnixMilli() - start.UnixMilli() //calculate time of http request

	if err != nil {
		//if there is no error, don't post anything to 'result' chan
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		return
	}
	//if there is no error don't forget to close body (prevent connection leakage)
	defer resp.Body.Close()

	select {
	case <-ctx.Done(): //if there was a faster search engine / timeout has elapsed, finish this function
		log.Debug().Msgf("context cancelled for '%s' cancelled", e.Name)
		return
	case result <- SearchEngineResponse{
		StatusCode:  resp.StatusCode,
		RequestTime: requestTime,
		Name:        e.Name,
		Ordinal:     i}: //if metering was finished successfully try to write to 'result' chan
	}
}

func (s *seClient) DoAll(ctx context.Context, wg *sync.WaitGroup, e SearchEngine, i int, result chan<- SearchEngineResponse) {
	//at the end indicate caller that it can close 'result' chan, if all other goroutines also called wgAllDone.Done()
	defer func() {
		wg.Done()
	}()

	//create request object
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.SearchUrl, nil)
	if err != nil {
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		return
	}

	start := time.Now()
	resp, err := s.client.Do(req)
	requestTime := time.Now().UnixMilli() - start.UnixMilli() //calculate time of http request

	if err != nil {
		//if there is an error write it anyway to 'result' chan
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		result <- SearchEngineResponse{Name: e.Name, Ordinal: i, Error: err}
		return
	}
	//if there is no error don't forget to close body (prevent connection leakage)
	defer resp.Body.Close()

	//if there is a result from search engine, write it to 'result' chan
	result <- SearchEngineResponse{StatusCode: resp.StatusCode, RequestTime: requestTime, Name: e.Name, Ordinal: i}
}
