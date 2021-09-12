package transport

import (
	"encoding/json"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/rs/zerolog/log"
	"github.com/stakkato95/searchmeter/internal/meter"
)

type searchEngineRequest struct {
	Request       string `json:"request"`
	TimeoutMillis int    `timeoutMillis:"timeoutMillis"`
}

type HTTP struct {
	engineMeter meter.Meter
}

func NewHTTP(meter meter.Meter) *HTTP {
	return &HTTP{engineMeter: meter}
}

func (h *HTTP) StartSearchEngineRequest(w http.ResponseWriter, r *http.Request) {
	var engineRequest searchEngineRequest

	if err := json.NewDecoder(r.Body).Decode(&engineRequest); err != nil {
		log.Debug().Err(err).Msg("can not decode searchEngineRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ok, err := valid.ValidateStruct(&engineRequest); !ok {
		log.Debug().Err(err).Msg("invalid searchEngineRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stats, err := h.engineMeter.Start(r.Context(), engineRequest.Request, engineRequest.TimeoutMillis)
	if err != nil {
		log.Debug().Err(err).Msg("metering failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(stats); err != nil {
		log.Debug().Err(err).Msg("failed to encode SearchEngineStats")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTP) StartAllSearchEnginesRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
	w.WriteHeader(http.StatusOK)
}
