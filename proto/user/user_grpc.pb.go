// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// FetchClient is the client API for Fetch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchClient interface {
	// Fetch one user details, given the user id
	FetchUserByID(ctx context.Context, in *FetchUserByIDRequest, opts ...grpc.CallOption) (*FetchUserByIDResponse, error)
	// Fetch multiple user details, give multiple user ids
	FetchUsers(ctx context.Context, in *FetchUsersRequest, opts ...grpc.CallOption) (*FetchUsersResponse, error)
}

type fetchClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchClient(cc grpc.ClientConnInterface) FetchClient {
	return &fetchClient{cc}
}

func (c *fetchClient) FetchUserByID(ctx context.Context, in *FetchUserByIDRequest, opts ...grpc.CallOption) (*FetchUserByIDResponse, error) {
	out := new(FetchUserByIDResponse)
	err := c.cc.Invoke(ctx, "/user.Fetch/FetchUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) FetchUsers(ctx context.Context, in *FetchUsersRequest, opts ...grpc.CallOption) (*FetchUsersResponse, error) {
	out := new(FetchUsersResponse)
	err := c.cc.Invoke(ctx, "/user.Fetch/FetchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchServer is the server API for Fetch service.
// All implementations must embed UnimplementedFetchServer
// for forward compatibility
type FetchServer interface {
	// Fetch one user details, given the user id
	FetchUserByID(context.Context, *FetchUserByIDRequest) (*FetchUserByIDResponse, error)
	// Fetch multiple user details, give multiple user ids
	FetchUsers(context.Context, *FetchUsersRequest) (*FetchUsersResponse, error)
	mustEmbedUnimplementedFetchServer()
}

// UnimplementedFetchServer must be embedded to have forward compatible implementations.
type UnimplementedFetchServer struct {
}

func (UnimplementedFetchServer) FetchUserByID(context.Context, *FetchUserByIDRequest) (*FetchUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUserByID not implemented")
}
func (UnimplementedFetchServer) FetchUsers(context.Context, *FetchUsersRequest) (*FetchUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUsers not implemented")
}
func (UnimplementedFetchServer) mustEmbedUnimplementedFetchServer() {}

// UnsafeFetchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchServer will
// result in compilation errors.
type UnsafeFetchServer interface {
	mustEmbedUnimplementedFetchServer()
}

func RegisterFetchServer(s grpc.ServiceRegistrar, srv FetchServer) {
	s.RegisterService(&Fetch_ServiceDesc, srv)
}

func _Fetch_FetchUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Fetch/FetchUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchUserByID(ctx, req.(*FetchUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_FetchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Fetch/FetchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchUsers(ctx, req.(*FetchUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fetch_ServiceDesc is the grpc.ServiceDesc for Fetch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fetch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Fetch",
	HandlerType: (*FetchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchUserByID",
			Handler:    _Fetch_FetchUserByID_Handler,
		},
		{
			MethodName: "FetchUsers",
			Handler:    _Fetch_FetchUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
