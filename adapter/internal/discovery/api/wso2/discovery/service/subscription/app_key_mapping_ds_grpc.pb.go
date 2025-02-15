// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package subscription

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ApplicationKeyMappingDiscoveryServiceClient is the client API for ApplicationKeyMappingDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationKeyMappingDiscoveryServiceClient interface {
	StreamApplicationKeyMappings(ctx context.Context, opts ...grpc.CallOption) (ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsClient, error)
}

type applicationKeyMappingDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationKeyMappingDiscoveryServiceClient(cc grpc.ClientConnInterface) ApplicationKeyMappingDiscoveryServiceClient {
	return &applicationKeyMappingDiscoveryServiceClient{cc}
}

func (c *applicationKeyMappingDiscoveryServiceClient) StreamApplicationKeyMappings(ctx context.Context, opts ...grpc.CallOption) (ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationKeyMappingDiscoveryService_serviceDesc.Streams[0], "/discovery.service.subscription.ApplicationKeyMappingDiscoveryService/StreamApplicationKeyMappings", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsClient{stream}
	return x, nil
}

type ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsClient struct {
	grpc.ClientStream
}

func (x *applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApplicationKeyMappingDiscoveryServiceServer is the server API for ApplicationKeyMappingDiscoveryService service.
// All implementations must embed UnimplementedApplicationKeyMappingDiscoveryServiceServer
// for forward compatibility
type ApplicationKeyMappingDiscoveryServiceServer interface {
	StreamApplicationKeyMappings(ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsServer) error
	mustEmbedUnimplementedApplicationKeyMappingDiscoveryServiceServer()
}

// UnimplementedApplicationKeyMappingDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationKeyMappingDiscoveryServiceServer struct {
}

func (UnimplementedApplicationKeyMappingDiscoveryServiceServer) StreamApplicationKeyMappings(ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamApplicationKeyMappings not implemented")
}
func (UnimplementedApplicationKeyMappingDiscoveryServiceServer) mustEmbedUnimplementedApplicationKeyMappingDiscoveryServiceServer() {
}

// UnsafeApplicationKeyMappingDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationKeyMappingDiscoveryServiceServer will
// result in compilation errors.
type UnsafeApplicationKeyMappingDiscoveryServiceServer interface {
	mustEmbedUnimplementedApplicationKeyMappingDiscoveryServiceServer()
}

func RegisterApplicationKeyMappingDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ApplicationKeyMappingDiscoveryServiceServer) {
	s.RegisterService(&_ApplicationKeyMappingDiscoveryService_serviceDesc, srv)
}

func _ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApplicationKeyMappingDiscoveryServiceServer).StreamApplicationKeyMappings(&applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsServer{stream})
}

type ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappingsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsServer struct {
	grpc.ServerStream
}

func (x *applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *applicationKeyMappingDiscoveryServiceStreamApplicationKeyMappingsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ApplicationKeyMappingDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.service.subscription.ApplicationKeyMappingDiscoveryService",
	HandlerType: (*ApplicationKeyMappingDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamApplicationKeyMappings",
			Handler:       _ApplicationKeyMappingDiscoveryService_StreamApplicationKeyMappings_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "wso2/discovery/service/subscription/app_key_mapping_ds.proto",
}
