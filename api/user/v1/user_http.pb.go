// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.2
// source: user/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserDeleteUser = "/api.user.v1.User/DeleteUser"
const OperationUserGetUser = "/api.user.v1.User/GetUser"
const OperationUserLoginUser = "/api.user.v1.User/LoginUser"
const OperationUserRegisterUser = "/api.user.v1.User/RegisterUser"
const OperationUserSendVerificationCode = "/api.user.v1.User/SendVerificationCode"
const OperationUserUpdateUserBirthday = "/api.user.v1.User/UpdateUserBirthday"
const OperationUserUpdateUserGender = "/api.user.v1.User/UpdateUserGender"
const OperationUserUpdateUserNickName = "/api.user.v1.User/UpdateUserNickName"
const OperationUserUpdateUserPassword = "/api.user.v1.User/UpdateUserPassword"

type UserHTTPServer interface {
	DeleteUser(context.Context, *DeleteRequest) (*UserInfoReply, error)
	GetUser(context.Context, *GetUserRequest) (*UserInfoReply, error)
	LoginUser(context.Context, *LoginRequest) (*LoginReply, error)
	RegisterUser(context.Context, *RegisterRequest) (*UserInfoReply, error)
	SendVerificationCode(context.Context, *SendCodeRequest) (*SendCodeReply, error)
	UpdateUserBirthday(context.Context, *UpdateBirthdayRequest) (*UserInfoReply, error)
	UpdateUserGender(context.Context, *UpdateGenderRequest) (*UserInfoReply, error)
	UpdateUserNickName(context.Context, *UpdateNickNameRequest) (*UserInfoReply, error)
	UpdateUserPassword(context.Context, *UpdatePasswordRequest) (*UserInfoReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user/getuser", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/v1/user/login", _User_LoginUser0_HTTP_Handler(srv))
	r.POST("/v1/user/create", _User_RegisterUser0_HTTP_Handler(srv))
	r.POST("/v1/user/update/nickname", _User_UpdateUserNickName0_HTTP_Handler(srv))
	r.POST("/v1/user/update/password", _User_UpdateUserPassword0_HTTP_Handler(srv))
	r.POST("/v1/user/update/birthday", _User_UpdateUserBirthday0_HTTP_Handler(srv))
	r.POST("/v1/user/update/gender", _User_UpdateUserGender0_HTTP_Handler(srv))
	r.POST("/v1/user/delete", _User_DeleteUser0_HTTP_Handler(srv))
	r.POST("/v1/user/sendcode", _User_SendVerificationCode0_HTTP_Handler(srv))
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_LoginUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginUser(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_RegisterUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegisterUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterUser(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserNickName0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateNickNameRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserNickName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserNickName(ctx, req.(*UpdateNickNameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserPassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserPassword(ctx, req.(*UpdatePasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserBirthday0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBirthdayRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserBirthday)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserBirthday(ctx, req.(*UpdateBirthdayRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserGender0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateGenderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserGender)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserGender(ctx, req.(*UpdateGenderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_SendVerificationCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSendVerificationCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendVerificationCode(ctx, req.(*SendCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendCodeReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	DeleteUser(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	LoginUser(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	RegisterUser(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	SendVerificationCode(ctx context.Context, req *SendCodeRequest, opts ...http.CallOption) (rsp *SendCodeReply, err error)
	UpdateUserBirthday(ctx context.Context, req *UpdateBirthdayRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	UpdateUserGender(ctx context.Context, req *UpdateGenderRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	UpdateUserNickName(ctx context.Context, req *UpdateNickNameRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	UpdateUserPassword(ctx context.Context, req *UpdatePasswordRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/getuser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) LoginUser(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegisterUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SendVerificationCode(ctx context.Context, in *SendCodeRequest, opts ...http.CallOption) (*SendCodeReply, error) {
	var out SendCodeReply
	pattern := "/v1/user/sendcode"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSendVerificationCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserBirthday(ctx context.Context, in *UpdateBirthdayRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/update/birthday"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserBirthday))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserGender(ctx context.Context, in *UpdateGenderRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/update/gender"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserGender))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserNickName(ctx context.Context, in *UpdateNickNameRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/update/nickname"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserNickName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/v1/user/update/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
