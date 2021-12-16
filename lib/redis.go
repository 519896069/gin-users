package lib

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	appConfig "user/config"
)

var (
	Redis *RedisConnect
)

type RedisConnect struct {
	Pool *redis.Pool
}

func init() {
	config := appConfig.CONFIG.Redis
	Redis = &RedisConnect{
		&redis.Pool{
			MaxIdle:     config.MaxIdle, //最大空闲连接数
			MaxActive:   config.Active,  //最大连接数
			IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
			Wait:        true, //超过连接数后是否等待
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))
				if err != nil {
					return nil, err
				}
				if config.Password != "" {
					if _, err := c.Do("AUTH", config.Password); err != nil {
						c.Close()
						return nil, err
					}
				}
				if config.Database != 0 {
					if _, err := c.Do("SELECT", config.Database); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
		},
	}
}
