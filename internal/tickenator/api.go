package tickenator

import (
	"context"

	"github.com/Q0un/architect/proto/api"
	proto "github.com/Q0un/architect/proto/tickenator"
)

type API struct {
	api.UnimplementedTickenatorServiceServer
	service *TickenatorService
}

// CreateTicket implements api.TickenatorServiceServer
func (a *API) CreateTicket(ctx context.Context, req *api.CreateTicketRequest) (*api.CreateTicketResponse, error) {
	id, err := a.service.CreateTicket(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return &api.CreateTicketResponse{
			Id: id,
		}, nil
	}
}

// UpdateTicket implements api.TickenatorServiceServer
func (a *API) UpdateTicket(ctx context.Context, req *api.UpdateTicketRequest) (*api.UpdateTicketResponse, error) {
	err := a.service.UpdateTicket(ctx, req)
	return &api.UpdateTicketResponse{}, err
}

// DeleteTicket implements api.TickenatorServiceServer
func (a *API) DeleteTicket(ctx context.Context, req *api.DeleteTicketRequest) (*api.DeleteTicketResponse, error) {
	err := a.service.DeleteTicket(ctx, req)
	return &api.DeleteTicketResponse{}, err
}

// GetTicket implements api.TickenatorServiceServer
func (a *API) GetTicket(ctx context.Context, req *api.GetTicketRequest) (*proto.Ticket, error) {
	return a.service.GetTicket(ctx, req)
}

// ListTickets implements api.TickenatorServiceServer
func (a *API) ListTickets(ctx context.Context, req *api.ListTicketsRequest) (*api.ListTicketsResponse, error) {
	tickets, err := a.service.ListTickets(ctx, req)
	if err != nil {
		return &api.ListTicketsResponse{}, err
	}
	return &api.ListTicketsResponse{
		Tickets: tickets,
	}, nil
}
