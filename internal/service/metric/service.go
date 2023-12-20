package metric

import (
	"context"

	"github.com/plusik10/metrics_collection_service/internal/model"
	"github.com/plusik10/metrics_collection_service/internal/repository"
	"github.com/plusik10/metrics_collection_service/internal/service"
)

type metricService struct {
	repo repository.Repository
}

func (m *metricService) Create(ctx context.Context, event *model.Event) error {
	err := m.repo.Create(ctx, event)
	return err
}

func NewMetricService(repo repository.Repository) service.MetricService {
	return &metricService{repo: repo}
}
