package app

import (
	"context"
	"log"

	"github.com/plusik10/metrics_collection_service/internal/config"
	"github.com/plusik10/metrics_collection_service/internal/repository"
	"github.com/plusik10/metrics_collection_service/internal/repository/metric"
	"github.com/plusik10/metrics_collection_service/internal/service"
	desc "github.com/plusik10/metrics_collection_service/internal/service/metric"
	"github.com/plusik10/metrics_collection_service/pkg/db"
)

type serviceProvider struct {
	db     db.Client
	config *config.Config

	repository repository.Repository

	metricService service.MetricService
}

func newServiceProvider() *serviceProvider {
	sp := &serviceProvider{}
	sp.config = sp.GetConfig()
	return sp
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %v", err)
		}
		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}
		s.db = dbc
	}
	return s.db
}

func (s *serviceProvider) GetRepository(ctx context.Context) repository.Repository {
	if s.repository == nil {
		s.repository = metric.NewMetricRepository(s.GetDB(ctx))
	}
	return s.repository
}

func (s *serviceProvider) GetMetricService(ctx context.Context) service.MetricService {
	if s.metricService == nil {
		s.metricService = desc.NewMetricService(s.GetRepository(ctx))
	}
	return s.metricService
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to create config: %v", err)
		}
		s.config = cfg
	}
	return s.config
}
