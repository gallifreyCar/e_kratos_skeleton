package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data 里面包含数据的客户端
type Data struct {
	// 数据库客户端
	gormDB *gorm.DB
	// redis客户端
	redis *redis.Client
}

// NewData . Data的构造函数
func NewData() (*Data, func(), error) {

	cleanup := func() {
		println("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
