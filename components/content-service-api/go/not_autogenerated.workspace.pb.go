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

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	// WorkspaceDownloadURL provides a URL from where the content of a workspace can be downloaded from
	WorkspaceDownloadURL(ctx context.Context, in *WorkspaceDownloadURLRequest, opts ...grpc.CallOption) (*WorkspaceDownloadURLResponse, error)
	// DeleteWorkspace deletes the content of a single workspace
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) WorkspaceDownloadURL(ctx context.Context, in *WorkspaceDownloadURLRequest, opts ...grpc.CallOption) (*WorkspaceDownloadURLResponse, error) {
	out := new(WorkspaceDownloadURLResponse)
	err := c.cc.Invoke(ctx, "/contentservice.WorkspaceService/WorkspaceDownloadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/contentservice.WorkspaceService/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
type WorkspaceServiceServer interface {
	// WorkspaceDownloadURL provides a URL from where the content of a workspace can be downloaded from
	WorkspaceDownloadURL(context.Context, *WorkspaceDownloadURLRequest) (*WorkspaceDownloadURLResponse, error)
	// DeleteWorkspace deletes the content of a single workspace
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
}

// UnimplementedWorkspaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (*UnimplementedWorkspaceServiceServer) WorkspaceDownloadURL(ctx context.Context, req *WorkspaceDownloadURLRequest) (*WorkspaceDownloadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkspaceDownloadURL not implemented")
}
func (*UnimplementedWorkspaceServiceServer) DeleteWorkspace(ctx context.Context, req *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}

func RegisterWorkspaceServiceServer(s *grpc.Server, srv WorkspaceServiceServer) {
	s.RegisterService(&_WorkspaceService_serviceDesc, srv)
}

func _WorkspaceService_WorkspaceDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkspaceDownloadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).WorkspaceDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.WorkspaceService/WorkspaceDownloadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).WorkspaceDownloadURL(ctx, req.(*WorkspaceDownloadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.WorkspaceService/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkspaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkspaceDownloadURL",
			Handler:    _WorkspaceService_WorkspaceDownloadURL_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace.proto",
}
