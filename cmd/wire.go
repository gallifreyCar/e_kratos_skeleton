//go:build wireinject
// +build wireinject

package main

import (
	"e_kratos_skeleton/internal/biz"
	"e_kratos_skeleton/internal/data"
	"e_kratos_skeleton/internal/handler"
	"github.com/google/wire"
)

// 初始化app 依赖注入
func wireApp() (App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		handler.ProviderSet,
		newApp,
	))
}
