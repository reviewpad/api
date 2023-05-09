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

// RobinClient is the client API for Robin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobinClient interface {
	Explain(ctx context.Context, in *ExplainRequest, opts ...grpc.CallOption) (*ExplainReply, error)
	Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptReply, error)
	RawPrompt(ctx context.Context, in *RawPromptRequest, opts ...grpc.CallOption) (*PromptReply, error)
	Summarize(ctx context.Context, in *SummarizeRequest, opts ...grpc.CallOption) (*SummarizeReply, error)
}

type robinClient struct {
	cc grpc.ClientConnInterface
}

func NewRobinClient(cc grpc.ClientConnInterface) RobinClient {
	return &robinClient{cc}
}

func (c *robinClient) Explain(ctx context.Context, in *ExplainRequest, opts ...grpc.CallOption) (*ExplainReply, error) {
	out := new(ExplainReply)
	err := c.cc.Invoke(ctx, "/services.Robin/Explain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robinClient) Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptReply, error) {
	out := new(PromptReply)
	err := c.cc.Invoke(ctx, "/services.Robin/Prompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robinClient) RawPrompt(ctx context.Context, in *RawPromptRequest, opts ...grpc.CallOption) (*PromptReply, error) {
	out := new(PromptReply)
	err := c.cc.Invoke(ctx, "/services.Robin/RawPrompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robinClient) Summarize(ctx context.Context, in *SummarizeRequest, opts ...grpc.CallOption) (*SummarizeReply, error) {
	out := new(SummarizeReply)
	err := c.cc.Invoke(ctx, "/services.Robin/Summarize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobinServer is the server API for Robin service.
// All implementations must embed UnimplementedRobinServer
// for forward compatibility
type RobinServer interface {
	Explain(context.Context, *ExplainRequest) (*ExplainReply, error)
	Prompt(context.Context, *PromptRequest) (*PromptReply, error)
	RawPrompt(context.Context, *RawPromptRequest) (*PromptReply, error)
	Summarize(context.Context, *SummarizeRequest) (*SummarizeReply, error)
	mustEmbedUnimplementedRobinServer()
}

// UnimplementedRobinServer must be embedded to have forward compatible implementations.
type UnimplementedRobinServer struct {
}

func (UnimplementedRobinServer) Explain(context.Context, *ExplainRequest) (*ExplainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Explain not implemented")
}
func (UnimplementedRobinServer) Prompt(context.Context, *PromptRequest) (*PromptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prompt not implemented")
}
func (UnimplementedRobinServer) RawPrompt(context.Context, *RawPromptRequest) (*PromptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawPrompt not implemented")
}
func (UnimplementedRobinServer) Summarize(context.Context, *SummarizeRequest) (*SummarizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Summarize not implemented")
}
func (UnimplementedRobinServer) mustEmbedUnimplementedRobinServer() {}

// UnsafeRobinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobinServer will
// result in compilation errors.
type UnsafeRobinServer interface {
	mustEmbedUnimplementedRobinServer()
}

func RegisterRobinServer(s grpc.ServiceRegistrar, srv RobinServer) {
	s.RegisterService(&Robin_ServiceDesc, srv)
}

func _Robin_Explain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExplainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobinServer).Explain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Robin/Explain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobinServer).Explain(ctx, req.(*ExplainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robin_Prompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobinServer).Prompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Robin/Prompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobinServer).Prompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robin_RawPrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawPromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobinServer).RawPrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Robin/RawPrompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobinServer).RawPrompt(ctx, req.(*RawPromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robin_Summarize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummarizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobinServer).Summarize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Robin/Summarize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobinServer).Summarize(ctx, req.(*SummarizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Robin_ServiceDesc is the grpc.ServiceDesc for Robin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Robin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Robin",
	HandlerType: (*RobinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Explain",
			Handler:    _Robin_Explain_Handler,
		},
		{
			MethodName: "Prompt",
			Handler:    _Robin_Prompt_Handler,
		},
		{
			MethodName: "RawPrompt",
			Handler:    _Robin_RawPrompt_Handler,
		},
		{
			MethodName: "Summarize",
			Handler:    _Robin_Summarize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/robin.proto",
}
