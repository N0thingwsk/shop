package biz

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
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
	UpdateUser(context.Context, User) (User, error)
	DeleteUser(context.Context, int) (User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

//func Test(pass1, pass2 string, options *password.Options) bool {
//	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
//	salt, encodedPwd := password.Encode(use.Password, options)
//	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
//	passwordInfo := strings.Split(pass1, "$")
//	return password.Verify(pass2, passwordInfo[2], passwordInfo[3], options)
//}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user User) (User, error) {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(user.Password, options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	us, err := u.repo.CreateUser(ctx, User{
		Mobile:   user.Mobile,
		Password: newPassword,
		NickName: user.NickName,
	})
	if err != nil {
		return User{}, err
	}
	return User{
		Mobile:   us.Mobile,
		NickName: us.NickName,
	}, nil
}
