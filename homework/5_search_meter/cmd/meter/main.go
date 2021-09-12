package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter"
	"github.com/stakkato95/searchmeter/internal/meter/transport"
)

func main() {
	searchEngineMeter := meter.NewSearchEngineMeter()
	searchHttp := transport.NewHTTP(searchEngineMeter)

	mux := http.NewServeMux()
	mux.HandleFunc("/search", searchHttp.StartSearchEngineRequest)

	addr := ":8080"
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Debug().Msgf("starting server at '%s'\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Debug().Err(err).Msgf("error when starting server at %s", addr)
	}
}
