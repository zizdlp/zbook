// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: service_zbook_user.proto

package pb

import (
	context "context"
	rpcs "github.com/zizdlp/zbook/pb/rpcs"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ZBookUser_CreateUser_FullMethodName           = "/pb.ZBookUser/CreateUser"
	ZBookUser_LoginUser_FullMethodName            = "/pb.ZBookUser/LoginUser"
	ZBookUser_UpdateUser_FullMethodName           = "/pb.ZBookUser/UpdateUser"
	ZBookUser_UpdateUserOnBoarding_FullMethodName = "/pb.ZBookUser/UpdateUserOnBoarding"
	ZBookUser_QueryUser_FullMethodName            = "/pb.ZBookUser/QueryUser"
	ZBookUser_GetUserInfo_FullMethodName          = "/pb.ZBookUser/GetUserInfo"
	ZBookUser_GetUserAvatar_FullMethodName        = "/pb.ZBookUser/GetUserAvatar"
	ZBookUser_ListUser_FullMethodName             = "/pb.ZBookUser/ListUser"
	ZBookUser_GetListUserCount_FullMethodName     = "/pb.ZBookUser/GetListUserCount"
	ZBookUser_GetQueryUserCount_FullMethodName    = "/pb.ZBookUser/GetQueryUserCount"
)

// ZBookUserClient is the client API for ZBookUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZBookUserClient interface {
	// 1.CreateUser
	CreateUser(ctx context.Context, in *rpcs.CreateUserRequest, opts ...grpc.CallOption) (*rpcs.CreateUserResponse, error)
	// 2.LoginUser
	LoginUser(ctx context.Context, in *rpcs.LoginUserRequest, opts ...grpc.CallOption) (*rpcs.LoginUserResponse, error)
	// 3.UpdateUser
	UpdateUser(ctx context.Context, in *rpcs.UpdateUserRequest, opts ...grpc.CallOption) (*rpcs.UpdateUserResponse, error)
	// 4.UpdateUserOnBoarding
	UpdateUserOnBoarding(ctx context.Context, in *rpcs.UpdateUserOnBoardingRequest, opts ...grpc.CallOption) (*rpcs.UpdateUserOnBoardingResponse, error)
	// 5.QueryUser
	QueryUser(ctx context.Context, in *rpcs.QueryUserRequest, opts ...grpc.CallOption) (*rpcs.QueryUserResponse, error)
	// 6.GetUserInfo
	GetUserInfo(ctx context.Context, in *rpcs.GetUserInfoRequest, opts ...grpc.CallOption) (*rpcs.GetUserInfoResponse, error)
	// 7.GetUserAvatar
	GetUserAvatar(ctx context.Context, in *rpcs.GetUserAvatarRequest, opts ...grpc.CallOption) (*rpcs.GetUserAvatarResponse, error)
	// 8.ListUser
	ListUser(ctx context.Context, in *rpcs.ListUserRequest, opts ...grpc.CallOption) (*rpcs.ListUserResponse, error)
	// 9.GetListUserCount
	GetListUserCount(ctx context.Context, in *rpcs.GetListUserCountRequest, opts ...grpc.CallOption) (*rpcs.GetListUserCountResponse, error)
	// 10.GetQueryUserCount
	GetQueryUserCount(ctx context.Context, in *rpcs.GetQueryUserCountRequest, opts ...grpc.CallOption) (*rpcs.GetQueryUserCountResponse, error)
}

type zBookUserClient struct {
	cc grpc.ClientConnInterface
}

func NewZBookUserClient(cc grpc.ClientConnInterface) ZBookUserClient {
	return &zBookUserClient{cc}
}

func (c *zBookUserClient) CreateUser(ctx context.Context, in *rpcs.CreateUserRequest, opts ...grpc.CallOption) (*rpcs.CreateUserResponse, error) {
	out := new(rpcs.CreateUserResponse)
	err := c.cc.Invoke(ctx, ZBookUser_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) LoginUser(ctx context.Context, in *rpcs.LoginUserRequest, opts ...grpc.CallOption) (*rpcs.LoginUserResponse, error) {
	out := new(rpcs.LoginUserResponse)
	err := c.cc.Invoke(ctx, ZBookUser_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) UpdateUser(ctx context.Context, in *rpcs.UpdateUserRequest, opts ...grpc.CallOption) (*rpcs.UpdateUserResponse, error) {
	out := new(rpcs.UpdateUserResponse)
	err := c.cc.Invoke(ctx, ZBookUser_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) UpdateUserOnBoarding(ctx context.Context, in *rpcs.UpdateUserOnBoardingRequest, opts ...grpc.CallOption) (*rpcs.UpdateUserOnBoardingResponse, error) {
	out := new(rpcs.UpdateUserOnBoardingResponse)
	err := c.cc.Invoke(ctx, ZBookUser_UpdateUserOnBoarding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) QueryUser(ctx context.Context, in *rpcs.QueryUserRequest, opts ...grpc.CallOption) (*rpcs.QueryUserResponse, error) {
	out := new(rpcs.QueryUserResponse)
	err := c.cc.Invoke(ctx, ZBookUser_QueryUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) GetUserInfo(ctx context.Context, in *rpcs.GetUserInfoRequest, opts ...grpc.CallOption) (*rpcs.GetUserInfoResponse, error) {
	out := new(rpcs.GetUserInfoResponse)
	err := c.cc.Invoke(ctx, ZBookUser_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) GetUserAvatar(ctx context.Context, in *rpcs.GetUserAvatarRequest, opts ...grpc.CallOption) (*rpcs.GetUserAvatarResponse, error) {
	out := new(rpcs.GetUserAvatarResponse)
	err := c.cc.Invoke(ctx, ZBookUser_GetUserAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) ListUser(ctx context.Context, in *rpcs.ListUserRequest, opts ...grpc.CallOption) (*rpcs.ListUserResponse, error) {
	out := new(rpcs.ListUserResponse)
	err := c.cc.Invoke(ctx, ZBookUser_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) GetListUserCount(ctx context.Context, in *rpcs.GetListUserCountRequest, opts ...grpc.CallOption) (*rpcs.GetListUserCountResponse, error) {
	out := new(rpcs.GetListUserCountResponse)
	err := c.cc.Invoke(ctx, ZBookUser_GetListUserCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookUserClient) GetQueryUserCount(ctx context.Context, in *rpcs.GetQueryUserCountRequest, opts ...grpc.CallOption) (*rpcs.GetQueryUserCountResponse, error) {
	out := new(rpcs.GetQueryUserCountResponse)
	err := c.cc.Invoke(ctx, ZBookUser_GetQueryUserCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZBookUserServer is the server API for ZBookUser service.
// All implementations must embed UnimplementedZBookUserServer
// for forward compatibility
type ZBookUserServer interface {
	// 1.CreateUser
	CreateUser(context.Context, *rpcs.CreateUserRequest) (*rpcs.CreateUserResponse, error)
	// 2.LoginUser
	LoginUser(context.Context, *rpcs.LoginUserRequest) (*rpcs.LoginUserResponse, error)
	// 3.UpdateUser
	UpdateUser(context.Context, *rpcs.UpdateUserRequest) (*rpcs.UpdateUserResponse, error)
	// 4.UpdateUserOnBoarding
	UpdateUserOnBoarding(context.Context, *rpcs.UpdateUserOnBoardingRequest) (*rpcs.UpdateUserOnBoardingResponse, error)
	// 5.QueryUser
	QueryUser(context.Context, *rpcs.QueryUserRequest) (*rpcs.QueryUserResponse, error)
	// 6.GetUserInfo
	GetUserInfo(context.Context, *rpcs.GetUserInfoRequest) (*rpcs.GetUserInfoResponse, error)
	// 7.GetUserAvatar
	GetUserAvatar(context.Context, *rpcs.GetUserAvatarRequest) (*rpcs.GetUserAvatarResponse, error)
	// 8.ListUser
	ListUser(context.Context, *rpcs.ListUserRequest) (*rpcs.ListUserResponse, error)
	// 9.GetListUserCount
	GetListUserCount(context.Context, *rpcs.GetListUserCountRequest) (*rpcs.GetListUserCountResponse, error)
	// 10.GetQueryUserCount
	GetQueryUserCount(context.Context, *rpcs.GetQueryUserCountRequest) (*rpcs.GetQueryUserCountResponse, error)
	mustEmbedUnimplementedZBookUserServer()
}

// UnimplementedZBookUserServer must be embedded to have forward compatible implementations.
type UnimplementedZBookUserServer struct {
}

func (UnimplementedZBookUserServer) CreateUser(context.Context, *rpcs.CreateUserRequest) (*rpcs.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedZBookUserServer) LoginUser(context.Context, *rpcs.LoginUserRequest) (*rpcs.LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedZBookUserServer) UpdateUser(context.Context, *rpcs.UpdateUserRequest) (*rpcs.UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedZBookUserServer) UpdateUserOnBoarding(context.Context, *rpcs.UpdateUserOnBoardingRequest) (*rpcs.UpdateUserOnBoardingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserOnBoarding not implemented")
}
func (UnimplementedZBookUserServer) QueryUser(context.Context, *rpcs.QueryUserRequest) (*rpcs.QueryUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}
func (UnimplementedZBookUserServer) GetUserInfo(context.Context, *rpcs.GetUserInfoRequest) (*rpcs.GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedZBookUserServer) GetUserAvatar(context.Context, *rpcs.GetUserAvatarRequest) (*rpcs.GetUserAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAvatar not implemented")
}
func (UnimplementedZBookUserServer) ListUser(context.Context, *rpcs.ListUserRequest) (*rpcs.ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedZBookUserServer) GetListUserCount(context.Context, *rpcs.GetListUserCountRequest) (*rpcs.GetListUserCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUserCount not implemented")
}
func (UnimplementedZBookUserServer) GetQueryUserCount(context.Context, *rpcs.GetQueryUserCountRequest) (*rpcs.GetQueryUserCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryUserCount not implemented")
}
func (UnimplementedZBookUserServer) mustEmbedUnimplementedZBookUserServer() {}

// UnsafeZBookUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZBookUserServer will
// result in compilation errors.
type UnsafeZBookUserServer interface {
	mustEmbedUnimplementedZBookUserServer()
}

func RegisterZBookUserServer(s grpc.ServiceRegistrar, srv ZBookUserServer) {
	s.RegisterService(&ZBookUser_ServiceDesc, srv)
}

func _ZBookUser_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).CreateUser(ctx, req.(*rpcs.CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).LoginUser(ctx, req.(*rpcs.LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).UpdateUser(ctx, req.(*rpcs.UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_UpdateUserOnBoarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.UpdateUserOnBoardingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).UpdateUserOnBoarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_UpdateUserOnBoarding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).UpdateUserOnBoarding(ctx, req.(*rpcs.UpdateUserOnBoardingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.QueryUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_QueryUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).QueryUser(ctx, req.(*rpcs.QueryUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).GetUserInfo(ctx, req.(*rpcs.GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_GetUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetUserAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).GetUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_GetUserAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).GetUserAvatar(ctx, req.(*rpcs.GetUserAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).ListUser(ctx, req.(*rpcs.ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_GetListUserCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetListUserCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).GetListUserCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_GetListUserCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).GetListUserCount(ctx, req.(*rpcs.GetListUserCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookUser_GetQueryUserCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetQueryUserCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookUserServer).GetQueryUserCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookUser_GetQueryUserCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookUserServer).GetQueryUserCount(ctx, req.(*rpcs.GetQueryUserCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZBookUser_ServiceDesc is the grpc.ServiceDesc for ZBookUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZBookUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ZBookUser",
	HandlerType: (*ZBookUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ZBookUser_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _ZBookUser_LoginUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ZBookUser_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserOnBoarding",
			Handler:    _ZBookUser_UpdateUserOnBoarding_Handler,
		},
		{
			MethodName: "QueryUser",
			Handler:    _ZBookUser_QueryUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _ZBookUser_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserAvatar",
			Handler:    _ZBookUser_GetUserAvatar_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _ZBookUser_ListUser_Handler,
		},
		{
			MethodName: "GetListUserCount",
			Handler:    _ZBookUser_GetListUserCount_Handler,
		},
		{
			MethodName: "GetQueryUserCount",
			Handler:    _ZBookUser_GetQueryUserCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_zbook_user.proto",
}
