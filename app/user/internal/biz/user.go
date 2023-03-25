package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"shop/app/user/internal/conf"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

// User 用户
type User struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Birthday string `gorm:"type:varchar(200)"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(6)"`
	Role     int    `gorm:"column:role;default:1;type:int"`
}

// UserRepo 用户仓库
type UserRepo interface {
	GetUserInfoById(context.Context, int) (User, error)
	GetUserInfoByMobile(context.Context, string) (User, error)
	GetUserCache(context.Context, string) ([]byte, error)
	SetUserCache(context.Context, string, string) error
	ProtoMarshal(context.Context, User) ([]byte, error)
	ProtoUnmarshal(context.Context, []byte) (User, error)
	CreateUser(context.Context, User) error
	UpdateUser(context.Context, User) error
	DeleteUser(context.Context, int) error
}

// UserUsecase 用户用例
type UserUsecase struct {
	repo UserRepo
	c    *conf.Server
	log  *log.Helper
}

// NewUserUsecase new a user usecase.
func NewUserUsecase(repo UserRepo, c *conf.Server, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, c: c, log: log.NewHelper(logger)}
}

// GetUserInfoById 根据id获取用户信息
func (uc *UserUsecase) GetUserInfoById(ctx context.Context, id int) (User, error) {
	return uc.repo.GetUserInfoById(ctx, id)
}
