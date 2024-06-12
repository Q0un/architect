package stats

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/sync/errgroup"
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
		return app.service.runKafkaConsumer(ctx)
	})

	return g.Wait()
}

func (app *App) RunHTTPServer(ctx context.Context) error {
	app.logger.Println("Starting http server at", app.config.HttpAddress)
	server := &http.Server{Addr: app.config.HttpAddress, Handler: app.router}
	return server.ListenAndServe()
}
