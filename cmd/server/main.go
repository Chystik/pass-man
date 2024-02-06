package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Chystik/pass-man/config/server"
	run "github.com/Chystik/pass-man/run/server"
)

func main() {
	var envFilePath string
	cfg := config.NewServerConfig()

	flag.StringVar(&envFilePath, "e", "", "path to the .env file")
	flag.Parse()

	if envFilePath != "" {
		err := loadEnvFromFile(envFilePath)
		if err != nil {
			panic(err)
		}
	}
	parseEnv(cfg)

	// Graceful shutdown setup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	run.Server(ctx, cfg)
}
