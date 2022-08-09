package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "shop/api/user/v1"
	"shop/app/user/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUsecase
	log *log.Helper
}
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) GetUserList(context context.Context, in *v1.PageInfo) (*v1.UserListResponse, error) {
	return nil, nil
}

func (u *UserService) GetUserByMobile(context context.Context, in *v1.MobileRequest) (*v1.UserInfoResponse, error) {
	return nil, nil
}

func (u *UserService) GetUserByID(context.Context, *v1.IdRequest) (*v1.UserInfoResponse, error) {
	return nil, nil
}

func (u *UserService) CreateUser(context.Context, *v1.CreateUserInfo) (*v1.UserInfoResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(context.Context, *v1.UpdateUserInfo) (*emptypb.Empty, error) {
	return nil, nil
}

func (u *UserService) CheckPassWorld(context.Context, *v1.PassWordCheckInfo) (*v1.CheckResponse, error) {
	return nil, nil
}