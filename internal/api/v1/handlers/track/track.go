package track

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/plusik10/metrics_collection_service/internal/model"
	"github.com/plusik10/metrics_collection_service/internal/service"
)

func New(ctx context.Context, service service.MetricService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event model.Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = service.Create(ctx, &event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	}
}
