package stats

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/Q0un/architect/proto/api"
)

type App struct {
	logger *log.Logger
	config *Config
	router chi.Router

	api     *API
	service *StatsService
}

func NewApp(logger *log.Logger, config *Config) (*App, error) {
	service, err := NewStatsService(logger, config)
	if err != nil {
		return nil, err
	}

	app := &App{
		logger:  logger,
		config:  config,
		api:     &API{service: service},
		service: service,
	}

	err = app.SetupRouter()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) SetupRouter() error {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Logger)

	app.api.Mount(router)

	app.router = router
	return nil
}

func (app *App) Run(ctx context.Context) error {
	app.logger.Println("Starting app")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return app.RunHTTPServer(ctx)
	})
	g.Go(func() error {
		return app.RunGRPCServer(ctx)
	})
	g.Go(func() error {
		return app.service.runKafkaConsumer(ctx)
	})

	return g.Wait()
}

func (app *App) RunHTTPServer(ctx context.Context) error {
	app.logger.Println("Starting http server at", app.config.HttpAddress)
	server := &http.Server{Addr: app.config.HttpAddress, Handler: app.router}
	return server.ListenAndServe()
}

func (app *App) RunGRPCServer(ctx context.Context) error {
	app.logger.Println("Starting grpc server at:", app.config.GrpcAddress)
	server := grpc.NewServer()
	api.RegisterStatsServiceServer(server, app.api)

	listen, err := net.Listen("tcp", app.config.GrpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen grpc server socket: %w", err)
	}

	if err = server.Serve(listen); err != nil {
		return fmt.Errorf("grpc server failed: %w", err)
	}

	return nil
}
