package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

type User struct {
	gorm.Model
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6)"`
	Role     int        `gorm:"column:role;default:1;type:int"`
}

type UserRepo interface {
	GetUserList(context.Context, int) ([]User, error)
	GetUserinfoByMobile(context.Context, string) (User, error)
	GetUserinfoById(context.Context, int) (User, error)
	CreateUser(context.Context, User) (User, error)
	UpdateUser(context.Context, int) (User, error)
	DeleteUser(context.Context, int) (User, error)
}

type UserUsecase struct {
	repo UserRepo
	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

