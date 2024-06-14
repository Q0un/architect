package users

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/Q0un/architect/proto/api"
)

type StatsClient struct {
	conn   *grpc.ClientConn
	client api.StatsServiceClient
}

func NewStatsClient(logger *log.Logger, c *Config) (*StatsClient, error) {
	logger.Println("Connecting to stats host:", c.StatsHost)
	cc, err := grpc.Dial(c.StatsHost, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &StatsClient{
		conn:   cc,
		client: api.NewStatsServiceClient(cc),
	}, nil
}


func (client *StatsClient) TicketStats(ctx context.Context, req *api.TicketStatsRequest) (*api.TicketStatsResponse, error) {
	return client.client.TicketStats(ctx, req)
}

func (client *StatsClient) TopTickets(ctx context.Context, req *api.TopTicketsRequest) (*api.TopTicketsResponse, error) {
	return client.client.TopTickets(ctx, req)
}

func (client *StatsClient) TopUsers(ctx context.Context, req *api.TopUsersRequest) (*api.TopUsersResponse, error) {
	return client.client.TopUsers(ctx, req)
}

func (client *StatsClient) Close() {
	_ = client.conn.Close()
}