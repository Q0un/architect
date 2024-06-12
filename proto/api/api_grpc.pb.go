// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api/api.proto

package api

import (
	context "context"
	tickenator "github.com/Q0un/architect/proto/tickenator"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UsersService_SignUp_FullMethodName       = "/api.UsersService/SignUp"
	UsersService_SignIn_FullMethodName       = "/api.UsersService/SignIn"
	UsersService_EditInfo_FullMethodName     = "/api.UsersService/EditInfo"
	UsersService_CreateTicket_FullMethodName = "/api.UsersService/CreateTicket"
	UsersService_UpdateTicket_FullMethodName = "/api.UsersService/UpdateTicket"
	UsersService_DeleteTicket_FullMethodName = "/api.UsersService/DeleteTicket"
	UsersService_GetTicket_FullMethodName    = "/api.UsersService/GetTicket"
	UsersService_ListTickets_FullMethodName  = "/api.UsersService/ListTickets"
	UsersService_ViewTicket_FullMethodName   = "/api.UsersService/ViewTicket"
	UsersService_LikeTicket_FullMethodName   = "/api.UsersService/LikeTicket"
)

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	EditInfo(ctx context.Context, in *EditInfoRequest, opts ...grpc.CallOption) (*EditInfoResponse, error)
	CreateTicket(ctx context.Context, in *CreateTicketHttpRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketHttpRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketHttpRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
	GetTicket(ctx context.Context, in *GetTicketHttpRequest, opts ...grpc.CallOption) (*tickenator.Ticket, error)
	ListTickets(ctx context.Context, in *ListTicketsHttpRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error)
	ViewTicket(ctx context.Context, in *ViewTicketRequest, opts ...grpc.CallOption) (*ViewTicketResponse, error)
	LikeTicket(ctx context.Context, in *LikeTicketRequest, opts ...grpc.CallOption) (*LikeTicketResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, UsersService_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, UsersService_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) EditInfo(ctx context.Context, in *EditInfoRequest, opts ...grpc.CallOption) (*EditInfoResponse, error) {
	out := new(EditInfoResponse)
	err := c.cc.Invoke(ctx, UsersService_EditInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) CreateTicket(ctx context.Context, in *CreateTicketHttpRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, UsersService_CreateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketHttpRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error) {
	out := new(UpdateTicketResponse)
	err := c.cc.Invoke(ctx, UsersService_UpdateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketHttpRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, UsersService_DeleteTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetTicket(ctx context.Context, in *GetTicketHttpRequest, opts ...grpc.CallOption) (*tickenator.Ticket, error) {
	out := new(tickenator.Ticket)
	err := c.cc.Invoke(ctx, UsersService_GetTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ListTickets(ctx context.Context, in *ListTicketsHttpRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error) {
	out := new(ListTicketsResponse)
	err := c.cc.Invoke(ctx, UsersService_ListTickets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ViewTicket(ctx context.Context, in *ViewTicketRequest, opts ...grpc.CallOption) (*ViewTicketResponse, error) {
	out := new(ViewTicketResponse)
	err := c.cc.Invoke(ctx, UsersService_ViewTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) LikeTicket(ctx context.Context, in *LikeTicketRequest, opts ...grpc.CallOption) (*LikeTicketResponse, error) {
	out := new(LikeTicketResponse)
	err := c.cc.Invoke(ctx, UsersService_LikeTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	EditInfo(context.Context, *EditInfoRequest) (*EditInfoResponse, error)
	CreateTicket(context.Context, *CreateTicketHttpRequest) (*CreateTicketResponse, error)
	UpdateTicket(context.Context, *UpdateTicketHttpRequest) (*UpdateTicketResponse, error)
	DeleteTicket(context.Context, *DeleteTicketHttpRequest) (*DeleteTicketResponse, error)
	GetTicket(context.Context, *GetTicketHttpRequest) (*tickenator.Ticket, error)
	ListTickets(context.Context, *ListTicketsHttpRequest) (*ListTicketsResponse, error)
	ViewTicket(context.Context, *ViewTicketRequest) (*ViewTicketResponse, error)
	LikeTicket(context.Context, *LikeTicketRequest) (*LikeTicketResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedUsersServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedUsersServiceServer) EditInfo(context.Context, *EditInfoRequest) (*EditInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditInfo not implemented")
}
func (UnimplementedUsersServiceServer) CreateTicket(context.Context, *CreateTicketHttpRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedUsersServiceServer) UpdateTicket(context.Context, *UpdateTicketHttpRequest) (*UpdateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
}
func (UnimplementedUsersServiceServer) DeleteTicket(context.Context, *DeleteTicketHttpRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedUsersServiceServer) GetTicket(context.Context, *GetTicketHttpRequest) (*tickenator.Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedUsersServiceServer) ListTickets(context.Context, *ListTicketsHttpRequest) (*ListTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}
func (UnimplementedUsersServiceServer) ViewTicket(context.Context, *ViewTicketRequest) (*ViewTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewTicket not implemented")
}
func (UnimplementedUsersServiceServer) LikeTicket(context.Context, *LikeTicketRequest) (*LikeTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeTicket not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_EditInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).EditInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_EditInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).EditInfo(ctx, req.(*EditInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateTicket(ctx, req.(*CreateTicketHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_UpdateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UpdateTicket(ctx, req.(*UpdateTicketHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_DeleteTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteTicket(ctx, req.(*DeleteTicketHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetTicket(ctx, req.(*GetTicketHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketsHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ListTickets(ctx, req.(*ListTicketsHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ViewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ViewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ViewTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ViewTicket(ctx, req.(*ViewTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_LikeTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).LikeTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_LikeTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).LikeTicket(ctx, req.(*LikeTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _UsersService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _UsersService_SignIn_Handler,
		},
		{
			MethodName: "EditInfo",
			Handler:    _UsersService_EditInfo_Handler,
		},
		{
			MethodName: "CreateTicket",
			Handler:    _UsersService_CreateTicket_Handler,
		},
		{
			MethodName: "UpdateTicket",
			Handler:    _UsersService_UpdateTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _UsersService_DeleteTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _UsersService_GetTicket_Handler,
		},
		{
			MethodName: "ListTickets",
			Handler:    _UsersService_ListTickets_Handler,
		},
		{
			MethodName: "ViewTicket",
			Handler:    _UsersService_ViewTicket_Handler,
		},
		{
			MethodName: "LikeTicket",
			Handler:    _UsersService_LikeTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
