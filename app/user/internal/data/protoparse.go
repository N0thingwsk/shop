package data

import (
	"context"
	"google.golang.org/protobuf/proto"
	v1 "shop/api/user/v1"
	"shop/app/user/internal/biz"
)

// ProtoMarshal proto序列化
func (r *userRepo) ProtoMarshal[T any](ctx context.Context, user T) ([]byte, error) {
	//q: go泛型语法
	//a:
	info := &v1.UserInfoReply{
		User: &v1.UserInfoReply_User{
			Mobile:   user.Mobile,
			NikeName: user.NickName,
			Birthday: user.Birthday,
			Gender:   user.Gender,
		},
	}
	marshal, err := proto.Marshal(user)
	if err != nil {
		return nil, err
	}
	//TODO
	return marshal, nil
}

// ProtoUnmarshal proto反序列化
func (r *userRepo) ProtoUnmarshal(ctx context.Context, user []byte) (biz.User, error) {
	return biz.User{}, nil
}
