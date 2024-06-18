// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: todo_item.proto

package pb

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

// TodoHandlerClient is the client API for TodoHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoHandlerClient interface {
	CreateTodoItem(ctx context.Context, in *CreateTodoItemRequest, opts ...grpc.CallOption) (*TodoItemResponse, error)
	BatchCreateTodoItems(ctx context.Context, opts ...grpc.CallOption) (TodoHandler_BatchCreateTodoItemsClient, error)
	GetTodoItem(ctx context.Context, in *ItemByIdRequest, opts ...grpc.CallOption) (*TodoItemResponse, error)
	ListTodoItems(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (TodoHandler_ListTodoItemsClient, error)
	SearchTodoItems(ctx context.Context, in *ItemSearchRequest, opts ...grpc.CallOption) (TodoHandler_SearchTodoItemsClient, error)
	UpdateTodoItem(ctx context.Context, in *UpdateTodoItemRequest, opts ...grpc.CallOption) (*TodoItemResponse, error)
	DeleteTodoItem(ctx context.Context, in *ItemByIdRequest, opts ...grpc.CallOption) (*DeletedStatusResponse, error)
	ChangeTodoState(ctx context.Context, in *ChangeTodoStateRequest, opts ...grpc.CallOption) (*TodoItemResponse, error)
}

type todoHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoHandlerClient(cc grpc.ClientConnInterface) TodoHandlerClient {
	return &todoHandlerClient{cc}
}

func (c *todoHandlerClient) CreateTodoItem(ctx context.Context, in *CreateTodoItemRequest, opts ...grpc.CallOption) (*TodoItemResponse, error) {
	out := new(TodoItemResponse)
	err := c.cc.Invoke(ctx, "/todoitem.TodoHandler/CreateTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoHandlerClient) BatchCreateTodoItems(ctx context.Context, opts ...grpc.CallOption) (TodoHandler_BatchCreateTodoItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoHandler_ServiceDesc.Streams[0], "/todoitem.TodoHandler/BatchCreateTodoItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoHandlerBatchCreateTodoItemsClient{stream}
	return x, nil
}

type TodoHandler_BatchCreateTodoItemsClient interface {
	Send(*BatchCreateTodoItemsRequest) error
	Recv() (*TodoItemResponse, error)
	grpc.ClientStream
}

type todoHandlerBatchCreateTodoItemsClient struct {
	grpc.ClientStream
}

func (x *todoHandlerBatchCreateTodoItemsClient) Send(m *BatchCreateTodoItemsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *todoHandlerBatchCreateTodoItemsClient) Recv() (*TodoItemResponse, error) {
	m := new(TodoItemResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoHandlerClient) GetTodoItem(ctx context.Context, in *ItemByIdRequest, opts ...grpc.CallOption) (*TodoItemResponse, error) {
	out := new(TodoItemResponse)
	err := c.cc.Invoke(ctx, "/todoitem.TodoHandler/GetTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoHandlerClient) ListTodoItems(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (TodoHandler_ListTodoItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoHandler_ServiceDesc.Streams[1], "/todoitem.TodoHandler/ListTodoItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoHandlerListTodoItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoHandler_ListTodoItemsClient interface {
	Recv() (*TodoItemResponse, error)
	grpc.ClientStream
}

type todoHandlerListTodoItemsClient struct {
	grpc.ClientStream
}

func (x *todoHandlerListTodoItemsClient) Recv() (*TodoItemResponse, error) {
	m := new(TodoItemResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoHandlerClient) SearchTodoItems(ctx context.Context, in *ItemSearchRequest, opts ...grpc.CallOption) (TodoHandler_SearchTodoItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoHandler_ServiceDesc.Streams[2], "/todoitem.TodoHandler/SearchTodoItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoHandlerSearchTodoItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoHandler_SearchTodoItemsClient interface {
	Recv() (*TodoItemResponse, error)
	grpc.ClientStream
}

type todoHandlerSearchTodoItemsClient struct {
	grpc.ClientStream
}

func (x *todoHandlerSearchTodoItemsClient) Recv() (*TodoItemResponse, error) {
	m := new(TodoItemResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoHandlerClient) UpdateTodoItem(ctx context.Context, in *UpdateTodoItemRequest, opts ...grpc.CallOption) (*TodoItemResponse, error) {
	out := new(TodoItemResponse)
	err := c.cc.Invoke(ctx, "/todoitem.TodoHandler/UpdateTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoHandlerClient) DeleteTodoItem(ctx context.Context, in *ItemByIdRequest, opts ...grpc.CallOption) (*DeletedStatusResponse, error) {
	out := new(DeletedStatusResponse)
	err := c.cc.Invoke(ctx, "/todoitem.TodoHandler/DeleteTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoHandlerClient) ChangeTodoState(ctx context.Context, in *ChangeTodoStateRequest, opts ...grpc.CallOption) (*TodoItemResponse, error) {
	out := new(TodoItemResponse)
	err := c.cc.Invoke(ctx, "/todoitem.TodoHandler/ChangeTodoState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoHandlerServer is the server API for TodoHandler service.
// All implementations must embed UnimplementedTodoHandlerServer
// for forward compatibility
type TodoHandlerServer interface {
	CreateTodoItem(context.Context, *CreateTodoItemRequest) (*TodoItemResponse, error)
	BatchCreateTodoItems(TodoHandler_BatchCreateTodoItemsServer) error
	GetTodoItem(context.Context, *ItemByIdRequest) (*TodoItemResponse, error)
	ListTodoItems(*ListTodoRequest, TodoHandler_ListTodoItemsServer) error
	SearchTodoItems(*ItemSearchRequest, TodoHandler_SearchTodoItemsServer) error
	UpdateTodoItem(context.Context, *UpdateTodoItemRequest) (*TodoItemResponse, error)
	DeleteTodoItem(context.Context, *ItemByIdRequest) (*DeletedStatusResponse, error)
	ChangeTodoState(context.Context, *ChangeTodoStateRequest) (*TodoItemResponse, error)
	mustEmbedUnimplementedTodoHandlerServer()
}

// UnimplementedTodoHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedTodoHandlerServer struct {
}

func (UnimplementedTodoHandlerServer) CreateTodoItem(context.Context, *CreateTodoItemRequest) (*TodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodoItem not implemented")
}
func (UnimplementedTodoHandlerServer) BatchCreateTodoItems(TodoHandler_BatchCreateTodoItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchCreateTodoItems not implemented")
}
func (UnimplementedTodoHandlerServer) GetTodoItem(context.Context, *ItemByIdRequest) (*TodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoItem not implemented")
}
func (UnimplementedTodoHandlerServer) ListTodoItems(*ListTodoRequest, TodoHandler_ListTodoItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTodoItems not implemented")
}
func (UnimplementedTodoHandlerServer) SearchTodoItems(*ItemSearchRequest, TodoHandler_SearchTodoItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchTodoItems not implemented")
}
func (UnimplementedTodoHandlerServer) UpdateTodoItem(context.Context, *UpdateTodoItemRequest) (*TodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodoItem not implemented")
}
func (UnimplementedTodoHandlerServer) DeleteTodoItem(context.Context, *ItemByIdRequest) (*DeletedStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoItem not implemented")
}
func (UnimplementedTodoHandlerServer) ChangeTodoState(context.Context, *ChangeTodoStateRequest) (*TodoItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTodoState not implemented")
}
func (UnimplementedTodoHandlerServer) mustEmbedUnimplementedTodoHandlerServer() {}

// UnsafeTodoHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoHandlerServer will
// result in compilation errors.
type UnsafeTodoHandlerServer interface {
	mustEmbedUnimplementedTodoHandlerServer()
}

func RegisterTodoHandlerServer(s grpc.ServiceRegistrar, srv TodoHandlerServer) {
	s.RegisterService(&TodoHandler_ServiceDesc, srv)
}

func _TodoHandler_CreateTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoHandlerServer).CreateTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoitem.TodoHandler/CreateTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoHandlerServer).CreateTodoItem(ctx, req.(*CreateTodoItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoHandler_BatchCreateTodoItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TodoHandlerServer).BatchCreateTodoItems(&todoHandlerBatchCreateTodoItemsServer{stream})
}

type TodoHandler_BatchCreateTodoItemsServer interface {
	Send(*TodoItemResponse) error
	Recv() (*BatchCreateTodoItemsRequest, error)
	grpc.ServerStream
}

type todoHandlerBatchCreateTodoItemsServer struct {
	grpc.ServerStream
}

func (x *todoHandlerBatchCreateTodoItemsServer) Send(m *TodoItemResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *todoHandlerBatchCreateTodoItemsServer) Recv() (*BatchCreateTodoItemsRequest, error) {
	m := new(BatchCreateTodoItemsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TodoHandler_GetTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoHandlerServer).GetTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoitem.TodoHandler/GetTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoHandlerServer).GetTodoItem(ctx, req.(*ItemByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoHandler_ListTodoItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTodoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoHandlerServer).ListTodoItems(m, &todoHandlerListTodoItemsServer{stream})
}

type TodoHandler_ListTodoItemsServer interface {
	Send(*TodoItemResponse) error
	grpc.ServerStream
}

type todoHandlerListTodoItemsServer struct {
	grpc.ServerStream
}

func (x *todoHandlerListTodoItemsServer) Send(m *TodoItemResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TodoHandler_SearchTodoItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoHandlerServer).SearchTodoItems(m, &todoHandlerSearchTodoItemsServer{stream})
}

type TodoHandler_SearchTodoItemsServer interface {
	Send(*TodoItemResponse) error
	grpc.ServerStream
}

type todoHandlerSearchTodoItemsServer struct {
	grpc.ServerStream
}

func (x *todoHandlerSearchTodoItemsServer) Send(m *TodoItemResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TodoHandler_UpdateTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoHandlerServer).UpdateTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoitem.TodoHandler/UpdateTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoHandlerServer).UpdateTodoItem(ctx, req.(*UpdateTodoItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoHandler_DeleteTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoHandlerServer).DeleteTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoitem.TodoHandler/DeleteTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoHandlerServer).DeleteTodoItem(ctx, req.(*ItemByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoHandler_ChangeTodoState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTodoStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoHandlerServer).ChangeTodoState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoitem.TodoHandler/ChangeTodoState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoHandlerServer).ChangeTodoState(ctx, req.(*ChangeTodoStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoHandler_ServiceDesc is the grpc.ServiceDesc for TodoHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todoitem.TodoHandler",
	HandlerType: (*TodoHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodoItem",
			Handler:    _TodoHandler_CreateTodoItem_Handler,
		},
		{
			MethodName: "GetTodoItem",
			Handler:    _TodoHandler_GetTodoItem_Handler,
		},
		{
			MethodName: "UpdateTodoItem",
			Handler:    _TodoHandler_UpdateTodoItem_Handler,
		},
		{
			MethodName: "DeleteTodoItem",
			Handler:    _TodoHandler_DeleteTodoItem_Handler,
		},
		{
			MethodName: "ChangeTodoState",
			Handler:    _TodoHandler_ChangeTodoState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BatchCreateTodoItems",
			Handler:       _TodoHandler_BatchCreateTodoItems_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListTodoItems",
			Handler:       _TodoHandler_ListTodoItems_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchTodoItems",
			Handler:       _TodoHandler_SearchTodoItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo_item.proto",
}
