package meter

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter/client"
)

type SearchEngineStats struct {
	TimeMillis int64  `json:"timeMillis"`
	EngineName string `json:"engineName"`
}

type Meter interface {
	Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error)
	StartAll(ctx context.Context, request string) []SearchEngineStats
}

type SearchEngineMeter struct {
	client client.SEClient
}

func NewSearchEngineMeter(c client.SEClient) Meter {
	return &SearchEngineMeter{client: c}
}

func (s *SearchEngineMeter) Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error) {
	ctx, cancel := context.WithCancel(ctx)

	resultChan := make(chan client.SearchEngineResponse)

	var wgStart sync.WaitGroup
	var wgWait sync.WaitGroup
	wgWait.Add(1)

	for i, engine := range client.GetSearchEngines() {
		wgStart.Add(1)
		go s.client.Do(ctx, &wgStart, &wgWait, engine, i, resultChan)
	}

	wgStart.Wait() //wait till all goroutines have created requests and are ready to do request
	wgWait.Done()  //signal all goroutines to make request
	response := <-resultChan
	cancel()
	close(resultChan)
	log.Debug().Msgf("response: %s", response)

	return SearchEngineStats{TimeMillis: response.RequestTime, EngineName: response.Name}, nil
}

func (s *SearchEngineMeter) StartAll(ctx context.Context, request string) []SearchEngineStats {
	var wg sync.WaitGroup

	resultChan := make(chan client.SearchEngineResponse)
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	engines := client.GetSearchEngines()
	wg.Add(len(engines))
	for i, engine := range engines {
		go s.client.DoAll(ctx, &wg, engine, i, resultChan)
	}

	stats := make([]SearchEngineStats, 0)
	for response := range resultChan {
		stats = append(stats, SearchEngineStats{TimeMillis: response.RequestTime, EngineName: response.Name})
	}

	return stats
}
