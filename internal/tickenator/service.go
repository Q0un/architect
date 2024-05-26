package tickenator

import (
	"context"
	"fmt"
	"log"
	
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Q0un/architect/proto/api"
	proto "github.com/Q0un/architect/proto/tickenator"
)

type TickenatorService struct {
	logger *log.Logger
	db     *sqlx.DB
}

func NewTickenatorService(logger *log.Logger, config *Config) (*TickenatorService, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			config.Db.User,
			config.Db.Password,
			config.Db.Name,
			config.Db.Host,
			config.Db.Port,
		),
	)
	if err != nil {
		return nil, err
	}

	return &TickenatorService{
		logger: logger,
		db:     db,
	}, nil
}

func (tickenator *TickenatorService) CreateTicket(ctx context.Context, req *api.CreateTicketRequest) (uint64, error) {
	var ticketId uint64
	tx := tickenator.db.MustBegin()
	err := tx.QueryRowx(
		"INSERT INTO tickets (author_id, name, description) VALUES ($1, $2, $3) RETURNING id",
		req.GetUserId(),
		req.GetName(),
		req.GetDescription(),
	).Scan(&ticketId)

	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("Ticket cannot be added to database")
	}

	tx.Commit()

	return ticketId, nil
}

func (tickenator *TickenatorService) UpdateTicket(ctx context.Context, req *api.UpdateTicketRequest) error {
	var ticket proto.Ticket
	err := tickenator.db.Get(&ticket, "SELECT * FROM tickets WHERE id=$1", req.GetId())

	if err != nil {
		return fmt.Errorf("There is no ticket with such id")
	}

	if req.GetUserId() != ticket.GetAuthorId() {
		return fmt.Errorf("You don't have access to edit this ticket")
	}

	if req.Name != nil {
		ticket.Name = req.GetName()
	}
	if req.Description != nil {
		ticket.Description = req.GetDescription()
	}

	_, err = tickenator.db.NamedExec(
		"UPDATE tickets SET name=:name, description=:description WHERE id = :id",
		ticket,
	)
	if err != nil {
		return fmt.Errorf("Ticket cannot be edited in database")
	}

	return nil
}

func (tickenator *TickenatorService) DeleteTicket(ctx context.Context, req *api.DeleteTicketRequest) error {
	var ticket proto.Ticket
	err := tickenator.db.Get(&ticket, "SELECT * FROM tickets WHERE id=$1", req.GetId())

	if err != nil {
		return fmt.Errorf("There is no ticket with such id")
	}

	if req.GetUserId() != ticket.GetAuthorId() {
		return fmt.Errorf("You don't have access to delete this ticket")
	}

	_, err = tickenator.db.NamedExec(
		"DELETE FROM tickets WHERE id = :id",
		ticket,
	)
	if err != nil {
		return fmt.Errorf("Ticket cannot be deleted in database")
	}

	return nil
}

func (tickenator *TickenatorService) GetTicket(ctx context.Context, req *api.GetTicketRequest) (*proto.Ticket, error) {
	var ticket proto.Ticket
	err := tickenator.db.Get(&ticket, "SELECT * FROM tickets WHERE id=$1", req.GetId())
	if err != nil {
		return nil, fmt.Errorf("There is no ticket with such id")
	}

	return &ticket, nil
}

func (tickenator *TickenatorService) ListTickets(ctx context.Context, req *api.ListTicketsRequest) ([]*proto.Ticket, error) {
	var tickets []*proto.Ticket
	err := tickenator.db.Select(&tickets, "SELECT * FROM tickets OFFSET $1 LIMIT $2", (req.GetPageNum() - 1) * req.GetPageSize(), req.GetPageSize())

	if err != nil {
		return []*proto.Ticket{}, nil
	}

	return tickets, nil
}
