package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Q0un/architect/internal/stats"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Stats service failed: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	conf := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	cfg, err := stats.LoadConfig(*conf)
	if err != nil {
		return fmt.Errorf("Failed to load config: %w", err)
	}

	var outputPaths []string
	outputPaths = append(outputPaths, "stdout")
	if cfg.LogFile != "" {
		outputPaths = append(outputPaths, cfg.LogFile)
	}

	log := log.New(os.Stdout, "", log.LstdFlags)

	app, err := stats.NewApp(log, cfg)
	if err != nil {
		return fmt.Errorf("Failed to create stats service: %w", err)
	}

	return app.Run(context.Background())
}
