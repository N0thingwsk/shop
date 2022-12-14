package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "shop/api/user/v1"
	"shop/app/user/internal/biz"
	"time"
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
	user, err := u.uc.GetUserinfo(ctx, int(ctx.Value("userid").(float64)))
	if err != nil {
		return &v1.UserInfoReply{}, err
	}
	return &v1.UserInfoReply{
		User: &v1.UserInfoReply_User{
			Mobile:   user.Mobile,
			NikeName: user.NickName,
			Birthday: user.Birthday.Format("2006-01-02 15:04:05"),
			Gender:   user.Gender,
		},
	}, nil
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
	user, err := u.uc.CreateUser(ctx, biz.User{
		Mobile:   in.User.Mobile,
		Password: in.User.Password,
	})
	if err != nil {
		return &v1.UserInfoReply{}, err
	}
	return &v1.UserInfoReply{
		User: &v1.UserInfoReply_User{
			Mobile: user.Mobile,
		},
	}, nil
}

func (u *UserService) UpdateUserNickName(ctx context.Context, in *v1.UpdateNickNameRequest) (*v1.UserInfoReply, error) {
	user, err := u.uc.UpdateUserNickName(ctx, biz.User{
		Model: gorm.Model{
			ID: uint(ctx.Value("userid").(float64)),
		},
		NickName: in.User.Nikename,
	})
	if err != nil {
		return &v1.UserInfoReply{}, err
	}
	return &v1.UserInfoReply{User: &v1.UserInfoReply_User{
		NikeName: user.NickName,
	}}, nil
}

func (u *UserService) UpdateUserPassword(ctx context.Context, in *v1.UpdatePasswordRequest) (*v1.UserInfoReply, error) {
	user, err := u.uc.UpdateUserPassword(ctx, biz.User{
		Model: gorm.Model{
			ID: uint(ctx.Value("userid").(float64)),
		},
		Password: in.User.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{User: &v1.UserInfoReply_User{
		Mobile: user.Mobile,
	},
	}, nil
}

func (u *UserService) UpdateUserBirthday(ctx context.Context, in *v1.UpdateBirthdayRequest) (*v1.UserInfoReply, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", in.User.Birthday, time.Local)
	if err != nil {
		return nil, err
	}
	user, err := u.uc.UpdateUserBirthday(ctx, biz.User{
		Model: gorm.Model{
			ID: uint(ctx.Value("userid").(float64)),
		},
		Birthday: &t,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{User: &v1.UserInfoReply_User{
		Mobile: user.Mobile,
	},
	}, nil
}

func (u *UserService) UpdateUserGender(ctx context.Context, in *v1.UpdateGenderRequest) (*v1.UserInfoReply, error) {
	user, err := u.uc.UpdateUserGender(ctx, biz.User{
		Model: gorm.Model{
			ID: uint(ctx.Value("userid").(float64)),
		},
		Gender: in.User.Gender,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{User: &v1.UserInfoReply_User{
		Mobile: user.Mobile,
	},
	}, nil
}

// DeleteUser TODO ??????????????????
func (u *UserService) DeleteUser(ctx context.Context, in *v1.DeleteRequest) (*v1.UserInfoReply, error) {
	return nil, nil
}

// SendVerificationCode TODO ??????????????????
func (u *UserService) SendVerificationCode(ctx context.Context, in *v1.SendCodeRequest) (*v1.SendCodeReply, error) {
	return nil, nil
}
