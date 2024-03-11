package users

import (
	"context"
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Q0un/architect/proto/api"
)

func HandlerMatcher(key string) (string, bool) {
	if key == "authorization" {
		return key, true
	}
	return "", false
}

type API struct {
	api.UnimplementedUsersServiceServer
	logger  *log.Logger
	service *UsersService
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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "Failed to get metadata")
	}

	if len(md["authorization"]) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing 'authorization' header")
	}

	return &api.EditInfoResponse{}, a.service.EditInfo(ctx, req, md["authorization"][0])
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
