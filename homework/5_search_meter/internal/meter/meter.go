package meter

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter/client"
)

//result of metering a single search engine
type SearchEngineStats struct {
	TimeMillis int64  `json:"timeMillis"`
	EngineName string `json:"engineName"`
}

type Meter interface {
	Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error)
	StartAll(ctx context.Context, request string) []SearchEngineStats
}

//private implementation of Meter interface
type searchEngineMeter struct {
	client client.SEClient
}

//since searchEngineMeter is private, return interface
func NewSearchEngineMeter(c client.SEClient) Meter {
	return &searchEngineMeter{client: c}
}

//returns fastest search engine
func (s *searchEngineMeter) Start(ctx context.Context, request string, timeoutMillis int) (SearchEngineStats, error) {
	//take ctx and extend it with a timeout. 'cancel' can be used to stop execution before deadline
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutMillis)*time.Millisecond)
	defer cancel()

	//chan to accept result of only of the fastest search engine
	resultChan := make(chan client.SearchEngineResponse)

	var wgAllDone sync.WaitGroup //when no more goroutines will write to resultChan, close resultChan
	var wgStart sync.WaitGroup   //wait till all goroutines are ready to make request
	var wgWait sync.WaitGroup    //signal all goroutines to make request

	engines := client.GetSearchEngines()
	wgAllDone.Add(len(engines))
	wgStart.Add(len(engines))
	wgWait.Add(1)

	go func() {
		wgAllDone.Wait() //when all goroutines are done, close resultChan to avoid goroutine leakage
		close(resultChan)
	}()
	for i, engine := range engines {
		//start request for each seach engine
		go s.client.Do(ctx, &wgStart, &wgWait, &wgAllDone, engine, i, resultChan)
	}

	wgStart.Wait()               //wait till all goroutines have created requests and are ready to do request
	wgWait.Done()                //signal all goroutines to make request
	response, ok := <-resultChan //wait for response from fastest search engine OR check that channel was closed
	if !ok {
		//there was no search engine that served resonse within 'timeoutMillis' deadline
		log.Debug().Msgf("no response received")
		return SearchEngineStats{}, nil
	}

	//log and return result
	log.Debug().Msgf("response: %s", response)
	return SearchEngineStats{TimeMillis: response.RequestTime, EngineName: response.Name}, nil
}

func (s *searchEngineMeter) StartAll(ctx context.Context, request string) []SearchEngineStats {
	var wg sync.WaitGroup //wait till all goroutines have created requests and are ready to do request

	//chan to accept result of all search engines, including errors
	resultChan := make(chan client.SearchEngineResponse)

	engines := client.GetSearchEngines()
	wg.Add(len(engines))

	go func() {
		wg.Wait() //when all goroutines are done, close resultChan to avoid goroutine leakage
		close(resultChan)
	}()
	for i, engine := range engines {
		//start request for each seach engine
		go s.client.DoAll(ctx, &wg, engine, i, resultChan)
	}

	stats := make([]SearchEngineStats, 0)
	for response := range resultChan {
		//get results of all search engine: errors and actual results
		stats = append(stats, SearchEngineStats{TimeMillis: response.RequestTime, EngineName: response.Name})
	}

	return stats
}
