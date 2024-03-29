package main

import (
	"context"
	v1 "e_kratos_skeleton/api/http/v1"
	"e_kratos_skeleton/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
)

type App struct {
	userHandler *handler.UserHandler
}

func newApp(userHandler *handler.UserHandler) App {
	return App{
		userHandler: userHandler,
	}
}

var (
	cfgFile string // 配置文件
	port    int    // 端口

	rootCmd = &cobra.Command{
		Use:   "ekratos",
		Short: "welcome to e kratos",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// 初始化gin
			r := &gin.Engine{}

			// 初始化app，调用的是 wire_gen.go 中的 initApp
			app, cleanup, err := wireApp()
			v1.RegisterUserHandler(r, app.userHandler)

			defer cleanup()
			if err != nil {
				log.Fatalf("init app failed, err:%v", err)
			}

			// 启动http服务
			r.Run(":8080")

			<-ctx.Done()
			log.Println("Server exiting")
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
