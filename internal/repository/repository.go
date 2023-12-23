package repository

import (
	"context"

	"github.com/plusik10/metrics_collection_service/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Repository
type Repository interface {
	Create(ctx context.Context, event *model.Event) error
}
