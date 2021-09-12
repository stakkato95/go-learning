package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type SearchEngineResponse struct {
	StatusCode  int
	RequestTime int64
	Error       error
}

type SEClient interface {
	MakeRequest(ctx context.Context, wgStart *sync.WaitGroup, wgWait *sync.WaitGroup, request string, engineName string,
		result chan<- SearchEngineResponse)
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

func (s *seClient) MakeRequest(
	ctx context.Context,
	wgStart *sync.WaitGroup,
	wgWait *sync.WaitGroup,
	request string,
	engineName string,
	result chan<- SearchEngineResponse) {
	addr, err := url.ParseRequestURI(fmt.Sprintf("https://www.google.com/search?q=%s", request))
	if err != nil {
		result <- SearchEngineResponse{Error: err}
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr.String(), nil)
	if err != nil {
		result <- SearchEngineResponse{Error: err}
		return
	}

	wgStart.Done()
	wgWait.Wait()

	start := time.Now()
	resp, err := s.client.Do(req)
	requestTime := time.Now().UnixMilli() - start.UnixMilli()

	if err != nil {
		result <- SearchEngineResponse{Error: err}
		return
	}
	defer resp.Body.Close()

	select {
	case <-ctx.Done():
		log.Debug().Msg("goroutine cancelled")
		return
	case result <- SearchEngineResponse{StatusCode: resp.StatusCode, RequestTime: requestTime}:
	}
}
