// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	session "github.com/Streamfair/common_proto/SessionService/pb/session"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	// Sessions
	CreateSession(ctx context.Context, in *session.CreateSessionRequest, opts ...grpc.CallOption) (*session.CreateSessionResponse, error)
	ExtendSession(ctx context.Context, in *session.ExtendSessionRequest, opts ...grpc.CallOption) (*session.ExtendSessionResponse, error)
	GetSession(ctx context.Context, in *session.GetSessionRequest, opts ...grpc.CallOption) (*session.GetSessionResponse, error)
	DeleteSession(ctx context.Context, in *session.DeleteSessionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeSession(ctx context.Context, in *session.RevokeSessionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VerifySession(ctx context.Context, in *session.VerifySessionRequest, opts ...grpc.CallOption) (*session.VerifySessionResponse, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *session.CreateSessionRequest, opts ...grpc.CallOption) (*session.CreateSessionResponse, error) {
	out := new(session.CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/pb.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) ExtendSession(ctx context.Context, in *session.ExtendSessionRequest, opts ...grpc.CallOption) (*session.ExtendSessionResponse, error) {
	out := new(session.ExtendSessionResponse)
	err := c.cc.Invoke(ctx, "/pb.SessionService/ExtendSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *session.GetSessionRequest, opts ...grpc.CallOption) (*session.GetSessionResponse, error) {
	out := new(session.GetSessionResponse)
	err := c.cc.Invoke(ctx, "/pb.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteSession(ctx context.Context, in *session.DeleteSessionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.SessionService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) RevokeSession(ctx context.Context, in *session.RevokeSessionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.SessionService/RevokeSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) VerifySession(ctx context.Context, in *session.VerifySessionRequest, opts ...grpc.CallOption) (*session.VerifySessionResponse, error) {
	out := new(session.VerifySessionResponse)
	err := c.cc.Invoke(ctx, "/pb.SessionService/VerifySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	// Sessions
	CreateSession(context.Context, *session.CreateSessionRequest) (*session.CreateSessionResponse, error)
	ExtendSession(context.Context, *session.ExtendSessionRequest) (*session.ExtendSessionResponse, error)
	GetSession(context.Context, *session.GetSessionRequest) (*session.GetSessionResponse, error)
	DeleteSession(context.Context, *session.DeleteSessionRequest) (*emptypb.Empty, error)
	RevokeSession(context.Context, *session.RevokeSessionRequest) (*emptypb.Empty, error)
	VerifySession(context.Context, *session.VerifySessionRequest) (*session.VerifySessionResponse, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) CreateSession(context.Context, *session.CreateSessionRequest) (*session.CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedSessionServiceServer) ExtendSession(context.Context, *session.ExtendSessionRequest) (*session.ExtendSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtendSession not implemented")
}
func (UnimplementedSessionServiceServer) GetSession(context.Context, *session.GetSessionRequest) (*session.GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedSessionServiceServer) DeleteSession(context.Context, *session.DeleteSessionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedSessionServiceServer) RevokeSession(context.Context, *session.RevokeSessionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSession not implemented")
}
func (UnimplementedSessionServiceServer) VerifySession(context.Context, *session.VerifySessionRequest) (*session.VerifySessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySession not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(session.CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*session.CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_ExtendSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(session.ExtendSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).ExtendSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SessionService/ExtendSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).ExtendSession(ctx, req.(*session.ExtendSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(session.GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*session.GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(session.DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SessionService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteSession(ctx, req.(*session.DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(session.RevokeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SessionService/RevokeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).RevokeSession(ctx, req.(*session.RevokeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_VerifySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(session.VerifySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).VerifySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SessionService/VerifySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).VerifySession(ctx, req.(*session.VerifySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "ExtendSession",
			Handler:    _SessionService_ExtendSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _SessionService_DeleteSession_Handler,
		},
		{
			MethodName: "RevokeSession",
			Handler:    _SessionService_RevokeSession_Handler,
		},
		{
			MethodName: "VerifySession",
			Handler:    _SessionService_VerifySession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session_svc.proto",
}
