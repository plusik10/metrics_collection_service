package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/plusik10/metrics_collection_service/internal/api/v1/handlers/track"
)

type App struct {
	handler         *chi.Mux
	serviceProvider *serviceProvider
}

func (a *App) initPublicHTTP(ctx context.Context) error {
	r := chi.NewRouter()
	r.Post("/track", track.New(ctx, a.serviceProvider.GetMetricService(ctx)))
	a.handler = r

	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		_ = a.serviceProvider.db.Close()
	}()

	err := a.runPublicHTTP()
	if err != nil {
		return err
	}
	return nil
}

//nolint:revive
func (a *App) runPublicHTTP() error {
	httpPort := a.serviceProvider.GetConfig().HTTP.Port
	fmt.Println("Starting public http server")
	if err := http.ListenAndServe(httpPort, a.handler); err != nil {
		return err
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initPublicHTTP,
	}
	for _, dep := range deps {
		err := dep(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//nolint:revive
func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
