// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/betting_system.proto

package proto

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
	BettingService_CreatePodiumBet_FullMethodName          = "/betting.BettingService/CreatePodiumBet"
	BettingService_CreatePolePositionBet_FullMethodName    = "/betting.BettingService/CreatePolePositionBet"
	BettingService_CreateRainBet_FullMethodName            = "/betting.BettingService/CreateRainBet"
	BettingService_CreateRetirementBet_FullMethodName      = "/betting.BettingService/CreateRetirementBet"
	BettingService_CreateFastestLapBet_FullMethodName      = "/betting.BettingService/CreateFastestLapBet"
	BettingService_CreateLapTimingBet_FullMethodName       = "/betting.BettingService/CreateLapTimingBet"
	BettingService_GetUserActiveBets_FullMethodName        = "/betting.BettingService/GetUserActiveBets"
	BettingService_GetUserBetHistory_FullMethodName        = "/betting.BettingService/GetUserBetHistory"
	BettingService_GetPendingBetsForSession_FullMethodName = "/betting.BettingService/GetPendingBetsForSession"
	BettingService_GetUserBalance_FullMethodName           = "/betting.BettingService/GetUserBalance"
)

// BettingServiceClient is the client API for BettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BettingServiceClient interface {
	CreatePodiumBet(ctx context.Context, in *PodiumBetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	CreatePolePositionBet(ctx context.Context, in *PolePositionBetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	CreateRainBet(ctx context.Context, in *RainBetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	CreateRetirementBet(ctx context.Context, in *RetirementBetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	CreateFastestLapBet(ctx context.Context, in *FastestLapBetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	CreateLapTimingBet(ctx context.Context, in *LapTimingBetRequest, opts ...grpc.CallOption) (*BetResponse, error)
	GetUserActiveBets(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserBetsResponse, error)
	GetUserBetHistory(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserBetsResponse, error)
	GetPendingBetsForSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionBetsResponse, error)
	GetUserBalance(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type bettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBettingServiceClient(cc grpc.ClientConnInterface) BettingServiceClient {
	return &bettingServiceClient{cc}
}

func (c *bettingServiceClient) CreatePodiumBet(ctx context.Context, in *PodiumBetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, BettingService_CreatePodiumBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) CreatePolePositionBet(ctx context.Context, in *PolePositionBetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, BettingService_CreatePolePositionBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) CreateRainBet(ctx context.Context, in *RainBetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, BettingService_CreateRainBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) CreateRetirementBet(ctx context.Context, in *RetirementBetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, BettingService_CreateRetirementBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) CreateFastestLapBet(ctx context.Context, in *FastestLapBetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, BettingService_CreateFastestLapBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) CreateLapTimingBet(ctx context.Context, in *LapTimingBetRequest, opts ...grpc.CallOption) (*BetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BetResponse)
	err := c.cc.Invoke(ctx, BettingService_CreateLapTimingBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) GetUserActiveBets(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserBetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserBetsResponse)
	err := c.cc.Invoke(ctx, BettingService_GetUserActiveBets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) GetUserBetHistory(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserBetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserBetsResponse)
	err := c.cc.Invoke(ctx, BettingService_GetUserBetHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) GetPendingBetsForSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionBetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SessionBetsResponse)
	err := c.cc.Invoke(ctx, BettingService_GetPendingBetsForSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bettingServiceClient) GetUserBalance(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, BettingService_GetUserBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BettingServiceServer is the server API for BettingService service.
// All implementations must embed UnimplementedBettingServiceServer
// for forward compatibility.
type BettingServiceServer interface {
	CreatePodiumBet(context.Context, *PodiumBetRequest) (*BetResponse, error)
	CreatePolePositionBet(context.Context, *PolePositionBetRequest) (*BetResponse, error)
	CreateRainBet(context.Context, *RainBetRequest) (*BetResponse, error)
	CreateRetirementBet(context.Context, *RetirementBetRequest) (*BetResponse, error)
	CreateFastestLapBet(context.Context, *FastestLapBetRequest) (*BetResponse, error)
	CreateLapTimingBet(context.Context, *LapTimingBetRequest) (*BetResponse, error)
	GetUserActiveBets(context.Context, *UserRequest) (*UserBetsResponse, error)
	GetUserBetHistory(context.Context, *UserRequest) (*UserBetsResponse, error)
	GetPendingBetsForSession(context.Context, *SessionRequest) (*SessionBetsResponse, error)
	GetUserBalance(context.Context, *UserRequest) (*BalanceResponse, error)
	mustEmbedUnimplementedBettingServiceServer()
}

// UnimplementedBettingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBettingServiceServer struct{}

func (UnimplementedBettingServiceServer) CreatePodiumBet(context.Context, *PodiumBetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePodiumBet not implemented")
}
func (UnimplementedBettingServiceServer) CreatePolePositionBet(context.Context, *PolePositionBetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolePositionBet not implemented")
}
func (UnimplementedBettingServiceServer) CreateRainBet(context.Context, *RainBetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRainBet not implemented")
}
func (UnimplementedBettingServiceServer) CreateRetirementBet(context.Context, *RetirementBetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRetirementBet not implemented")
}
func (UnimplementedBettingServiceServer) CreateFastestLapBet(context.Context, *FastestLapBetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFastestLapBet not implemented")
}
func (UnimplementedBettingServiceServer) CreateLapTimingBet(context.Context, *LapTimingBetRequest) (*BetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLapTimingBet not implemented")
}
func (UnimplementedBettingServiceServer) GetUserActiveBets(context.Context, *UserRequest) (*UserBetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserActiveBets not implemented")
}
func (UnimplementedBettingServiceServer) GetUserBetHistory(context.Context, *UserRequest) (*UserBetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBetHistory not implemented")
}
func (UnimplementedBettingServiceServer) GetPendingBetsForSession(context.Context, *SessionRequest) (*SessionBetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingBetsForSession not implemented")
}
func (UnimplementedBettingServiceServer) GetUserBalance(context.Context, *UserRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedBettingServiceServer) mustEmbedUnimplementedBettingServiceServer() {}
func (UnimplementedBettingServiceServer) testEmbeddedByValue()                        {}

// UnsafeBettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BettingServiceServer will
// result in compilation errors.
type UnsafeBettingServiceServer interface {
	mustEmbedUnimplementedBettingServiceServer()
}

func RegisterBettingServiceServer(s grpc.ServiceRegistrar, srv BettingServiceServer) {
	// If the following call pancis, it indicates UnimplementedBettingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BettingService_ServiceDesc, srv)
}

func _BettingService_CreatePodiumBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodiumBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).CreatePodiumBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_CreatePodiumBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).CreatePodiumBet(ctx, req.(*PodiumBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_CreatePolePositionBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolePositionBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).CreatePolePositionBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_CreatePolePositionBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).CreatePolePositionBet(ctx, req.(*PolePositionBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_CreateRainBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RainBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).CreateRainBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_CreateRainBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).CreateRainBet(ctx, req.(*RainBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_CreateRetirementBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetirementBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).CreateRetirementBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_CreateRetirementBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).CreateRetirementBet(ctx, req.(*RetirementBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_CreateFastestLapBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastestLapBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).CreateFastestLapBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_CreateFastestLapBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).CreateFastestLapBet(ctx, req.(*FastestLapBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_CreateLapTimingBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LapTimingBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).CreateLapTimingBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_CreateLapTimingBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).CreateLapTimingBet(ctx, req.(*LapTimingBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_GetUserActiveBets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).GetUserActiveBets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_GetUserActiveBets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).GetUserActiveBets(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_GetUserBetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).GetUserBetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_GetUserBetHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).GetUserBetHistory(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_GetPendingBetsForSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).GetPendingBetsForSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_GetPendingBetsForSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).GetPendingBetsForSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BettingService_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BettingServiceServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BettingService_GetUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BettingServiceServer).GetUserBalance(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BettingService_ServiceDesc is the grpc.ServiceDesc for BettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "betting.BettingService",
	HandlerType: (*BettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePodiumBet",
			Handler:    _BettingService_CreatePodiumBet_Handler,
		},
		{
			MethodName: "CreatePolePositionBet",
			Handler:    _BettingService_CreatePolePositionBet_Handler,
		},
		{
			MethodName: "CreateRainBet",
			Handler:    _BettingService_CreateRainBet_Handler,
		},
		{
			MethodName: "CreateRetirementBet",
			Handler:    _BettingService_CreateRetirementBet_Handler,
		},
		{
			MethodName: "CreateFastestLapBet",
			Handler:    _BettingService_CreateFastestLapBet_Handler,
		},
		{
			MethodName: "CreateLapTimingBet",
			Handler:    _BettingService_CreateLapTimingBet_Handler,
		},
		{
			MethodName: "GetUserActiveBets",
			Handler:    _BettingService_GetUserActiveBets_Handler,
		},
		{
			MethodName: "GetUserBetHistory",
			Handler:    _BettingService_GetUserBetHistory_Handler,
		},
		{
			MethodName: "GetPendingBetsForSession",
			Handler:    _BettingService_GetPendingBetsForSession_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _BettingService_GetUserBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/betting_system.proto",
}
