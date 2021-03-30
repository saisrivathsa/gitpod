// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package api

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentServiceClient interface {
	// DeleteUserContent deletes all content associated with a user.
	DeleteUserContent(ctx context.Context, in *DeleteUserContentRequest, opts ...grpc.CallOption) (*DeleteUserContentResponse, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) DeleteUserContent(ctx context.Context, in *DeleteUserContentRequest, opts ...grpc.CallOption) (*DeleteUserContentResponse, error) {
	out := new(DeleteUserContentResponse)
	err := c.cc.Invoke(ctx, "/contentservice.ContentService/DeleteUserContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
type ContentServiceServer interface {
	// DeleteUserContent deletes all content associated with a user.
	DeleteUserContent(context.Context, *DeleteUserContentRequest) (*DeleteUserContentResponse, error)
}

// UnimplementedContentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (*UnimplementedContentServiceServer) DeleteUserContent(ctx context.Context, req *DeleteUserContentRequest) (*DeleteUserContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserContent not implemented")
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_DeleteUserContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteUserContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.ContentService/DeleteUserContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteUserContent(ctx, req.(*DeleteUserContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserContent",
			Handler:    _ContentService_DeleteUserContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}
