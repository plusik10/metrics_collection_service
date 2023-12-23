package track

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/plusik10/metrics_collection_service/internal/model"
	"github.com/plusik10/metrics_collection_service/internal/service"
)

type ResponseTrack struct {
	OK  bool
	Err string
}

func New(ctx context.Context, service service.MetricService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var event model.Event

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			sendResponse(w, false, errors.New("Expected json object"), http.StatusBadRequest)
			return
		}

		err = event.Validate()
		if err != nil {
			sendResponse(w, false, err, http.StatusBadRequest)
			return
		}

		err = service.Create(ctx, &event)
		if err != nil {
			sendResponse(w, false, err, http.StatusInternalServerError)
		}

		sendResponse(w, true, nil, http.StatusCreated)
	}
}

func sendResponse(w http.ResponseWriter, ok bool, err error, statusCode int) {
	resp := ResponseTrack{OK: ok}
	if err != nil {
		resp.Err = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
