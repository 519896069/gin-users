// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: app/services/user.proto

package pb

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

// UserServicesClient is the client API for UserServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServicesClient interface {
	//用token获取用户 - 用户鉴权
	GetUserByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error)
}

type userServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServicesClient(cc grpc.ClientConnInterface) UserServicesClient {
	return &userServicesClient{cc}
}

func (c *userServicesClient) GetUserByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserServices/GetUserByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServicesServer is the server API for UserServices service.
// All implementations must embed UnimplementedUserServicesServer
// for forward compatibility
type UserServicesServer interface {
	//用token获取用户 - 用户鉴权
	GetUserByToken(context.Context, *Token) (*User, error)
	mustEmbedUnimplementedUserServicesServer()
}

// UnimplementedUserServicesServer must be embedded to have forward compatible implementations.
type UnimplementedUserServicesServer struct {
}

func (UnimplementedUserServicesServer) GetUserByToken(context.Context, *Token) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByToken not implemented")
}
func (UnimplementedUserServicesServer) mustEmbedUnimplementedUserServicesServer() {}

// UnsafeUserServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServicesServer will
// result in compilation errors.
type UnsafeUserServicesServer interface {
	mustEmbedUnimplementedUserServicesServer()
}

func RegisterUserServicesServer(s grpc.ServiceRegistrar, srv UserServicesServer) {
	s.RegisterService(&UserServices_ServiceDesc, srv)
}

func _UserServices_GetUserByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).GetUserByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServices/GetUserByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).GetUserByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServices_ServiceDesc is the grpc.ServiceDesc for UserServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserServices",
	HandlerType: (*UserServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByToken",
			Handler:    _UserServices_GetUserByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/services/user.proto",
}
