package users

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Q0un/architect/proto/api"
	proto "github.com/Q0un/architect/proto/tickenator"
)

func HandlerMatcher(key string) (string, bool) {
	if key == "authorization" {
		return key, true
	}
	return "", false
}

type API struct {
	api.UnimplementedUsersServiceServer
	service *UsersService
}

func (a *API) getUserId(ctx context.Context) (uint64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, fmt.Errorf("Pass auth header")
	}
	if len(md["authorization"]) == 0 {
		return 0, fmt.Errorf("Pass auth header")
	}

	token, err := jwt.Parse(md["authorization"][0], func(token *jwt.Token) (interface{}, error) {
		return a.service.jwtPublic, nil
	})
	if err != nil || !token.Valid {
		return 0, fmt.Errorf("Bad auth header")
	}

	idValue, ok := token.Claims.(jwt.MapClaims)["id"]
	if !ok {
		return 0, fmt.Errorf("Bad auth header")
	}

	idStr, ok := idValue.(string)
	if !ok {
		return 0, fmt.Errorf("Bad auth header")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Bad auth header")
	}

	return id, nil
}

// SignUp implements api.UsersServiceServer
func (a *API) SignUp(ctx context.Context, req *api.SignUpRequest) (*api.SignUpResponse, error) {
	token, err := a.service.SignUp(ctx, req)
	if err != nil {
		return nil, err
	}

	tokenHeader := metadata.New(map[string]string{"authorization": token})
	if err := grpc.SendHeader(ctx, tokenHeader); err != nil {
		return nil, fmt.Errorf("Unable to send 'authorization' header")
	}

	return &api.SignUpResponse{}, nil
}

// SignIn implements api.UsersServiceServer
func (a *API) SignIn(ctx context.Context, req *api.SignInRequest) (*api.SignInResponse, error) {
	token, err := a.service.SignIn(ctx, req)
	if err != nil {
		return nil, err
	}

	tokenHeader := metadata.New(map[string]string{"authorization": token})
	if err := grpc.SendHeader(ctx, tokenHeader); err != nil {
		return nil, fmt.Errorf("Unable to send 'authorization' header")
	}

	return &api.SignInResponse{}, nil
}

// SignIn implements api.UsersServiceServer
func (a *API) EditInfo(ctx context.Context, req *api.EditInfoRequest) (*api.EditInfoResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}
	return &api.EditInfoResponse{}, a.service.EditInfo(ctx, req, id)
}

// CreateTicket implements api.TickenatorServiceServer
func (a *API) CreateTicket(ctx context.Context, req *api.CreateTicketHttpRequest) (*api.CreateTicketResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}
	return a.service.tickenator.CreateTicket(
		ctx, 
		&api.CreateTicketRequest{
			UserId:      id,
			Name:        req.Name,
			Description: req.Description,
		},
	)
}

// UpdateTicket implements api.TickenatorServiceServer
func (a *API) UpdateTicket(ctx context.Context, req *api.UpdateTicketHttpRequest) (*api.UpdateTicketResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}
	return a.service.tickenator.UpdateTicket(
		ctx, 
		&api.UpdateTicketRequest{
			UserId:      id,
			Id:          req.Id,
			Name:        req.Name,
			Description: req.Description,
		},
	)
}

// DeleteTicket implements api.TickenatorServiceServer
func (a *API) DeleteTicket(ctx context.Context, req *api.DeleteTicketHttpRequest) (*api.DeleteTicketResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}
	return a.service.tickenator.DeleteTicket(
		ctx, 
		&api.DeleteTicketRequest{
			UserId: id,
			Id:     req.Id,
		},
	)
}

// GetTicket implements api.TickenatorServiceServer
func (a *API) GetTicket(ctx context.Context, req *api.GetTicketHttpRequest) (*proto.Ticket, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}
	return a.service.tickenator.GetTicket(
		ctx, 
		&api.GetTicketRequest{
			Id: req.Id,
		},
	)
}

// ListTickets implements api.TickenatorServiceServer
func (a *API) ListTickets(ctx context.Context, req *api.ListTicketsHttpRequest) (*api.ListTicketsResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}
	return a.service.tickenator.ListTickets(
		ctx,
		&api.ListTicketsRequest{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		},
	)
}

// ViewTicket implements api.TickenatorServiceServer
func (a *API) ViewTicket(ctx context.Context, req *api.ViewTicketRequest) (*api.ViewTicketResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}

	err = a.service.SendKafkaEvent(ctx, req.GetId(), id, "view")
	if err != nil {
		return nil, err
	}

	return &api.ViewTicketResponse{}, nil
}

// LikeTicket implements api.TickenatorServiceServer
func (a *API) LikeTicket(ctx context.Context, req *api.LikeTicketRequest) (*api.LikeTicketResponse, error) {
	id, err := a.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	if !a.service.CheckUser(id) {
		return nil, fmt.Errorf("Bad auth header")
	}

	err = a.service.SendKafkaEvent(ctx, req.GetId(), id, "like")
	if err != nil {
		return nil, err
	}

	return &api.LikeTicketResponse{}, nil
}

func (a *API) Mount(router chi.Router) {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(HandlerMatcher),
		runtime.WithOutgoingHeaderMatcher(HandlerMatcher),
	)
	router.Mount("/api", mux)
	_ = api.RegisterUsersServiceHandlerServer(context.Background(), mux, a)
}

var _ api.UsersServiceServer = (*API)(nil)
