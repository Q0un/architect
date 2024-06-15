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

// TicketStats implements api.StatsServiceServer
func (a *API) TicketStats(ctx context.Context, req *api.TicketStatsRequest) (*api.TicketStatsResponse, error) {
	views, likes, err := a.service.TicketStats(req.GetId())
	if err != nil {
		return nil, err
	}

	return &api.TicketStatsResponse{
		Views: views,
		Likes: likes,
	}, nil
}

// TopTickets implements api.StatsServiceServer
func (a *API) TopTickets(ctx context.Context, req *api.TopTicketsRequest) (*api.TopTicketsResponse, error) {
	top, err := a.service.TopTickets(req.GetType())
	if err != nil {
		return nil, err
	}

	return &api.TopTicketsResponse{
		Top: top,
	}, nil
}

// TopUsers implements api.StatsServiceServer
func (a *API) TopUsers(ctx context.Context, req *api.TopUsersRequest) (*api.TopUsersResponse, error) {
	top, err := a.service.TopUsers()
	if err != nil {
		return nil, err
	}

	return &api.TopUsersResponse{
		Top: top,
	}, nil
}

func (a *API) Mount(router chi.Router) {
	mux := runtime.NewServeMux()
	router.Mount("/", mux)
	_ = api.RegisterStatsServiceHandlerServer(context.Background(), mux, a)
}
