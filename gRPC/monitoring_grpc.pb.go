// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package monitoring

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

// ElasticSearchMonitoringClient is the client API for ElasticSearchMonitoring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElasticSearchMonitoringClient interface {
	// get cluister's status
	GetStatus(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Status, error)
}

type elasticSearchMonitoringClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticSearchMonitoringClient(cc grpc.ClientConnInterface) ElasticSearchMonitoringClient {
	return &elasticSearchMonitoringClient{cc}
}

func (c *elasticSearchMonitoringClient) GetStatus(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/monitoring.ElasticSearchMonitoring/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticSearchMonitoringServer is the server API for ElasticSearchMonitoring service.
// All implementations must embed UnimplementedElasticSearchMonitoringServer
// for forward compatibility
type ElasticSearchMonitoringServer interface {
	// get cluister's status
	GetStatus(context.Context, *Cluster) (*Status, error)
	mustEmbedUnimplementedElasticSearchMonitoringServer()
}

// UnimplementedElasticSearchMonitoringServer must be embedded to have forward compatible implementations.
type UnimplementedElasticSearchMonitoringServer struct {
}

func (UnimplementedElasticSearchMonitoringServer) GetStatus(context.Context, *Cluster) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedElasticSearchMonitoringServer) mustEmbedUnimplementedElasticSearchMonitoringServer() {
}

// UnsafeElasticSearchMonitoringServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElasticSearchMonitoringServer will
// result in compilation errors.
type UnsafeElasticSearchMonitoringServer interface {
	mustEmbedUnimplementedElasticSearchMonitoringServer()
}

func RegisterElasticSearchMonitoringServer(s grpc.ServiceRegistrar, srv ElasticSearchMonitoringServer) {
	s.RegisterService(&ElasticSearchMonitoring_ServiceDesc, srv)
}

func _ElasticSearchMonitoring_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticSearchMonitoringServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.ElasticSearchMonitoring/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticSearchMonitoringServer).GetStatus(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

// ElasticSearchMonitoring_ServiceDesc is the grpc.ServiceDesc for ElasticSearchMonitoring service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElasticSearchMonitoring_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.ElasticSearchMonitoring",
	HandlerType: (*ElasticSearchMonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _ElasticSearchMonitoring_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/monitoring.proto",
}