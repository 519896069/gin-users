package fzp

import (
	"gorm.io/gorm"
	"user/fzp/db"
	"user/fzp/redis"
)

var Runtime = &Application{}

type Application struct {
	Db    *gorm.DB
	Redis *redis.RedisConnect
}

func init() {
	Runtime.Db = db.InitDb()
	Runtime.Redis = redis.InitRedis()
}
