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

// SemanticClient is the client API for Semantic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SemanticClient interface {
	GetSymbols(ctx context.Context, in *GetSymbolsRequest, opts ...grpc.CallOption) (*GetSymbolsReply, error)
}

type semanticClient struct {
	cc grpc.ClientConnInterface
}

func NewSemanticClient(cc grpc.ClientConnInterface) SemanticClient {
	return &semanticClient{cc}
}

func (c *semanticClient) GetSymbols(ctx context.Context, in *GetSymbolsRequest, opts ...grpc.CallOption) (*GetSymbolsReply, error) {
	out := new(GetSymbolsReply)
	err := c.cc.Invoke(ctx, "/services.Semantic/GetSymbols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SemanticServer is the server API for Semantic service.
// All implementations must embed UnimplementedSemanticServer
// for forward compatibility
type SemanticServer interface {
	GetSymbols(context.Context, *GetSymbolsRequest) (*GetSymbolsReply, error)
	mustEmbedUnimplementedSemanticServer()
}

// UnimplementedSemanticServer must be embedded to have forward compatible implementations.
type UnimplementedSemanticServer struct {
}

func (UnimplementedSemanticServer) GetSymbols(context.Context, *GetSymbolsRequest) (*GetSymbolsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbols not implemented")
}
func (UnimplementedSemanticServer) mustEmbedUnimplementedSemanticServer() {}

// UnsafeSemanticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SemanticServer will
// result in compilation errors.
type UnsafeSemanticServer interface {
	mustEmbedUnimplementedSemanticServer()
}

func RegisterSemanticServer(s grpc.ServiceRegistrar, srv SemanticServer) {
	s.RegisterService(&Semantic_ServiceDesc, srv)
}

func _Semantic_GetSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SemanticServer).GetSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Semantic/GetSymbols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SemanticServer).GetSymbols(ctx, req.(*GetSymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Semantic_ServiceDesc is the grpc.ServiceDesc for Semantic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Semantic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Semantic",
	HandlerType: (*SemanticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSymbols",
			Handler:    _Semantic_GetSymbols_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/semantic.proto",
}
