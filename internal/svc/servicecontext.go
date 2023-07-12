package svc

import (
	"easyStart/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	DbStore    *gorm.DB
	RedisStore *redis.Client
}

func NewServiceContext(c config.Config, dbStore *gorm.DB, redisStore *redis.Client) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		DbStore:    dbStore,
		RedisStore: redisStore,
	}
}
