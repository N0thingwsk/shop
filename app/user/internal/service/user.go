package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "shop/api/user/v1"
	"shop/app/user/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.UserInfoReply, error) {
	fmt.Println("test")
	return &v1.UserInfoReply{}, nil
}

func (u *UserService) LoginUser(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	token, err := u.uc.LoginUser(ctx, biz.User{
		Mobile:   in.User.Mobile,
		Password: in.User.Password,
	})
	if err != nil {
		return &v1.LoginReply{}, err
	}
	return &v1.LoginReply{
		User: &v1.LoginReply_User{
			Mobile: in.User.Mobile,
		},
		Token: token,
	}, nil
}

func (u *UserService) RegisterUser(ctx context.Context, in *v1.RegisterRequest) (*v1.UserInfoReply, error) {

	return nil, nil
}

func (u *UserService) UpdateUserNickName(ctx context.Context, in *v1.UpdateNickNameRequest) (*v1.UserInfoReply, error) {
	return nil, nil
}

func (u *UserService) UpdateUserPassword(ctx context.Context, in *v1.UpdatePasswordRequest) (*v1.UserInfoReply, error) {
	_, err := u.uc.UpdateUserPassword(ctx, biz.User{
		Model: gorm.Model{
			ID: ctx.Value("userid").(uint),
		},
		Password: in.User.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{}, nil
}

func (u *UserService) UpdateUserBirthday(ctx context.Context, in *v1.UpdateBirthdayRequest) (*v1.UserInfoReply, error) {
	return nil, nil
}

func (u *UserService) UpdateUserGender(ctx context.Context, in *v1.UpdateGenderRequest) (*v1.UserInfoReply, error) {
	return nil, nil
}

func (u *UserService) DeleteUser(ctx context.Context, in *v1.DeleteRequest) (*v1.UserInfoReply, error) {
	return nil, nil
}

func (u *UserService) SendVerificationCode(ctx context.Context, in *v1.SendCodeRequest) (*v1.SendCodeReply, error) {
	return nil, nil
}
