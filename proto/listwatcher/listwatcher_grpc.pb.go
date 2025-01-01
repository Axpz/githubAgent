// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: listwatcher.proto

package listwatcher

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ListWatchService_ListWatch_FullMethodName = "/listwatcher.ListWatchService/ListWatch"
)

// ListWatchServiceClient is the client API for ListWatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the ListWatch service
type ListWatchServiceClient interface {
	// ListWatch method: streams the current state of all items and updates.
	ListWatch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Event, Event], error)
}

type listWatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListWatchServiceClient(cc grpc.ClientConnInterface) ListWatchServiceClient {
	return &listWatchServiceClient{cc}
}

func (c *listWatchServiceClient) ListWatch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Event, Event], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ListWatchService_ServiceDesc.Streams[0], ListWatchService_ListWatch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Event, Event]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ListWatchService_ListWatchClient = grpc.BidiStreamingClient[Event, Event]

// ListWatchServiceServer is the server API for ListWatchService service.
// All implementations must embed UnimplementedListWatchServiceServer
// for forward compatibility.
//
// Define the ListWatch service
type ListWatchServiceServer interface {
	// ListWatch method: streams the current state of all items and updates.
	ListWatch(grpc.BidiStreamingServer[Event, Event]) error
	mustEmbedUnimplementedListWatchServiceServer()
}

// UnimplementedListWatchServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListWatchServiceServer struct{}

func (UnimplementedListWatchServiceServer) ListWatch(grpc.BidiStreamingServer[Event, Event]) error {
	return status.Errorf(codes.Unimplemented, "method ListWatch not implemented")
}
func (UnimplementedListWatchServiceServer) mustEmbedUnimplementedListWatchServiceServer() {}
func (UnimplementedListWatchServiceServer) testEmbeddedByValue()                          {}

// UnsafeListWatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListWatchServiceServer will
// result in compilation errors.
type UnsafeListWatchServiceServer interface {
	mustEmbedUnimplementedListWatchServiceServer()
}

func RegisterListWatchServiceServer(s grpc.ServiceRegistrar, srv ListWatchServiceServer) {
	// If the following call pancis, it indicates UnimplementedListWatchServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListWatchService_ServiceDesc, srv)
}

func _ListWatchService_ListWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListWatchServiceServer).ListWatch(&grpc.GenericServerStream[Event, Event]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ListWatchService_ListWatchServer = grpc.BidiStreamingServer[Event, Event]

// ListWatchService_ServiceDesc is the grpc.ServiceDesc for ListWatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListWatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listwatcher.ListWatchService",
	HandlerType: (*ListWatchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListWatch",
			Handler:       _ListWatchService_ListWatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "listwatcher.proto",
}
