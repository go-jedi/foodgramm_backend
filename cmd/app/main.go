package main

import (
	"context"
	"log"

	"github.com/go-jedi/foodgramm_backend/internal/app"
)

// @title API application for telegram web apps application foodgrammm
// @version 1.0
// @description This is a application for telegram web apps application.

// @host localhost:50050
// @BasePath /v1
func main() {
	ctx := context.Background()

	// initialize app
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	// run application
	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
