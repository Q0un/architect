package stats

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/Q0un/architect/proto/api"
)

type API struct {
	api.UnimplementedStatsServiceServer
	service *StatsService
}

// HealthCheck implements api.StatsServiceServer
func (a *API) HealthCheck(ctx context.Context, req *api.HealthCheckRequest) (*api.HealthCheckResponse, error) {
	return &api.HealthCheckResponse{}, nil
}

func (a *API) Mount(router chi.Router) {
	mux := runtime.NewServeMux()
	router.Mount("/", mux)
	_ = api.RegisterStatsServiceHandlerServer(context.Background(), mux, a)
}
