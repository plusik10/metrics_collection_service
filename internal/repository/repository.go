package repository

import (
	"context"

	"github.com/plusik10/metrics_collection_service/internal/model"
)

type Repository interface {
	Create(ctx context.Context, event *model.Event) error
}
