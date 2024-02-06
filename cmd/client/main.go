package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Chystik/pass-man/config/client"
	run "github.com/Chystik/pass-man/run/client"
)

func main() {
	cfg := config.NewClientConfig()

	parseFlags(cfg)

	// Graceful shutdown setup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	run.Client(ctx, cfg)
}
