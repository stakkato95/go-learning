package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter"
	"github.com/stakkato95/searchmeter/internal/meter/client"
	"github.com/stakkato95/searchmeter/internal/meter/transport"
)

func main() {
	seclient := client.NewClient()
	searchEngineMeter := meter.NewSearchEngineMeter(seclient)
	searchHttp := transport.NewHTTP(searchEngineMeter)

	mux := http.NewServeMux()
	mux.HandleFunc("/meter/single", searchHttp.StartSearchEngineRequest)
	mux.HandleFunc("/meter/all", searchHttp.StartAllSearchEnginesRequest)

	addr := ":8080"
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Debug().Msgf("starting server at '%s'", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Debug().Err(err).Msgf("error when starting server at %s", addr)
	}
}
