package biz

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strings"
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
	GetUserinfo(context.Context, int) (User, error)
	LoginUser(context.Context, User) (User, error)
	CreateUser(context.Context, User) (User, error)
	UpdateUserNickName(context.Context, User) (User, error)
	UpdateUserPassword(context.Context, User) (User, error)
	UpdateUserBirthday(context.Context, User) (User, error)
	UpdateUserGender(context.Context, User) (User, error)
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

func (u *UserUsecase) GetUserinfo(ctx context.Context, id int) (User, error) {
	user, err := u.repo.GetUserinfo(ctx, id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserUsecase) LoginUser(ctx context.Context, user User) (User, error) {
	use, err := u.repo.LoginUser(ctx, user)
	if err != nil {
		return User{}, err
	}
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(use.Password, "$")
	isPass := password.Verify(user.Password, passwordInfo[2], passwordInfo[3], options)
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

func (u *UserUsecase) UpdateUserNickName(context.Context, User) (User, error) {

}

func (u *UserUsecase) UpdateUserPassword(context.Context, User) (User, error) {

}

func (u *UserUsecase) UpdateUserBirthday(context.Context, User) (User, error) {

}

func (u *UserUsecase) UpdateUserGender(context.Context, User) (User, error) {

}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) (User, error) {
	us, err := u.DeleteUser(ctx, id)
	if err != nil {
		return User{}, err
	}
	return us, nil
}
