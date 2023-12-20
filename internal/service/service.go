package service

import (
	"context"

	"github.com/plusik10/metrics_collection_service/internal/model"
)

type MetricService interface {
	Create(ctx context.Context, event *model.Event) error
}
