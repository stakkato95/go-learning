package meter

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type SearchEngineStats struct {
	TimeMillis int    `json:"timeMillis"`
	EngineName string `json:"engineName"`
}

type searchEngineResponse struct {
	StatusCode  int
	RequestTime int64
	Error       error
}

type Meter interface {
	Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error)
	StartForAllEngines(ctx context.Context, request string) ([]SearchEngineStats, error)
}

type SearchEngineMeter struct {
	//client *seclient
}

func NewSearchEngineMeter() Meter {
	return &SearchEngineMeter{}
}

func (s *SearchEngineMeter) Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error) {
	ctx, cancel := context.WithCancel(ctx)

	result := make(chan searchEngineResponse)
	var wgStart sync.WaitGroup
	var wgWait sync.WaitGroup
	wgStart.Add(1)
	wgWait.Add(1)

	go makeRequest(ctx, &wgStart, &wgWait, request, "google", result)
	// go makeRequest(ctx, &wgStart, &wgWait, request, "google", result)

	wgStart.Wait()
	wgWait.Done()
	resonse := <-result
	cancel()
	log.Debug().Msgf("response: %s", resonse)

	return SearchEngineStats{TimeMillis: 100500, EngineName: "google"}, nil
}

func (s *SearchEngineMeter) StartForAllEngines(ctx context.Context, request string) ([]SearchEngineStats, error) {
	stats := []SearchEngineStats{
		{TimeMillis: 100500, EngineName: "google"},
		{TimeMillis: 1, EngineName: "bing"},
	}

	return stats, nil
}

func makeRequest(ctx context.Context, wgStart *sync.WaitGroup, wgWait *sync.WaitGroup, request string, engineName string, result chan<- searchEngineResponse) {
	start := time.Now()
	wgStart.Done()
	wgWait.Wait()
	resp, err := http.Get("https://www.google.com/search?q=hello")
	requestTime := time.Now().UnixMilli() - start.UnixMilli()

	select {
	case <-ctx.Done():
	case result <- searchEngineResponse{StatusCode: resp.StatusCode, RequestTime: requestTime, Error: err}:
	}
}
