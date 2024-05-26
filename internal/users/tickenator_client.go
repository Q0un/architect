package users

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/Q0un/architect/proto/api"
	proto "github.com/Q0un/architect/proto/tickenator"
)

type TickenatorClient struct {
	conn   *grpc.ClientConn
	client api.TickenatorServiceClient
}

func NewTickenatorClient(logger *log.Logger, c *Config) (*TickenatorClient, error) {
	logger.Println("Connecting to tickenator host:", c.TickenatorHost)
	cc, err := grpc.Dial(c.TickenatorHost, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &TickenatorClient{
		conn:   cc,
		client: api.NewTickenatorServiceClient(cc),
	}, nil
}


func (client *TickenatorClient) CreateTicket(ctx context.Context, req *api.CreateTicketRequest) (*api.CreateTicketResponse, error) {
	return client.client.CreateTicket(ctx, req)
}

func (client *TickenatorClient) UpdateTicket(ctx context.Context, req *api.UpdateTicketRequest) (*api.UpdateTicketResponse, error) {
	return client.client.UpdateTicket(ctx, req)
}

func (client *TickenatorClient) DeleteTicket(ctx context.Context, req *api.DeleteTicketRequest) (*api.DeleteTicketResponse, error) {
	return client.client.DeleteTicket(ctx, req)
}

func (client *TickenatorClient) GetTicket(ctx context.Context, req *api.GetTicketRequest) (*proto.Ticket, error) {
	return client.client.GetTicket(ctx, req)
}

func (client *TickenatorClient) ListTickets(ctx context.Context, req *api.ListTicketsRequest) (*api.ListTicketsResponse, error) {
	return client.client.ListTickets(ctx, req)
}

func (client *TickenatorClient) Close() {
	_ = client.conn.Close()
}