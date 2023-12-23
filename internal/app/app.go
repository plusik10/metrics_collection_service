package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/plusik10/metrics_collection_service/internal/api/v1/handlers/track"
)

type App struct {
	httpServer      *http.Server
	serviceProvider *serviceProvider
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		err := a.runPublicHTTP()
		if err != nil {
			log.Println("error running: ", err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v\n", err)
	}
	log.Println("HTTP server shutdown")

	return nil
}

func (a *App) initPublicHTTP(ctx context.Context) error {
	r := chi.NewRouter()
	r.Post("/track", track.New(ctx, a.serviceProvider.GetMetricService(ctx)))

	a.httpServer = &http.Server{
		Handler:           r,
		Addr:              a.serviceProvider.GetConfig().HTTP.Port,
		ReadHeaderTimeout: a.serviceProvider.GetConfig().HTTP.Timeout,
	}

	return nil
}

//nolint:revive
func (a *App) runPublicHTTP() error {
	fmt.Println("Starting public http server")
	if err := a.httpServer.ListenAndServe(); err != nil {
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

func (a *App) initServiceProvider(ctx context.Context) error {
	_ = ctx
	a.serviceProvider = newServiceProvider()
	return nil
}
