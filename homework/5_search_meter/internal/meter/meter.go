package meter

import (
	"context"
	"sync"
	"time"

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
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutMillis)*time.Millisecond)
	defer cancel()

	resultChan := make(chan client.SearchEngineResponse)

	var wgAllDone sync.WaitGroup
	var wgStart sync.WaitGroup
	var wgWait sync.WaitGroup

	engines := client.GetSearchEngines()
	wgAllDone.Add(len(engines))
	wgWait.Add(1)

	go func() {
		wgAllDone.Wait()
		close(resultChan)
	}()
	for i, engine := range engines {
		wgStart.Add(1)
		go s.client.Do(ctx, &wgStart, &wgWait, &wgAllDone, engine, i, resultChan)
	}

	wgStart.Wait() //wait till all goroutines have created requests and are ready to do request
	wgWait.Done()  //signal all goroutines to make request
	response, ok := <-resultChan
	if !ok {
		log.Debug().Msgf("no response received")
		return SearchEngineStats{}, nil
	}

	log.Debug().Msgf("response: %s", response)
	return SearchEngineStats{TimeMillis: response.RequestTime, EngineName: response.Name}, nil
}

func (s *SearchEngineMeter) StartAll(ctx context.Context, request string) []SearchEngineStats {
	var wg sync.WaitGroup

	resultChan := make(chan client.SearchEngineResponse)

	engines := client.GetSearchEngines()
	wg.Add(len(engines))

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for i, engine := range engines {
		go s.client.DoAll(ctx, &wg, engine, i, resultChan)
	}

	stats := make([]SearchEngineStats, 0)
	for response := range resultChan {
		stats = append(stats, SearchEngineStats{TimeMillis: response.RequestTime, EngineName: response.Name})
	}

	return stats
}
