// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api/tickenator.proto

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
	TickenatorService_CreateTicket_FullMethodName = "/api.TickenatorService/CreateTicket"
	TickenatorService_UpdateTicket_FullMethodName = "/api.TickenatorService/UpdateTicket"
	TickenatorService_DeleteTicket_FullMethodName = "/api.TickenatorService/DeleteTicket"
	TickenatorService_GetTicket_FullMethodName    = "/api.TickenatorService/GetTicket"
	TickenatorService_ListTickets_FullMethodName  = "/api.TickenatorService/ListTickets"
)

// TickenatorServiceClient is the client API for TickenatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TickenatorServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*tickenator.Ticket, error)
	ListTickets(ctx context.Context, in *ListTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error)
}

type tickenatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTickenatorServiceClient(cc grpc.ClientConnInterface) TickenatorServiceClient {
	return &tickenatorServiceClient{cc}
}

func (c *tickenatorServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, TickenatorService_CreateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tickenatorServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error) {
	out := new(UpdateTicketResponse)
	err := c.cc.Invoke(ctx, TickenatorService_UpdateTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tickenatorServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, TickenatorService_DeleteTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tickenatorServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*tickenator.Ticket, error) {
	out := new(tickenator.Ticket)
	err := c.cc.Invoke(ctx, TickenatorService_GetTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tickenatorServiceClient) ListTickets(ctx context.Context, in *ListTicketsRequest, opts ...grpc.CallOption) (*ListTicketsResponse, error) {
	out := new(ListTicketsResponse)
	err := c.cc.Invoke(ctx, TickenatorService_ListTickets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TickenatorServiceServer is the server API for TickenatorService service.
// All implementations must embed UnimplementedTickenatorServiceServer
// for forward compatibility
type TickenatorServiceServer interface {
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error)
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*tickenator.Ticket, error)
	ListTickets(context.Context, *ListTicketsRequest) (*ListTicketsResponse, error)
	mustEmbedUnimplementedTickenatorServiceServer()
}

// UnimplementedTickenatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTickenatorServiceServer struct {
}

func (UnimplementedTickenatorServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTickenatorServiceServer) UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
}
func (UnimplementedTickenatorServiceServer) DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTickenatorServiceServer) GetTicket(context.Context, *GetTicketRequest) (*tickenator.Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedTickenatorServiceServer) ListTickets(context.Context, *ListTicketsRequest) (*ListTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}
func (UnimplementedTickenatorServiceServer) mustEmbedUnimplementedTickenatorServiceServer() {}

// UnsafeTickenatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TickenatorServiceServer will
// result in compilation errors.
type UnsafeTickenatorServiceServer interface {
	mustEmbedUnimplementedTickenatorServiceServer()
}

func RegisterTickenatorServiceServer(s grpc.ServiceRegistrar, srv TickenatorServiceServer) {
	s.RegisterService(&TickenatorService_ServiceDesc, srv)
}

func _TickenatorService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickenatorServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickenatorService_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickenatorServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TickenatorService_UpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickenatorServiceServer).UpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickenatorService_UpdateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickenatorServiceServer).UpdateTicket(ctx, req.(*UpdateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TickenatorService_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickenatorServiceServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickenatorService_DeleteTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickenatorServiceServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TickenatorService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickenatorServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickenatorService_GetTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickenatorServiceServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TickenatorService_ListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickenatorServiceServer).ListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickenatorService_ListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickenatorServiceServer).ListTickets(ctx, req.(*ListTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TickenatorService_ServiceDesc is the grpc.ServiceDesc for TickenatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TickenatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TickenatorService",
	HandlerType: (*TickenatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _TickenatorService_CreateTicket_Handler,
		},
		{
			MethodName: "UpdateTicket",
			Handler:    _TickenatorService_UpdateTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TickenatorService_DeleteTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _TickenatorService_GetTicket_Handler,
		},
		{
			MethodName: "ListTickets",
			Handler:    _TickenatorService_ListTickets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tickenator.proto",
}