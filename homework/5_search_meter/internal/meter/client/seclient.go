package client

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type SearchEngineResponse struct {
	StatusCode  int
	RequestTime int64
	Error       error
	Name        string
	Ordinal     int
}

type SEClient interface {
	Do(ctx context.Context, wgStart *sync.WaitGroup, wgWait *sync.WaitGroup, e SearchEngine, i int, result chan<- SearchEngineResponse)
	DoAll(ctx context.Context, wg *sync.WaitGroup, e SearchEngine, i int, result chan<- SearchEngineResponse)
}

type seClient struct {
	client http.Client
}

func NewClient() SEClient {
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

func (s *seClient) Do(ctx context.Context, wgStart *sync.WaitGroup, wgWait *sync.WaitGroup, e SearchEngine, i int, result chan<- SearchEngineResponse) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.SearchUrl, nil)
	if err != nil {
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		return
	}

	wgStart.Done()
	wgWait.Wait()

	start := time.Now()
	resp, err := s.client.Do(req)
	requestTime := time.Now().UnixMilli() - start.UnixMilli()

	if err != nil {
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		return
	}
	defer resp.Body.Close()

	select {
	case <-ctx.Done():
		log.Debug().Msgf("context cancelled for '%s' cancelled", e.Name)
		return
	case result <- SearchEngineResponse{
		StatusCode:  resp.StatusCode,
		RequestTime: requestTime,
		Name:        e.Name,
		Ordinal:     i}:
	}
}

func (s *seClient) DoAll(ctx context.Context, wg *sync.WaitGroup, e SearchEngine, i int, result chan<- SearchEngineResponse) {
	defer func() {
		wg.Done()
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.SearchUrl, nil)
	if err != nil {
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		return
	}

	start := time.Now()
	resp, err := s.client.Do(req)
	requestTime := time.Now().UnixMilli() - start.UnixMilli()

	if err != nil {
		log.Debug().Err(err).Msgf("error for engine '%s'", e.Name)
		result <- SearchEngineResponse{Name: e.Name, Ordinal: i, Error: err}
		return
	}
	defer resp.Body.Close()

	result <- SearchEngineResponse{StatusCode: resp.StatusCode, RequestTime: requestTime, Name: e.Name, Ordinal: i}
}
