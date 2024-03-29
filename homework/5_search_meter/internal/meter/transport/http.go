package transport

import (
	"encoding/json"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter"
)

//body for '/meter/single'
type singleSearchEngineRequest struct {
	Request       string `json:"request"`
	TimeoutMillis int    `timeoutMillis:"timeoutMillis"`
}

//body for '/meter/all'
type allSearchEnginesRequest struct {
	Request string `json:"request"`
}

type HTTP struct {
	engineMeter meter.Meter
}

func NewHTTP(meter meter.Meter) *HTTP {
	return &HTTP{engineMeter: meter}
}

func (h *HTTP) StartSearchEngineRequest(w http.ResponseWriter, r *http.Request) {
	var engineRequest singleSearchEngineRequest

	//validate body json and its fields
	if err := json.NewDecoder(r.Body).Decode(&engineRequest); err != nil {
		log.Debug().Err(err).Msg("can not decode singleSearchEngineRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ok, err := valid.ValidateStruct(&engineRequest); !ok {
		log.Debug().Err(err).Msg("invalid singleSearchEngineRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//start metering and get the fastest search engine
	stats, err := h.engineMeter.Start(r.Context(), engineRequest.Request, engineRequest.TimeoutMillis)
	if err != nil {
		log.Debug().Err(err).Msg("metering failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//error while encoding json should not happen, but check it anyway
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		log.Debug().Err(err).Msg("failed to encode SearchEngineStats")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *HTTP) StartAllSearchEnginesRequest(w http.ResponseWriter, r *http.Request) {
	var enginesRequest allSearchEnginesRequest

	if err := json.NewDecoder(r.Body).Decode(&enginesRequest); err != nil {
		log.Debug().Err(err).Msg("can not decode allSearchEnginesRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ok, err := valid.ValidateStruct(&enginesRequest); !ok {
		log.Debug().Err(err).Msg("invalid allSearchEnginesRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//start metering and get stats for all search engines
	stats := h.engineMeter.StartAll(r.Context(), enginesRequest.Request)

	//error while encoding json should not happen, but check it anyway
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		log.Debug().Err(err).Msg("failed to encode []SearchEngineStats")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
