package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/plusik10/metrics_collection_service/internal/api/v1/handlers/track"
)

type app struct {
	handler         *chi.Mux
	serviceProvider *serviceProvider
}

func (a *app) initPublicHttp(ctx context.Context) error {
	r := chi.NewRouter()
	r.Post("/track", track.New(ctx, a.serviceProvider.GetMetricService(ctx)))
	a.handler = r
	return nil
}

func NewApp(ctx context.Context) (*app, error) {
	a := &app{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *app) Run() error {
	defer func() {
		_ = a.serviceProvider.db.Close()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	err := a.runPublicHttp()
	if err != nil {
		log.Fatalf("failed to process mux: %v", err)
	}
	<-stop
	return nil
}

// runPublicHttp - runs the public http server
// TODO: Graceful shutdown
func (a *app) runPublicHttp() error {
	httpPort := a.serviceProvider.GetConfig().HTTP.Port

	if err := http.ListenAndServe(httpPort, a.handler); err != nil {
		return err
	}

	return nil
}

func (a *app) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initPublicHttp,
	}
	for _, dep := range deps {
		err := dep(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *app) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
