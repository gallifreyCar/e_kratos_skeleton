// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"e_kratos_skeleton/internal/biz"
	"e_kratos_skeleton/internal/data"
	"e_kratos_skeleton/internal/handler"
)

// Injectors from wire.go:

// 初始化app 依赖注入
func wireApp() (App, func(), error) {
	dataData, cleanup, err := data.NewData()
	if err != nil {
		return App{}, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	userCase := biz.NewUserCase(userRepo)
	userHandler := handler.NewUserHandler(userCase)
	app := newApp(userHandler)
	return app, func() {
		cleanup()
	}, nil
}
