package handler

import (
	"context"
	v1 "e_kratos_skeleton/api/http/v1"
	"e_kratos_skeleton/internal/biz"
)

// UserHandler UserHandler
type UserHandler struct {
	uc biz.UserCase
}

type UserCase interface {
	CreateUser(ctx context.Context, user *biz.User) (*biz.User, error)
	GetUser(ctx context.Context, id int) (*biz.User, error)
}

// NewUserHandler UserHandler的构造函数
func NewUserHandler(uc *biz.UserCase) *UserHandler {
	return &UserHandler{
		*uc,
	}
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(ctx context.Context, request v1.CreateUserRequest) (v1.CommonResponse, error) {
	user, err := h.uc.CreateUser(ctx, &biz.User{
		Name: request.Name,
		Age:  request.Age,
	})

	if err != nil {
		return v1.CommonResponse{}, err
	}

	return v1.CommonResponse{
		Code:    0,
		Message: "success",
		Data:    user,
	}, nil

}

// FindUserByID 根据ID查找用户
func (h *UserHandler) FindUserByID(ctx context.Context, request v1.FindUserByIDRequest) (v1.CommonResponse, error) {
	user, err := h.uc.GetUser(ctx, request.ID)
	if err != nil {
		return v1.CommonResponse{}, err
	}
	return v1.CommonResponse{
		Code:    0,
		Message: "success",
		Data:    user,
	}, nil
}
