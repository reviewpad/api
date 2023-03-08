// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

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

// HostsClient is the client API for Hosts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostsClient interface {
	// Comments
	PostGeneralComment(ctx context.Context, in *PostGeneralCommentRequest, opts ...grpc.CallOption) (*PostGeneralCommentReply, error)
}

type hostsClient struct {
	cc grpc.ClientConnInterface
}

func NewHostsClient(cc grpc.ClientConnInterface) HostsClient {
	return &hostsClient{cc}
}

func (c *hostsClient) PostGeneralComment(ctx context.Context, in *PostGeneralCommentRequest, opts ...grpc.CallOption) (*PostGeneralCommentReply, error) {
	out := new(PostGeneralCommentReply)
	err := c.cc.Invoke(ctx, "/services.Hosts/PostGeneralComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostsServer is the server API for Hosts service.
// All implementations must embed UnimplementedHostsServer
// for forward compatibility
type HostsServer interface {
	// Comments
	PostGeneralComment(context.Context, *PostGeneralCommentRequest) (*PostGeneralCommentReply, error)
	mustEmbedUnimplementedHostsServer()
}

// UnimplementedHostsServer must be embedded to have forward compatible implementations.
type UnimplementedHostsServer struct {
}

func (UnimplementedHostsServer) PostGeneralComment(context.Context, *PostGeneralCommentRequest) (*PostGeneralCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGeneralComment not implemented")
}
func (UnimplementedHostsServer) mustEmbedUnimplementedHostsServer() {}

// UnsafeHostsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostsServer will
// result in compilation errors.
type UnsafeHostsServer interface {
	mustEmbedUnimplementedHostsServer()
}

func RegisterHostsServer(s grpc.ServiceRegistrar, srv HostsServer) {
	s.RegisterService(&Hosts_ServiceDesc, srv)
}

func _Hosts_PostGeneralComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostGeneralCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServer).PostGeneralComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Hosts/PostGeneralComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServer).PostGeneralComment(ctx, req.(*PostGeneralCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hosts_ServiceDesc is the grpc.ServiceDesc for Hosts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hosts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Hosts",
	HandlerType: (*HostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostGeneralComment",
			Handler:    _Hosts_PostGeneralComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/hosts.proto",
}
