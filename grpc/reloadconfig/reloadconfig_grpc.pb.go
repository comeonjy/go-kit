// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reloadconfig

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

// ReloadConfigClient is the client API for ReloadConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReloadConfigClient interface {
	ReloadConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type reloadConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewReloadConfigClient(cc grpc.ClientConnInterface) ReloadConfigClient {
	return &reloadConfigClient{cc}
}

func (c *reloadConfigClient) ReloadConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/reloadconfig.ReloadConfig/ReloadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReloadConfigServer is the server API for ReloadConfig service.
// All implementations must embed UnimplementedReloadConfigServer
// for forward compatibility
type ReloadConfigServer interface {
	ReloadConfig(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedReloadConfigServer()
}

// UnimplementedReloadConfigServer must be embedded to have forward compatible implementations.
type UnimplementedReloadConfigServer struct {
}

func (UnimplementedReloadConfigServer) ReloadConfig(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadConfig not implemented")
}
func (UnimplementedReloadConfigServer) mustEmbedUnimplementedReloadConfigServer() {}

// UnsafeReloadConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReloadConfigServer will
// result in compilation errors.
type UnsafeReloadConfigServer interface {
	mustEmbedUnimplementedReloadConfigServer()
}

func RegisterReloadConfigServer(s grpc.ServiceRegistrar, srv ReloadConfigServer) {
	s.RegisterService(&ReloadConfig_ServiceDesc, srv)
}

func _ReloadConfig_ReloadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReloadConfigServer).ReloadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reloadconfig.ReloadConfig/ReloadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReloadConfigServer).ReloadConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ReloadConfig_ServiceDesc is the grpc.ServiceDesc for ReloadConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReloadConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reloadconfig.ReloadConfig",
	HandlerType: (*ReloadConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReloadConfig",
			Handler:    _ReloadConfig_ReloadConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/reloadconfig/reloadconfig.proto",
}
