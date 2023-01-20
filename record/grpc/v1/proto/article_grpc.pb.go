// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: v1/proto/article.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	// 新建文章
	CreateArticle(ctx context.Context, in *RequestCreateArticle, opts ...grpc.CallOption) (*ReplyCreateArticle, error)
	// 修改文章
	UpdateArticle(ctx context.Context, in *RequestUpdateArticle, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除文章
	DeleteArticle(ctx context.Context, in *RequestDeleteArticle, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 查询文章
	QueryArticle(ctx context.Context, in *RequestQueryArticle, opts ...grpc.CallOption) (*ReplyQueryArticle, error)
	// 文章列表
	ArticleList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReplyArticleList, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *RequestCreateArticle, opts ...grpc.CallOption) (*ReplyCreateArticle, error) {
	out := new(ReplyCreateArticle)
	err := c.cc.Invoke(ctx, "/article.v1.ArticleService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticle(ctx context.Context, in *RequestUpdateArticle, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/article.v1.ArticleService/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticle(ctx context.Context, in *RequestDeleteArticle, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/article.v1.ArticleService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) QueryArticle(ctx context.Context, in *RequestQueryArticle, opts ...grpc.CallOption) (*ReplyQueryArticle, error) {
	out := new(ReplyQueryArticle)
	err := c.cc.Invoke(ctx, "/article.v1.ArticleService/QueryArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReplyArticleList, error) {
	out := new(ReplyArticleList)
	err := c.cc.Invoke(ctx, "/article.v1.ArticleService/ArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	// 新建文章
	CreateArticle(context.Context, *RequestCreateArticle) (*ReplyCreateArticle, error)
	// 修改文章
	UpdateArticle(context.Context, *RequestUpdateArticle) (*emptypb.Empty, error)
	// 删除文章
	DeleteArticle(context.Context, *RequestDeleteArticle) (*emptypb.Empty, error)
	// 查询文章
	QueryArticle(context.Context, *RequestQueryArticle) (*ReplyQueryArticle, error)
	// 文章列表
	ArticleList(context.Context, *emptypb.Empty) (*ReplyArticleList, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *RequestCreateArticle) (*ReplyCreateArticle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticle(context.Context, *RequestUpdateArticle) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticle(context.Context, *RequestDeleteArticle) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServiceServer) QueryArticle(context.Context, *RequestQueryArticle) (*ReplyQueryArticle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryArticle not implemented")
}
func (UnimplementedArticleServiceServer) ArticleList(context.Context, *emptypb.Empty) (*ReplyArticleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleList not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.v1.ArticleService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*RequestCreateArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.v1.ArticleService/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, req.(*RequestUpdateArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.v1.ArticleService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, req.(*RequestDeleteArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_QueryArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestQueryArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).QueryArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.v1.ArticleService/QueryArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).QueryArticle(ctx, req.(*RequestQueryArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.v1.ArticleService/ArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article.v1.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticleService_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticleService_DeleteArticle_Handler,
		},
		{
			MethodName: "QueryArticle",
			Handler:    _ArticleService_QueryArticle_Handler,
		},
		{
			MethodName: "ArticleList",
			Handler:    _ArticleService_ArticleList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/proto/article.proto",
}
