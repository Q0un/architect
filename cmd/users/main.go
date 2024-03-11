package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Q0un/architect/internal/users"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Users service failed: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	conf := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	cfg, err := users.LoadConfig(*conf)
	if err != nil {
		return fmt.Errorf("Failed to load config: %w", err)
	}

	var outputPaths []string
	outputPaths = append(outputPaths, "stdout")
	if cfg.LogFile != "" {
		outputPaths = append(outputPaths, cfg.LogFile)
	}

	log := log.New(os.Stdout, "", log.LstdFlags)

	app, err := users.NewApp(log, cfg)
	if err != nil {
		return fmt.Errorf("Failed to create users service: %w", err)
	}

	return app.Run(context.Background())
}
