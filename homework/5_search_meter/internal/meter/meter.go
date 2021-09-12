package meter

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter/client"
)

type SearchEngineStats struct {
	TimeMillis int    `json:"timeMillis"`
	EngineName string `json:"engineName"`
}

type Meter interface {
	Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error)
	StartForAllEngines(ctx context.Context, request string) ([]SearchEngineStats, error)
}

type SearchEngineMeter struct {
	client client.SEClient
}

func NewSearchEngineMeter(c client.SEClient) Meter {
	return &SearchEngineMeter{client: c}
}

func (s *SearchEngineMeter) Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error) {
	ctx, cancel := context.WithCancel(ctx)

	result := make(chan client.SearchEngineResponse)
	var wgStart sync.WaitGroup
	var wgWait sync.WaitGroup
	wgStart.Add(1)
	wgWait.Add(1)

	go s.client.MakeRequest(ctx, &wgStart, &wgWait, request, "google", result)

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
