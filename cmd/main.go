package main

import (
	"context"
	"log"

	"github.com/plusik10/metrics_collection_service/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("error creating app: %v", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("error running app: %v", err)
	}
}
