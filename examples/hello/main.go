package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/go-minstack/core"
	"github.com/go-minstack/logger"
)

func run(logger *slog.Logger) {
	logger.Info("Hello from MinStack!", "app", "hello")
}

func main() {
	app := core.New(logger.Module())
	app.Invoke(run)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
	defer app.Stop(ctx)
}
