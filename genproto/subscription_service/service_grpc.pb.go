// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: service.proto

package subscription_service

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
	ResourceCategoryService_Create_FullMethodName = "/subscription_service.ResourceCategoryService/Create"
	ResourceCategoryService_Get_FullMethodName    = "/subscription_service.ResourceCategoryService/Get"
	ResourceCategoryService_Find_FullMethodName   = "/subscription_service.ResourceCategoryService/Find"
	ResourceCategoryService_Update_FullMethodName = "/subscription_service.ResourceCategoryService/Update"
	ResourceCategoryService_Delete_FullMethodName = "/subscription_service.ResourceCategoryService/Delete"
)

// ResourceCategoryServiceClient is the client API for ResourceCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceCategoryServiceClient interface {
	Create(ctx context.Context, in *ResourceCategory, opts ...grpc.CallOption) (*ResourceCategory, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResourceCategory, error)
	Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*ResourceCategories, error)
	Update(ctx context.Context, in *ResourceCategory, opts ...grpc.CallOption) (*ResourceCategory, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type resourceCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceCategoryServiceClient(cc grpc.ClientConnInterface) ResourceCategoryServiceClient {
	return &resourceCategoryServiceClient{cc}
}

func (c *resourceCategoryServiceClient) Create(ctx context.Context, in *ResourceCategory, opts ...grpc.CallOption) (*ResourceCategory, error) {
	out := new(ResourceCategory)
	err := c.cc.Invoke(ctx, ResourceCategoryService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceCategoryServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResourceCategory, error) {
	out := new(ResourceCategory)
	err := c.cc.Invoke(ctx, ResourceCategoryService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceCategoryServiceClient) Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*ResourceCategories, error) {
	out := new(ResourceCategories)
	err := c.cc.Invoke(ctx, ResourceCategoryService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceCategoryServiceClient) Update(ctx context.Context, in *ResourceCategory, opts ...grpc.CallOption) (*ResourceCategory, error) {
	out := new(ResourceCategory)
	err := c.cc.Invoke(ctx, ResourceCategoryService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceCategoryServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ResourceCategoryService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceCategoryServiceServer is the server API for ResourceCategoryService service.
// All implementations must embed UnimplementedResourceCategoryServiceServer
// for forward compatibility
type ResourceCategoryServiceServer interface {
	Create(context.Context, *ResourceCategory) (*ResourceCategory, error)
	Get(context.Context, *Id) (*ResourceCategory, error)
	Find(context.Context, *GetListFilter) (*ResourceCategories, error)
	Update(context.Context, *ResourceCategory) (*ResourceCategory, error)
	Delete(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedResourceCategoryServiceServer()
}

// UnimplementedResourceCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceCategoryServiceServer struct {
}

func (UnimplementedResourceCategoryServiceServer) Create(context.Context, *ResourceCategory) (*ResourceCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourceCategoryServiceServer) Get(context.Context, *Id) (*ResourceCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResourceCategoryServiceServer) Find(context.Context, *GetListFilter) (*ResourceCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedResourceCategoryServiceServer) Update(context.Context, *ResourceCategory) (*ResourceCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResourceCategoryServiceServer) Delete(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResourceCategoryServiceServer) mustEmbedUnimplementedResourceCategoryServiceServer() {
}

// UnsafeResourceCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceCategoryServiceServer will
// result in compilation errors.
type UnsafeResourceCategoryServiceServer interface {
	mustEmbedUnimplementedResourceCategoryServiceServer()
}

func RegisterResourceCategoryServiceServer(s grpc.ServiceRegistrar, srv ResourceCategoryServiceServer) {
	s.RegisterService(&ResourceCategoryService_ServiceDesc, srv)
}

func _ResourceCategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceCategoryService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCategoryServiceServer).Create(ctx, req.(*ResourceCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceCategoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCategoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceCategoryService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCategoryServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceCategoryService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCategoryServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceCategoryService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCategoryServiceServer).Find(ctx, req.(*GetListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceCategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceCategoryService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCategoryServiceServer).Update(ctx, req.(*ResourceCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceCategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceCategoryService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCategoryServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceCategoryService_ServiceDesc is the grpc.ServiceDesc for ResourceCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription_service.ResourceCategoryService",
	HandlerType: (*ResourceCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResourceCategoryService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResourceCategoryService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _ResourceCategoryService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceCategoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceCategoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}