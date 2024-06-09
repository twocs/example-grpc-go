// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: example-grpc-go/example-grpc-go.proto

package example_grpc_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GetUsersClient is the client API for GetUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUsersClient interface {
	// A simple RPC.
	//
	// Obtains the user with a given id.
	//
	// A user with an empty name is returned if there's no user at the given
	// position.
	GetUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*User, error)
	// A list of users
	//
	// Obtains the list of users corresponding to a list of ids. Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListUsers(ctx context.Context, in *IDs, opts ...grpc.CallOption) (GetUsers_ListUsersClient, error)
	// search endpoint
	//
	// Obtains the list of users corresponding to a search. Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	SearchUsers(ctx context.Context, in *Query, opts ...grpc.CallOption) (GetUsers_SearchUsersClient, error)
}

type getUsersClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUsersClient(cc grpc.ClientConnInterface) GetUsersClient {
	return &getUsersClient{cc}
}

func (c *getUsersClient) GetUser(ctx context.Context, in *ID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/example_grpc_go.GetUsers/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getUsersClient) ListUsers(ctx context.Context, in *IDs, opts ...grpc.CallOption) (GetUsers_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &GetUsers_ServiceDesc.Streams[0], "/example_grpc_go.GetUsers/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &getUsersListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GetUsers_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type getUsersListUsersClient struct {
	grpc.ClientStream
}

func (x *getUsersListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getUsersClient) SearchUsers(ctx context.Context, in *Query, opts ...grpc.CallOption) (GetUsers_SearchUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &GetUsers_ServiceDesc.Streams[1], "/example_grpc_go.GetUsers/SearchUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &getUsersSearchUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GetUsers_SearchUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type getUsersSearchUsersClient struct {
	grpc.ClientStream
}

func (x *getUsersSearchUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GetUsersServer is the server API for GetUsers service.
// All implementations must embed UnimplementedGetUsersServer
// for forward compatibility
type GetUsersServer interface {
	// A simple RPC.
	//
	// Obtains the user with a given id.
	//
	// A user with an empty name is returned if there's no user at the given
	// position.
	GetUser(context.Context, *ID) (*User, error)
	// A list of users
	//
	// Obtains the list of users corresponding to a list of ids. Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListUsers(*IDs, GetUsers_ListUsersServer) error
	// search endpoint
	//
	// Obtains the list of users corresponding to a search. Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	SearchUsers(*Query, GetUsers_SearchUsersServer) error
	mustEmbedUnimplementedGetUsersServer()
}

// UnimplementedGetUsersServer must be embedded to have forward compatible implementations.
type UnimplementedGetUsersServer struct {
}

func (UnimplementedGetUsersServer) GetUser(context.Context, *ID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGetUsersServer) ListUsers(*IDs, GetUsers_ListUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedGetUsersServer) SearchUsers(*Query, GetUsers_SearchUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedGetUsersServer) mustEmbedUnimplementedGetUsersServer() {}

// UnsafeGetUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUsersServer will
// result in compilation errors.
type UnsafeGetUsersServer interface {
	mustEmbedUnimplementedGetUsersServer()
}

func RegisterGetUsersServer(s grpc.ServiceRegistrar, srv GetUsersServer) {
	s.RegisterService(&GetUsers_ServiceDesc, srv)
}

func _GetUsers_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example_grpc_go.GetUsers/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUsersServer).GetUser(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetUsers_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IDs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GetUsersServer).ListUsers(m, &getUsersListUsersServer{stream})
}

type GetUsers_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type getUsersListUsersServer struct {
	grpc.ServerStream
}

func (x *getUsersListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _GetUsers_SearchUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GetUsersServer).SearchUsers(m, &getUsersSearchUsersServer{stream})
}

type GetUsers_SearchUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type getUsersSearchUsersServer struct {
	grpc.ServerStream
}

func (x *getUsersSearchUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

// GetUsers_ServiceDesc is the grpc.ServiceDesc for GetUsers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUsers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example_grpc_go.GetUsers",
	HandlerType: (*GetUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _GetUsers_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _GetUsers_ListUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchUsers",
			Handler:       _GetUsers_SearchUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example-grpc-go/example-grpc-go.proto",
}
