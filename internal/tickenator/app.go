package tickenator

import (
	"context"
	"fmt"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/Q0un/architect/proto/api"
)

type App struct {
	logger *log.Logger
	config *Config

	api     *API
	service *TickenatorService
}

func NewApp(logger *log.Logger, config *Config) (*App, error) {
	service, err := NewTickenatorService(logger, config)
	if err != nil {
		return nil, err
	}

	app := &App{
		logger:  logger,
		config:  config,
		api:     &API{service: service},
		service: service,
	}

	return app, nil
}

func (app *App) Run(ctx context.Context) error {
	app.logger.Println("Starting app")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return app.RunGRPCServer(ctx)
	})

	return g.Wait()
}

func (app *App) RunGRPCServer(ctx context.Context) error {
	app.logger.Println("Starting grpc server at:", app.config.GrpcAddress)
	server := grpc.NewServer()
	api.RegisterTickenatorServiceServer(server, app.api)

	listen, err := net.Listen("tcp", app.config.GrpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen grpc server socket: %w", err)
	}

	if err = server.Serve(listen); err != nil {
		return fmt.Errorf("grpc server failed: %w", err)
	}

	return nil
}
