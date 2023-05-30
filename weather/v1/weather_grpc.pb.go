// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: weather/v1/weather.proto

package weatherv1

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

const (
	WeatherVisionService_CurrentWeather_FullMethodName  = "/weather.v1.WeatherVisionService/CurrentWeather"
	WeatherVisionService_ExpectedWeather_FullMethodName = "/weather.v1.WeatherVisionService/ExpectedWeather"
)

// WeatherVisionServiceClient is the client API for WeatherVisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherVisionServiceClient interface {
	CurrentWeather(ctx context.Context, in *CurrentWeatherRequest, opts ...grpc.CallOption) (*CurrentWeatherResponse, error)
	ExpectedWeather(ctx context.Context, in *ExpectedWeatherRequest, opts ...grpc.CallOption) (*ExpectedWeatherResponse, error)
}

type weatherVisionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherVisionServiceClient(cc grpc.ClientConnInterface) WeatherVisionServiceClient {
	return &weatherVisionServiceClient{cc}
}

func (c *weatherVisionServiceClient) CurrentWeather(ctx context.Context, in *CurrentWeatherRequest, opts ...grpc.CallOption) (*CurrentWeatherResponse, error) {
	out := new(CurrentWeatherResponse)
	err := c.cc.Invoke(ctx, WeatherVisionService_CurrentWeather_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherVisionServiceClient) ExpectedWeather(ctx context.Context, in *ExpectedWeatherRequest, opts ...grpc.CallOption) (*ExpectedWeatherResponse, error) {
	out := new(ExpectedWeatherResponse)
	err := c.cc.Invoke(ctx, WeatherVisionService_ExpectedWeather_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherVisionServiceServer is the server API for WeatherVisionService service.
// All implementations must embed UnimplementedWeatherVisionServiceServer
// for forward compatibility
type WeatherVisionServiceServer interface {
	CurrentWeather(context.Context, *CurrentWeatherRequest) (*CurrentWeatherResponse, error)
	ExpectedWeather(context.Context, *ExpectedWeatherRequest) (*ExpectedWeatherResponse, error)
	mustEmbedUnimplementedWeatherVisionServiceServer()
}

// UnimplementedWeatherVisionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherVisionServiceServer struct {
}

func (UnimplementedWeatherVisionServiceServer) CurrentWeather(context.Context, *CurrentWeatherRequest) (*CurrentWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentWeather not implemented")
}
func (UnimplementedWeatherVisionServiceServer) ExpectedWeather(context.Context, *ExpectedWeatherRequest) (*ExpectedWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpectedWeather not implemented")
}
func (UnimplementedWeatherVisionServiceServer) mustEmbedUnimplementedWeatherVisionServiceServer() {}

// UnsafeWeatherVisionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherVisionServiceServer will
// result in compilation errors.
type UnsafeWeatherVisionServiceServer interface {
	mustEmbedUnimplementedWeatherVisionServiceServer()
}

func RegisterWeatherVisionServiceServer(s grpc.ServiceRegistrar, srv WeatherVisionServiceServer) {
	s.RegisterService(&WeatherVisionService_ServiceDesc, srv)
}

func _WeatherVisionService_CurrentWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherVisionServiceServer).CurrentWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherVisionService_CurrentWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherVisionServiceServer).CurrentWeather(ctx, req.(*CurrentWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherVisionService_ExpectedWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpectedWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherVisionServiceServer).ExpectedWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeatherVisionService_ExpectedWeather_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherVisionServiceServer).ExpectedWeather(ctx, req.(*ExpectedWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherVisionService_ServiceDesc is the grpc.ServiceDesc for WeatherVisionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherVisionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.v1.WeatherVisionService",
	HandlerType: (*WeatherVisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentWeather",
			Handler:    _WeatherVisionService_CurrentWeather_Handler,
		},
		{
			MethodName: "ExpectedWeather",
			Handler:    _WeatherVisionService_ExpectedWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather/v1/weather.proto",
}
