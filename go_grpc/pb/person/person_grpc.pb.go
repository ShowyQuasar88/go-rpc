// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: person/person.proto

package person

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
	SearchService_Search_FullMethodName    = "/person.SearchService/Search"
	SearchService_SearchIn_FullMethodName  = "/person.SearchService/SearchIn"
	SearchService_SearchOut_FullMethodName = "/person.SearchService/SearchOut"
	SearchService_SearchIO_FullMethodName  = "/person.SearchService/SearchIO"
)

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonResp, error)
	SearchIn(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PersonReq, PersonResp], error)
	SearchOut(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PersonResp], error)
	SearchIO(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PersonReq, PersonResp], error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonResp)
	err := c.cc.Invoke(ctx, SearchService_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchIn(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PersonReq, PersonResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[0], SearchService_SearchIn_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PersonReq, PersonResp]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SearchService_SearchInClient = grpc.ClientStreamingClient[PersonReq, PersonResp]

func (c *searchServiceClient) SearchOut(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PersonResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[1], SearchService_SearchOut_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PersonReq, PersonResp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SearchService_SearchOutClient = grpc.ServerStreamingClient[PersonResp]

func (c *searchServiceClient) SearchIO(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PersonReq, PersonResp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[2], SearchService_SearchIO_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PersonReq, PersonResp]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SearchService_SearchIOClient = grpc.BidiStreamingClient[PersonReq, PersonResp]

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility.
type SearchServiceServer interface {
	Search(context.Context, *PersonReq) (*PersonResp, error)
	SearchIn(grpc.ClientStreamingServer[PersonReq, PersonResp]) error
	SearchOut(*PersonReq, grpc.ServerStreamingServer[PersonResp]) error
	SearchIO(grpc.BidiStreamingServer[PersonReq, PersonResp]) error
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchServiceServer struct{}

func (UnimplementedSearchServiceServer) Search(context.Context, *PersonReq) (*PersonResp, error) {

	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearchServiceServer) SearchIn(grpc.ClientStreamingServer[PersonReq, PersonResp]) error {
	return status.Errorf(codes.Unimplemented, "method SearchIn not implemented")
}
func (UnimplementedSearchServiceServer) SearchOut(*PersonReq, grpc.ServerStreamingServer[PersonResp]) error {
	return status.Errorf(codes.Unimplemented, "method SearchOut not implemented")
}
func (UnimplementedSearchServiceServer) SearchIO(grpc.BidiStreamingServer[PersonReq, PersonResp]) error {
	return status.Errorf(codes.Unimplemented, "method SearchIO not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}
func (UnimplementedSearchServiceServer) testEmbeddedByValue()                       {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	// If the following call pancis, it indicates UnimplementedSearchServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*PersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchServiceServer).SearchIn(&grpc.GenericServerStream[PersonReq, PersonResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SearchService_SearchInServer = grpc.ClientStreamingServer[PersonReq, PersonResp]

func _SearchService_SearchOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PersonReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServiceServer).SearchOut(m, &grpc.GenericServerStream[PersonReq, PersonResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SearchService_SearchOutServer = grpc.ServerStreamingServer[PersonResp]

func _SearchService_SearchIO_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchServiceServer).SearchIO(&grpc.GenericServerStream[PersonReq, PersonResp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SearchService_SearchIOServer = grpc.BidiStreamingServer[PersonReq, PersonResp]

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchIn",
			Handler:       _SearchService_SearchIn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SearchOut",
			Handler:       _SearchService_SearchOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchIO",
			Handler:       _SearchService_SearchIO_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "person/person.proto",
}