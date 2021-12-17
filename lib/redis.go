package lib

import (
	"encoding/json"
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

func (r RedisConnect) Get(key string) ([]byte, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := conn.Do("get", key)
	if err != nil {
		return []byte{}, false
	}
	return reply.([]byte), true
}

func (r RedisConnect) Set(key string, value interface{}) bool {
	valueType := fmt.Sprintf("%T", value)
	if valueType != "string" && valueType != "[]byte" {
		valueByte, err := json.Marshal(value)
		if err != nil {
			return false
		}
		value = string(valueByte)
	}
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, value)
	return err == nil
}

func (r RedisConnect) Expired(key string, second int) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("expire", key, second)
	return err == nil
}

func (r RedisConnect) Setex(key string, value interface{}, second int) bool {
	ok := r.Set(key, value)
	if !ok {
		return ok
	}
	ok = r.Expired(key, second)
	return ok
}
