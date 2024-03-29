package data

import (
	"context"
	"e_kratos_skeleton/internal/biz"
)

type userRepo struct {
	data *Data
}

// NewUserRepo  UserRepo的构造函数
func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

// Create 创建用户
func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	r.data.gormDB.Create(u)
	return u, nil
}

// FindByID 根据ID查找用户
func (r *userRepo) FindByID(ctx context.Context, id int) (*biz.User, error) {
	var user biz.User
	r.data.gormDB.First(&user, id)
	return &user, nil
}
