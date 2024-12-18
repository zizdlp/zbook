// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: service_zbook_repo_relation.proto

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
	ZBookRepoRelation_CreateRepoRelation_FullMethodName         = "/pb.ZBookRepoRelation/CreateRepoRelation"
	ZBookRepoRelation_DeleteRepoRelation_FullMethodName         = "/pb.ZBookRepoRelation/DeleteRepoRelation"
	ZBookRepoRelation_CreateRepoVisibility_FullMethodName       = "/pb.ZBookRepoRelation/CreateRepoVisibility"
	ZBookRepoRelation_DeleteRepoVisibility_FullMethodName       = "/pb.ZBookRepoRelation/DeleteRepoVisibility"
	ZBookRepoRelation_ListSelectedUserByRepo_FullMethodName     = "/pb.ZBookRepoRelation/ListSelectedUserByRepo"
	ZBookRepoRelation_GetSelectedUserByRepoCount_FullMethodName = "/pb.ZBookRepoRelation/GetSelectedUserByRepoCount"
	ZBookRepoRelation_QueryUserByRepo_FullMethodName            = "/pb.ZBookRepoRelation/QueryUserByRepo"
)

// ZBookRepoRelationClient is the client API for ZBookRepoRelation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZBookRepoRelationClient interface {
	// 1.CreateRepoRelation
	CreateRepoRelation(ctx context.Context, in *rpcs.CreateRepoRelationRequest, opts ...grpc.CallOption) (*rpcs.CreateRepoRelationResponse, error)
	// 2.DeleteRepoRelation
	DeleteRepoRelation(ctx context.Context, in *rpcs.DeleteRepoRelationRequest, opts ...grpc.CallOption) (*rpcs.DeleteRepoRelationResponse, error)
	// 3.CreateRepoVisibility
	CreateRepoVisibility(ctx context.Context, in *rpcs.CreateRepoVisibilityRequest, opts ...grpc.CallOption) (*rpcs.CreateRepoVisibilityResponse, error)
	// 4.DeleteRepoVisibility
	DeleteRepoVisibility(ctx context.Context, in *rpcs.DeleteRepoVisibilityRequest, opts ...grpc.CallOption) (*rpcs.DeleteRepoVisibilityResponse, error)
	// 5.ListSelectedUserByRepo
	ListSelectedUserByRepo(ctx context.Context, in *rpcs.ListSelectedUserByRepoRequest, opts ...grpc.CallOption) (*rpcs.ListSelectedUserByRepoResponse, error)
	// 6.GetSelectedUserByRepoCount
	GetSelectedUserByRepoCount(ctx context.Context, in *rpcs.GetSelectedUserByRepoCountRequest, opts ...grpc.CallOption) (*rpcs.GetSelectedUserByRepoCountResponse, error)
	// 7.QueryUserByRepo
	QueryUserByRepo(ctx context.Context, in *rpcs.QueryUserByRepoRequest, opts ...grpc.CallOption) (*rpcs.QueryUserByRepoResponse, error)
}

type zBookRepoRelationClient struct {
	cc grpc.ClientConnInterface
}

func NewZBookRepoRelationClient(cc grpc.ClientConnInterface) ZBookRepoRelationClient {
	return &zBookRepoRelationClient{cc}
}

func (c *zBookRepoRelationClient) CreateRepoRelation(ctx context.Context, in *rpcs.CreateRepoRelationRequest, opts ...grpc.CallOption) (*rpcs.CreateRepoRelationResponse, error) {
	out := new(rpcs.CreateRepoRelationResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_CreateRepoRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookRepoRelationClient) DeleteRepoRelation(ctx context.Context, in *rpcs.DeleteRepoRelationRequest, opts ...grpc.CallOption) (*rpcs.DeleteRepoRelationResponse, error) {
	out := new(rpcs.DeleteRepoRelationResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_DeleteRepoRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookRepoRelationClient) CreateRepoVisibility(ctx context.Context, in *rpcs.CreateRepoVisibilityRequest, opts ...grpc.CallOption) (*rpcs.CreateRepoVisibilityResponse, error) {
	out := new(rpcs.CreateRepoVisibilityResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_CreateRepoVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookRepoRelationClient) DeleteRepoVisibility(ctx context.Context, in *rpcs.DeleteRepoVisibilityRequest, opts ...grpc.CallOption) (*rpcs.DeleteRepoVisibilityResponse, error) {
	out := new(rpcs.DeleteRepoVisibilityResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_DeleteRepoVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookRepoRelationClient) ListSelectedUserByRepo(ctx context.Context, in *rpcs.ListSelectedUserByRepoRequest, opts ...grpc.CallOption) (*rpcs.ListSelectedUserByRepoResponse, error) {
	out := new(rpcs.ListSelectedUserByRepoResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_ListSelectedUserByRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookRepoRelationClient) GetSelectedUserByRepoCount(ctx context.Context, in *rpcs.GetSelectedUserByRepoCountRequest, opts ...grpc.CallOption) (*rpcs.GetSelectedUserByRepoCountResponse, error) {
	out := new(rpcs.GetSelectedUserByRepoCountResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_GetSelectedUserByRepoCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zBookRepoRelationClient) QueryUserByRepo(ctx context.Context, in *rpcs.QueryUserByRepoRequest, opts ...grpc.CallOption) (*rpcs.QueryUserByRepoResponse, error) {
	out := new(rpcs.QueryUserByRepoResponse)
	err := c.cc.Invoke(ctx, ZBookRepoRelation_QueryUserByRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZBookRepoRelationServer is the server API for ZBookRepoRelation service.
// All implementations must embed UnimplementedZBookRepoRelationServer
// for forward compatibility
type ZBookRepoRelationServer interface {
	// 1.CreateRepoRelation
	CreateRepoRelation(context.Context, *rpcs.CreateRepoRelationRequest) (*rpcs.CreateRepoRelationResponse, error)
	// 2.DeleteRepoRelation
	DeleteRepoRelation(context.Context, *rpcs.DeleteRepoRelationRequest) (*rpcs.DeleteRepoRelationResponse, error)
	// 3.CreateRepoVisibility
	CreateRepoVisibility(context.Context, *rpcs.CreateRepoVisibilityRequest) (*rpcs.CreateRepoVisibilityResponse, error)
	// 4.DeleteRepoVisibility
	DeleteRepoVisibility(context.Context, *rpcs.DeleteRepoVisibilityRequest) (*rpcs.DeleteRepoVisibilityResponse, error)
	// 5.ListSelectedUserByRepo
	ListSelectedUserByRepo(context.Context, *rpcs.ListSelectedUserByRepoRequest) (*rpcs.ListSelectedUserByRepoResponse, error)
	// 6.GetSelectedUserByRepoCount
	GetSelectedUserByRepoCount(context.Context, *rpcs.GetSelectedUserByRepoCountRequest) (*rpcs.GetSelectedUserByRepoCountResponse, error)
	// 7.QueryUserByRepo
	QueryUserByRepo(context.Context, *rpcs.QueryUserByRepoRequest) (*rpcs.QueryUserByRepoResponse, error)
	mustEmbedUnimplementedZBookRepoRelationServer()
}

// UnimplementedZBookRepoRelationServer must be embedded to have forward compatible implementations.
type UnimplementedZBookRepoRelationServer struct {
}

func (UnimplementedZBookRepoRelationServer) CreateRepoRelation(context.Context, *rpcs.CreateRepoRelationRequest) (*rpcs.CreateRepoRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepoRelation not implemented")
}
func (UnimplementedZBookRepoRelationServer) DeleteRepoRelation(context.Context, *rpcs.DeleteRepoRelationRequest) (*rpcs.DeleteRepoRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepoRelation not implemented")
}
func (UnimplementedZBookRepoRelationServer) CreateRepoVisibility(context.Context, *rpcs.CreateRepoVisibilityRequest) (*rpcs.CreateRepoVisibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepoVisibility not implemented")
}
func (UnimplementedZBookRepoRelationServer) DeleteRepoVisibility(context.Context, *rpcs.DeleteRepoVisibilityRequest) (*rpcs.DeleteRepoVisibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepoVisibility not implemented")
}
func (UnimplementedZBookRepoRelationServer) ListSelectedUserByRepo(context.Context, *rpcs.ListSelectedUserByRepoRequest) (*rpcs.ListSelectedUserByRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSelectedUserByRepo not implemented")
}
func (UnimplementedZBookRepoRelationServer) GetSelectedUserByRepoCount(context.Context, *rpcs.GetSelectedUserByRepoCountRequest) (*rpcs.GetSelectedUserByRepoCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelectedUserByRepoCount not implemented")
}
func (UnimplementedZBookRepoRelationServer) QueryUserByRepo(context.Context, *rpcs.QueryUserByRepoRequest) (*rpcs.QueryUserByRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserByRepo not implemented")
}
func (UnimplementedZBookRepoRelationServer) mustEmbedUnimplementedZBookRepoRelationServer() {}

// UnsafeZBookRepoRelationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZBookRepoRelationServer will
// result in compilation errors.
type UnsafeZBookRepoRelationServer interface {
	mustEmbedUnimplementedZBookRepoRelationServer()
}

func RegisterZBookRepoRelationServer(s grpc.ServiceRegistrar, srv ZBookRepoRelationServer) {
	s.RegisterService(&ZBookRepoRelation_ServiceDesc, srv)
}

func _ZBookRepoRelation_CreateRepoRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.CreateRepoRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).CreateRepoRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_CreateRepoRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).CreateRepoRelation(ctx, req.(*rpcs.CreateRepoRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookRepoRelation_DeleteRepoRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.DeleteRepoRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).DeleteRepoRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_DeleteRepoRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).DeleteRepoRelation(ctx, req.(*rpcs.DeleteRepoRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookRepoRelation_CreateRepoVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.CreateRepoVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).CreateRepoVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_CreateRepoVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).CreateRepoVisibility(ctx, req.(*rpcs.CreateRepoVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookRepoRelation_DeleteRepoVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.DeleteRepoVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).DeleteRepoVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_DeleteRepoVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).DeleteRepoVisibility(ctx, req.(*rpcs.DeleteRepoVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookRepoRelation_ListSelectedUserByRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.ListSelectedUserByRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).ListSelectedUserByRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_ListSelectedUserByRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).ListSelectedUserByRepo(ctx, req.(*rpcs.ListSelectedUserByRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookRepoRelation_GetSelectedUserByRepoCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.GetSelectedUserByRepoCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).GetSelectedUserByRepoCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_GetSelectedUserByRepoCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).GetSelectedUserByRepoCount(ctx, req.(*rpcs.GetSelectedUserByRepoCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZBookRepoRelation_QueryUserByRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpcs.QueryUserByRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZBookRepoRelationServer).QueryUserByRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZBookRepoRelation_QueryUserByRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZBookRepoRelationServer).QueryUserByRepo(ctx, req.(*rpcs.QueryUserByRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZBookRepoRelation_ServiceDesc is the grpc.ServiceDesc for ZBookRepoRelation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZBookRepoRelation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ZBookRepoRelation",
	HandlerType: (*ZBookRepoRelationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepoRelation",
			Handler:    _ZBookRepoRelation_CreateRepoRelation_Handler,
		},
		{
			MethodName: "DeleteRepoRelation",
			Handler:    _ZBookRepoRelation_DeleteRepoRelation_Handler,
		},
		{
			MethodName: "CreateRepoVisibility",
			Handler:    _ZBookRepoRelation_CreateRepoVisibility_Handler,
		},
		{
			MethodName: "DeleteRepoVisibility",
			Handler:    _ZBookRepoRelation_DeleteRepoVisibility_Handler,
		},
		{
			MethodName: "ListSelectedUserByRepo",
			Handler:    _ZBookRepoRelation_ListSelectedUserByRepo_Handler,
		},
		{
			MethodName: "GetSelectedUserByRepoCount",
			Handler:    _ZBookRepoRelation_GetSelectedUserByRepoCount_Handler,
		},
		{
			MethodName: "QueryUserByRepo",
			Handler:    _ZBookRepoRelation_QueryUserByRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_zbook_repo_relation.proto",
}
