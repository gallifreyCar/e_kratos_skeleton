package v1

import (
	"context"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	// Name 名称
	Name string `json:"name"`
	// Age 年龄
	Age int `json:"age"`
}

type FindUserByIDRequest struct {
	// ID 用户ID
	ID int `json:"id"`
}

// UserHandler Handler层的UserHandler接口定义
type UserHandler interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (CommonResponse, error)
	FindUserByID(ctx context.Context, request FindUserByIDRequest) (CommonResponse, error)
}

func userFindUserById0Handler(h UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req FindUserByIDRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := h.FindUserByID(c, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func createUserHandler(h UserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := h.CreateUser(c, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}

}

type userController struct {
	h UserHandler
}

func (ctrl userController) createUserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := ctrl.h.CreateUser(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)

}

func RegisterUserHandler(r *gin.Engine, h UserHandler) {
	r.POST("/create", createUserHandler(h))
	r.GET("/find", userFindUserById0Handler(h))

}
