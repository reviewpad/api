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

// DiffClient is the client API for Diff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiffClient interface {
	GetDefinedSymbolsDiff(ctx context.Context, in *GetDefinedSymbolsDiffRequest, opts ...grpc.CallOption) (*GetDefinedSymbolsDiffReply, error)
}

type diffClient struct {
	cc grpc.ClientConnInterface
}

func NewDiffClient(cc grpc.ClientConnInterface) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) GetDefinedSymbolsDiff(ctx context.Context, in *GetDefinedSymbolsDiffRequest, opts ...grpc.CallOption) (*GetDefinedSymbolsDiffReply, error) {
	out := new(GetDefinedSymbolsDiffReply)
	err := c.cc.Invoke(ctx, "/services.Diff/GetDefinedSymbolsDiff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServer is the server API for Diff service.
// All implementations must embed UnimplementedDiffServer
// for forward compatibility
type DiffServer interface {
	GetDefinedSymbolsDiff(context.Context, *GetDefinedSymbolsDiffRequest) (*GetDefinedSymbolsDiffReply, error)
	mustEmbedUnimplementedDiffServer()
}

// UnimplementedDiffServer must be embedded to have forward compatible implementations.
type UnimplementedDiffServer struct {
}

func (UnimplementedDiffServer) GetDefinedSymbolsDiff(context.Context, *GetDefinedSymbolsDiffRequest) (*GetDefinedSymbolsDiffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefinedSymbolsDiff not implemented")
}
func (UnimplementedDiffServer) mustEmbedUnimplementedDiffServer() {}

// UnsafeDiffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiffServer will
// result in compilation errors.
type UnsafeDiffServer interface {
	mustEmbedUnimplementedDiffServer()
}

func RegisterDiffServer(s grpc.ServiceRegistrar, srv DiffServer) {
	s.RegisterService(&Diff_ServiceDesc, srv)
}

func _Diff_GetDefinedSymbolsDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefinedSymbolsDiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).GetDefinedSymbolsDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Diff/GetDefinedSymbolsDiff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).GetDefinedSymbolsDiff(ctx, req.(*GetDefinedSymbolsDiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Diff_ServiceDesc is the grpc.ServiceDesc for Diff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDefinedSymbolsDiff",
			Handler:    _Diff_GetDefinedSymbolsDiff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/diff.proto",
}