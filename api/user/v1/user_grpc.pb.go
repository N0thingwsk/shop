// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: user/v1/user.proto

package v1

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	UpdateUserNickName(ctx context.Context, in *UpdateNickNameRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	UpdateUserPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	UpdateUserBirthday(ctx context.Context, in *UpdateBirthdayRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	UpdateUserGender(ctx context.Context, in *UpdateGenderRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	SendVerificationCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserNickName(ctx context.Context, in *UpdateNickNameRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/UpdateUserNickName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/UpdateUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserBirthday(ctx context.Context, in *UpdateBirthdayRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/UpdateUserBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserGender(ctx context.Context, in *UpdateGenderRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/UpdateUserGender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendVerificationCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeReply, error) {
	out := new(SendCodeReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/SendVerificationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetUser(context.Context, *GetUserRequest) (*UserInfoReply, error)
	LoginUser(context.Context, *LoginRequest) (*LoginReply, error)
	RegisterUser(context.Context, *RegisterRequest) (*UserInfoReply, error)
	UpdateUserNickName(context.Context, *UpdateNickNameRequest) (*UserInfoReply, error)
	UpdateUserPassword(context.Context, *UpdatePasswordRequest) (*UserInfoReply, error)
	UpdateUserBirthday(context.Context, *UpdateBirthdayRequest) (*UserInfoReply, error)
	UpdateUserGender(context.Context, *UpdateGenderRequest) (*UserInfoReply, error)
	DeleteUser(context.Context, *DeleteRequest) (*UserInfoReply, error)
	SendVerificationCode(context.Context, *SendCodeRequest) (*SendCodeReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) LoginUser(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserServer) RegisterUser(context.Context, *RegisterRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServer) UpdateUserNickName(context.Context, *UpdateNickNameRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserNickName not implemented")
}
func (UnimplementedUserServer) UpdateUserPassword(context.Context, *UpdatePasswordRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
func (UnimplementedUserServer) UpdateUserBirthday(context.Context, *UpdateBirthdayRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBirthday not implemented")
}
func (UnimplementedUserServer) UpdateUserGender(context.Context, *UpdateGenderRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGender not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *DeleteRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) SendVerificationCode(context.Context, *SendCodeRequest) (*SendCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationCode not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserNickName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNickNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserNickName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/UpdateUserNickName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserNickName(ctx, req.(*UpdateNickNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/UpdateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserPassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBirthdayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/UpdateUserBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserBirthday(ctx, req.(*UpdateBirthdayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserGender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserGender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/UpdateUserGender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserGender(ctx, req.(*UpdateGenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/SendVerificationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendVerificationCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _User_LoginUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _User_RegisterUser_Handler,
		},
		{
			MethodName: "UpdateUserNickName",
			Handler:    _User_UpdateUserNickName_Handler,
		},
		{
			MethodName: "UpdateUserPassword",
			Handler:    _User_UpdateUserPassword_Handler,
		},
		{
			MethodName: "UpdateUserBirthday",
			Handler:    _User_UpdateUserBirthday_Handler,
		},
		{
			MethodName: "UpdateUserGender",
			Handler:    _User_UpdateUserGender_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "SendVerificationCode",
			Handler:    _User_SendVerificationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
