// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: llm.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LLMService_GenerateText_FullMethodName      = "/llm.LLMService/GenerateText"
	LLMService_ClearConversation_FullMethodName = "/llm.LLMService/ClearConversation"
)

// LLMServiceClient is the client API for LLMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LLMServiceClient interface {
	GenerateText(ctx context.Context, in *ChatHistoryRequest, opts ...grpc.CallOption) (*TextResponse, error)
	ClearConversation(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type lLMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLLMServiceClient(cc grpc.ClientConnInterface) LLMServiceClient {
	return &lLMServiceClient{cc}
}

func (c *lLMServiceClient) GenerateText(ctx context.Context, in *ChatHistoryRequest, opts ...grpc.CallOption) (*TextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextResponse)
	err := c.cc.Invoke(ctx, LLMService_GenerateText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMServiceClient) ClearConversation(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LLMService_ClearConversation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LLMServiceServer is the server API for LLMService service.
// All implementations must embed UnimplementedLLMServiceServer
// for forward compatibility.
type LLMServiceServer interface {
	GenerateText(context.Context, *ChatHistoryRequest) (*TextResponse, error)
	ClearConversation(context.Context, *User) (*emptypb.Empty, error)
	mustEmbedUnimplementedLLMServiceServer()
}

// UnimplementedLLMServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLLMServiceServer struct{}

func (UnimplementedLLMServiceServer) GenerateText(context.Context, *ChatHistoryRequest) (*TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateText not implemented")
}
func (UnimplementedLLMServiceServer) ClearConversation(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearConversation not implemented")
}
func (UnimplementedLLMServiceServer) mustEmbedUnimplementedLLMServiceServer() {}
func (UnimplementedLLMServiceServer) testEmbeddedByValue()                    {}

// UnsafeLLMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LLMServiceServer will
// result in compilation errors.
type UnsafeLLMServiceServer interface {
	mustEmbedUnimplementedLLMServiceServer()
}

func RegisterLLMServiceServer(s grpc.ServiceRegistrar, srv LLMServiceServer) {
	// If the following call pancis, it indicates UnimplementedLLMServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LLMService_ServiceDesc, srv)
}

func _LLMService_GenerateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).GenerateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_GenerateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).GenerateText(ctx, req.(*ChatHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLMService_ClearConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).ClearConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_ClearConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).ClearConversation(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// LLMService_ServiceDesc is the grpc.ServiceDesc for LLMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LLMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llm.LLMService",
	HandlerType: (*LLMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateText",
			Handler:    _LLMService_GenerateText_Handler,
		},
		{
			MethodName: "ClearConversation",
			Handler:    _LLMService_ClearConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "llm.proto",
}
