package biz

import (
	"google.golang.org/protobuf/proto"
	v1 "shop/api/user/v1"
)

func Marshal(info *v1.UserInfoReply) ([]byte, error) {
	marshal, err := proto.Marshal(info)
	if err != nil {
		return nil, err
	}
	return marshal, err
}

func Unmarshal(info []byte) (User, error) {
	var v v1.UserInfoReply
	err := proto.Unmarshal(info, &v)
	if err != nil {
		return User{}, err
	}
	return User{
		Mobile:   v.User.GetMobile(),
		NickName: v.User.GetNikeName(),
		Birthday: v.User.GetBirthday(),
		Gender:   v.User.GetGender(),
	}, nil
}
