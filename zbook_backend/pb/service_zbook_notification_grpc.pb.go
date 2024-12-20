// clang-format off

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: service_zbook_notification.proto

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
	ZBookNotification_ListFollowerNotification_FullMethodName                 = "/pb.ZBookNotification/ListFollowerNotification"
	ZBookNotification_ListRepoNotification_FullMethodName                     = "/pb.ZBookNotification/ListRepoNotification"
	ZBookNotification_ListCommentNotification_FullMethodName                  = "/pb.ZBookNotification/ListCommentNotification"
	ZBookNotification_ListSystemNotification_FullMethodName                   = "/pb.ZBookNotification/ListSystemNotification"
	ZBookNotification_MarkFollowerNotificationReaded_FullMethodName           = "/pb.ZBookNotification/MarkFollowerNotificationReaded"
	ZBookNotification_MarkSystemNotificationReaded_FullMethodName             = "/pb.ZBookNotification/MarkSystemNotificationReaded"
	ZBookNotification_MarkCommentNotificationReaded_FullMethodName            = "/pb.ZBookNotification/MarkCommentNotificationReaded"
	ZBookNotification_MarkRepoNotificationReaded_FullMethodName               = "/pb.ZBookNotification/MarkRepoNotificationReaded"
	ZBookNotification_GetUnReadCount_FullMethodName                           = "/pb.ZBookNotification/GetUnReadCount"
	ZBookNotification_ResetUnreadCount_FullMethodName                         = "/pb.ZBookNotification/ResetUnreadCount"
	ZBookNotification_GetListFollowerNotificationUnreadedCount_FullMethodName = "/pb.ZBookNotification/GetListFollowerNotificationUnreadedCount"
	ZBookNotification_GetListRepoNotificationUnreadedCount_FullMethodName     = "/pb.ZBookNotification/GetListRepoNotificationUnreadedCount"
	ZBookNotification_GetListCommentNotificationUnreadedCount_FullMethodName  = "/pb.ZBookNotification/GetListCommentNotificationUnreadedCount"
	ZBookNotification_GetListSystemNotificationUnreadedCount_FullMethodName   = "/pb.ZBookNotification/GetListSystemNotificationUnreadedCount"
)

// ZBookNotificationClient is the client API for ZBookNotification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZBookNotificationClient interface {
	// 1.ListFollowerNotification
	ListFollowerNotification(ctx context.Context, in *rpcs.ListFollowerNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListFollowerNotificationResponse, error)
	// 2.ListRepoNotification
	ListRepoNotification(ctx context.Context, in *rpcs.ListRepoNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListRepoNotificationResponse, error)
	// 3.ListCommentNotification
	ListCommentNotification(ctx context.Context, in *rpcs.ListCommentNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListCommentNotificationResponse, error)
	// 4.ListSystemNotification
	ListSystemNotification(ctx context.Context, in *rpcs.ListSystemNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListSystemNotificationResponse, error)
	// 5.MarkFollowerNotificationReaded
	MarkFollowerNotificationReaded(ctx context.Context, in *rpcs.MarkFollowerNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error)
	// 6.MarkSystemNotificationReaded
	MarkSystemNotificationReaded(ctx context.Context, in *rpcs.MarkSystemNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error)
	// 7.MarkCommentNotificationReaded
	MarkCommentNotificationReaded(ctx context.Context, in *rpcs.MarkCommentNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error)
	// 8.MarkRepoNotificationReaded
	MarkRepoNotificationReaded(ctx context.Context, in *rpcs.MarkRepoNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error)
	// 9.GetUnReadCount
	GetUnReadCount(ctx context.Context, in *rpcs.GetUnReadCountRequest, opts ...grpc.CallOption) (*rpcs.GetUnReadCountResponse, error)
	// 10.ResetUnreadCount
	ResetUnreadCount(ctx context.Context, in *rpcs.ResetUnreadCountRequest, opts ...grpc.CallOption) (*rpcs.ResetUnreadCountResponse, error)
	// 11.GetListFollowerNotificationUnreadedCount
	GetListFollowerNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListFollowerNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListFollowerNotificationUnreadedCountResponse, error)
	// 12.GetListRepoNotificationUnreadedCount
	GetListRepoNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListRepoNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListRepoNotificationUnreadedCountResponse, error)
	// 13.GetListCommentNotificationUnreadedCount
	GetListCommentNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListCommentNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListCommentNotificationUnreadedCountResponse, error)
	// 14.GetListSystemNotificationUnreadedCount
	GetListSystemNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListSystemNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListSystemNotificationUnreadedCountResponse, error)
}

type zBookNotificationClient struct {
	cc grpc.ClientConnInterface
}

func NewZBookNotificationClient(cc grpc.ClientConnInterface) ZBookNotificationClient {
	return &zBookNotificationClient{cc}
}

func (c *zBookNotificationClient) ListFollowerNotification(ctx context.Context, in *rpcs.ListFollowerNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListFollowerNotificationResponse, error) {
	out := new(rpcs.ListFollowerNotificationResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_ListFollowerNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) ListRepoNotification(ctx context.Context, in *rpcs.ListRepoNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListRepoNotificationResponse, error) {
	out := new(rpcs.ListRepoNotificationResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_ListRepoNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) ListCommentNotification(ctx context.Context, in *rpcs.ListCommentNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListCommentNotificationResponse, error) {
	out := new(rpcs.ListCommentNotificationResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_ListCommentNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) ListSystemNotification(ctx context.Context, in *rpcs.ListSystemNotificationRequest, opts ...grpc.CallOption) (*rpcs.ListSystemNotificationResponse, error) {
	out := new(rpcs.ListSystemNotificationResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_ListSystemNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) MarkFollowerNotificationReaded(ctx context.Context, in *rpcs.MarkFollowerNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error) {
	out := new(rpcs.SetNotiReadResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_MarkFollowerNotificationReaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) MarkSystemNotificationReaded(ctx context.Context, in *rpcs.MarkSystemNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error) {
	out := new(rpcs.SetNotiReadResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_MarkSystemNotificationReaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) MarkCommentNotificationReaded(ctx context.Context, in *rpcs.MarkCommentNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error) {
	out := new(rpcs.SetNotiReadResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_MarkCommentNotificationReaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) MarkRepoNotificationReaded(ctx context.Context, in *rpcs.MarkRepoNotificationReadedRequest, opts ...grpc.CallOption) (*rpcs.SetNotiReadResponse, error) {
	out := new(rpcs.SetNotiReadResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_MarkRepoNotificationReaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) GetUnReadCount(ctx context.Context, in *rpcs.GetUnReadCountRequest, opts ...grpc.CallOption) (*rpcs.GetUnReadCountResponse, error) {
	out := new(rpcs.GetUnReadCountResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_GetUnReadCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) ResetUnreadCount(ctx context.Context, in *rpcs.ResetUnreadCountRequest, opts ...grpc.CallOption) (*rpcs.ResetUnreadCountResponse, error) {
	out := new(rpcs.ResetUnreadCountResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_ResetUnreadCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) GetListFollowerNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListFollowerNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListFollowerNotificationUnreadedCountResponse, error) {
	out := new(rpcs.GetListFollowerNotificationUnreadedCountResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_GetListFollowerNotificationUnreadedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) GetListRepoNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListRepoNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListRepoNotificationUnreadedCountResponse, error) {
	out := new(rpcs.GetListRepoNotificationUnreadedCountResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_GetListRepoNotificationUnreadedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) GetListCommentNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListCommentNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListCommentNotificationUnreadedCountResponse, error) {
	out := new(rpcs.GetListCommentNotificationUnreadedCountResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_GetListCommentNotificationUnreadedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookNotificationClient) GetListSystemNotificationUnreadedCount(ctx context.Context, in *rpcs.GetListSystemNotificationUnreadedCountRequest, opts ...grpc.CallOption) (*rpcs.GetListSystemNotificationUnreadedCountResponse, error) {
	out := new(rpcs.GetListSystemNotificationUnreadedCountResponse)
	err := c.cc.Invoke(ctx, ZBookNotification_GetListSystemNotificationUnreadedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZBookNotificationServer is the server API for ZBookNotification service.
// All implementations must embed UnimplementedZBookNotificationServer
// for forward compatibility
type ZBookNotificationServer interface {
	// 1.ListFollowerNotification
	ListFollowerNotification(context.Context, *rpcs.ListFollowerNotificationRequest) (*rpcs.ListFollowerNotificationResponse, error)
	// 2.ListRepoNotification
	ListRepoNotification(context.Context, *rpcs.ListRepoNotificationRequest) (*rpcs.ListRepoNotificationResponse, error)
	// 3.ListCommentNotification
	ListCommentNotification(context.Context, *rpcs.ListCommentNotificationRequest) (*rpcs.ListCommentNotificationResponse, error)
	// 4.ListSystemNotification
	ListSystemNotification(context.Context, *rpcs.ListSystemNotificationRequest) (*rpcs.ListSystemNotificationResponse, error)
	// 5.MarkFollowerNotificationReaded
	MarkFollowerNotificationReaded(context.Context, *rpcs.MarkFollowerNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error)
	// 6.MarkSystemNotificationReaded
	MarkSystemNotificationReaded(context.Context, *rpcs.MarkSystemNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error)
	// 7.MarkCommentNotificationReaded
	MarkCommentNotificationReaded(context.Context, *rpcs.MarkCommentNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error)
	// 8.MarkRepoNotificationReaded
	MarkRepoNotificationReaded(context.Context, *rpcs.MarkRepoNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error)
	// 9.GetUnReadCount
	GetUnReadCount(context.Context, *rpcs.GetUnReadCountRequest) (*rpcs.GetUnReadCountResponse, error)
	// 10.ResetUnreadCount
	ResetUnreadCount(context.Context, *rpcs.ResetUnreadCountRequest) (*rpcs.ResetUnreadCountResponse, error)
	// 11.GetListFollowerNotificationUnreadedCount
	GetListFollowerNotificationUnreadedCount(context.Context, *rpcs.GetListFollowerNotificationUnreadedCountRequest) (*rpcs.GetListFollowerNotificationUnreadedCountResponse, error)
	// 12.GetListRepoNotificationUnreadedCount
	GetListRepoNotificationUnreadedCount(context.Context, *rpcs.GetListRepoNotificationUnreadedCountRequest) (*rpcs.GetListRepoNotificationUnreadedCountResponse, error)
	// 13.GetListCommentNotificationUnreadedCount
	GetListCommentNotificationUnreadedCount(context.Context, *rpcs.GetListCommentNotificationUnreadedCountRequest) (*rpcs.GetListCommentNotificationUnreadedCountResponse, error)
	// 14.GetListSystemNotificationUnreadedCount
	GetListSystemNotificationUnreadedCount(context.Context, *rpcs.GetListSystemNotificationUnreadedCountRequest) (*rpcs.GetListSystemNotificationUnreadedCountResponse, error)
	mustEmbedUnimplementedZBookNotificationServer()
}

// UnimplementedZBookNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedZBookNotificationServer struct {
}

func (UnimplementedZBookNotificationServer) ListFollowerNotification(context.Context, *rpcs.ListFollowerNotificationRequest) (*rpcs.ListFollowerNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollowerNotification not implemented")
}
func (UnimplementedZBookNotificationServer) ListRepoNotification(context.Context, *rpcs.ListRepoNotificationRequest) (*rpcs.ListRepoNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepoNotification not implemented")
}
func (UnimplementedZBookNotificationServer) ListCommentNotification(context.Context, *rpcs.ListCommentNotificationRequest) (*rpcs.ListCommentNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommentNotification not implemented")
}
func (UnimplementedZBookNotificationServer) ListSystemNotification(context.Context, *rpcs.ListSystemNotificationRequest) (*rpcs.ListSystemNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystemNotification not implemented")
}
func (UnimplementedZBookNotificationServer) MarkFollowerNotificationReaded(context.Context, *rpcs.MarkFollowerNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkFollowerNotificationReaded not implemented")
}
func (UnimplementedZBookNotificationServer) MarkSystemNotificationReaded(context.Context, *rpcs.MarkSystemNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkSystemNotificationReaded not implemented")
}
func (UnimplementedZBookNotificationServer) MarkCommentNotificationReaded(context.Context, *rpcs.MarkCommentNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkCommentNotificationReaded not implemented")
}
func (UnimplementedZBookNotificationServer) MarkRepoNotificationReaded(context.Context, *rpcs.MarkRepoNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkRepoNotificationReaded not implemented")
}
func (UnimplementedZBookNotificationServer) GetUnReadCount(context.Context, *rpcs.GetUnReadCountRequest) (*rpcs.GetUnReadCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnReadCount not implemented")
}
func (UnimplementedZBookNotificationServer) ResetUnreadCount(context.Context, *rpcs.ResetUnreadCountRequest) (*rpcs.ResetUnreadCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetUnreadCount not implemented")
}
func (UnimplementedZBookNotificationServer) GetListFollowerNotificationUnreadedCount(context.Context, *rpcs.GetListFollowerNotificationUnreadedCountRequest) (*rpcs.GetListFollowerNotificationUnreadedCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFollowerNotificationUnreadedCount not implemented")
}
func (UnimplementedZBookNotificationServer) GetListRepoNotificationUnreadedCount(context.Context, *rpcs.GetListRepoNotificationUnreadedCountRequest) (*rpcs.GetListRepoNotificationUnreadedCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListRepoNotificationUnreadedCount not implemented")
}
func (UnimplementedZBookNotificationServer) GetListCommentNotificationUnreadedCount(context.Context, *rpcs.GetListCommentNotificationUnreadedCountRequest) (*rpcs.GetListCommentNotificationUnreadedCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListCommentNotificationUnreadedCount not implemented")
}
func (UnimplementedZBookNotificationServer) GetListSystemNotificationUnreadedCount(context.Context, *rpcs.GetListSystemNotificationUnreadedCountRequest) (*rpcs.GetListSystemNotificationUnreadedCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListSystemNotificationUnreadedCount not implemented")
}
func (UnimplementedZBookNotificationServer) mustEmbedUnimplementedZBookNotificationServer() {}

// UnsafeZBookNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZBookNotificationServer will
// result in compilation errors.
type UnsafeZBookNotificationServer interface {
	mustEmbedUnimplementedZBookNotificationServer()
}

func RegisterZBookNotificationServer(s grpc.ServiceRegistrar, srv ZBookNotificationServer) {
	s.RegisterService(&ZBookNotification_ServiceDesc, srv)
}

func _ZBookNotification_ListFollowerNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ListFollowerNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).ListFollowerNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_ListFollowerNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).ListFollowerNotification(ctx, req.(*rpcs.ListFollowerNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_ListRepoNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ListRepoNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).ListRepoNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_ListRepoNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).ListRepoNotification(ctx, req.(*rpcs.ListRepoNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_ListCommentNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ListCommentNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).ListCommentNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_ListCommentNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).ListCommentNotification(ctx, req.(*rpcs.ListCommentNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_ListSystemNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ListSystemNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).ListSystemNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_ListSystemNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).ListSystemNotification(ctx, req.(*rpcs.ListSystemNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_MarkFollowerNotificationReaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.MarkFollowerNotificationReadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).MarkFollowerNotificationReaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_MarkFollowerNotificationReaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).MarkFollowerNotificationReaded(ctx, req.(*rpcs.MarkFollowerNotificationReadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_MarkSystemNotificationReaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.MarkSystemNotificationReadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).MarkSystemNotificationReaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_MarkSystemNotificationReaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).MarkSystemNotificationReaded(ctx, req.(*rpcs.MarkSystemNotificationReadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_MarkCommentNotificationReaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.MarkCommentNotificationReadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).MarkCommentNotificationReaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_MarkCommentNotificationReaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).MarkCommentNotificationReaded(ctx, req.(*rpcs.MarkCommentNotificationReadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_MarkRepoNotificationReaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.MarkRepoNotificationReadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).MarkRepoNotificationReaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_MarkRepoNotificationReaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).MarkRepoNotificationReaded(ctx, req.(*rpcs.MarkRepoNotificationReadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_GetUnReadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetUnReadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).GetUnReadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_GetUnReadCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).GetUnReadCount(ctx, req.(*rpcs.GetUnReadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_ResetUnreadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ResetUnreadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).ResetUnreadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_ResetUnreadCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).ResetUnreadCount(ctx, req.(*rpcs.ResetUnreadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_GetListFollowerNotificationUnreadedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetListFollowerNotificationUnreadedCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).GetListFollowerNotificationUnreadedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_GetListFollowerNotificationUnreadedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).GetListFollowerNotificationUnreadedCount(ctx, req.(*rpcs.GetListFollowerNotificationUnreadedCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_GetListRepoNotificationUnreadedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetListRepoNotificationUnreadedCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).GetListRepoNotificationUnreadedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_GetListRepoNotificationUnreadedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).GetListRepoNotificationUnreadedCount(ctx, req.(*rpcs.GetListRepoNotificationUnreadedCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_GetListCommentNotificationUnreadedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetListCommentNotificationUnreadedCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).GetListCommentNotificationUnreadedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_GetListCommentNotificationUnreadedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).GetListCommentNotificationUnreadedCount(ctx, req.(*rpcs.GetListCommentNotificationUnreadedCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookNotification_GetListSystemNotificationUnreadedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetListSystemNotificationUnreadedCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookNotificationServer).GetListSystemNotificationUnreadedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookNotification_GetListSystemNotificationUnreadedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookNotificationServer).GetListSystemNotificationUnreadedCount(ctx, req.(*rpcs.GetListSystemNotificationUnreadedCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZBookNotification_ServiceDesc is the grpc.ServiceDesc for ZBookNotification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZBookNotification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ZBookNotification",
	HandlerType: (*ZBookNotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFollowerNotification",
			Handler:    _ZBookNotification_ListFollowerNotification_Handler,
		},
		{
			MethodName: "ListRepoNotification",
			Handler:    _ZBookNotification_ListRepoNotification_Handler,
		},
		{
			MethodName: "ListCommentNotification",
			Handler:    _ZBookNotification_ListCommentNotification_Handler,
		},
		{
			MethodName: "ListSystemNotification",
			Handler:    _ZBookNotification_ListSystemNotification_Handler,
		},
		{
			MethodName: "MarkFollowerNotificationReaded",
			Handler:    _ZBookNotification_MarkFollowerNotificationReaded_Handler,
		},
		{
			MethodName: "MarkSystemNotificationReaded",
			Handler:    _ZBookNotification_MarkSystemNotificationReaded_Handler,
		},
		{
			MethodName: "MarkCommentNotificationReaded",
			Handler:    _ZBookNotification_MarkCommentNotificationReaded_Handler,
		},
		{
			MethodName: "MarkRepoNotificationReaded",
			Handler:    _ZBookNotification_MarkRepoNotificationReaded_Handler,
		},
		{
			MethodName: "GetUnReadCount",
			Handler:    _ZBookNotification_GetUnReadCount_Handler,
		},
		{
			MethodName: "ResetUnreadCount",
			Handler:    _ZBookNotification_ResetUnreadCount_Handler,
		},
		{
			MethodName: "GetListFollowerNotificationUnreadedCount",
			Handler:    _ZBookNotification_GetListFollowerNotificationUnreadedCount_Handler,
		},
		{
			MethodName: "GetListRepoNotificationUnreadedCount",
			Handler:    _ZBookNotification_GetListRepoNotificationUnreadedCount_Handler,
		},
		{
			MethodName: "GetListCommentNotificationUnreadedCount",
			Handler:    _ZBookNotification_GetListCommentNotificationUnreadedCount_Handler,
		},
		{
			MethodName: "GetListSystemNotificationUnreadedCount",
			Handler:    _ZBookNotification_GetListSystemNotificationUnreadedCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_zbook_notification.proto",
}
