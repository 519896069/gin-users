package redis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	appConfig "user/config"
)

type RedisConnect struct {
	Pool *redis.Pool
}

func InitRedis() *RedisConnect {
	config := appConfig.CONFIG.Setting.Redis
	return &RedisConnect{
		Pool: &redis.Pool{
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

func (r RedisConnect) format(reply string, receiver interface{}) (interface{}, bool) {
	if receiver == nil {
		return reply, true
	}
	err := json.Unmarshal([]byte(reply), receiver)
	if err != nil {
		return nil, false
	}
	return reply, true
}

//string
func (r RedisConnect) Get(key string, receiver interface{}) (interface{}, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		return nil, false
	}
	return r.format(reply, receiver)
}

func (r RedisConnect) Set(key string, value string, second int) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	if second == 0 {
		second = -1
	}
	_, err := conn.Do("set", key, value, second)
	return err == nil
}

func (r RedisConnect) Setnx(key string, value string) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := redis.Int(conn.Do("setnx", key, value))
	return err == nil && reply == 1
}

func (r RedisConnect) Expired(key string, second int) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("expire", key, second)
	return err == nil
}

//list
func (r RedisConnect) Lpush(key string, value string) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("lpush", key, value)
	return err == nil
}

func (r RedisConnect) Rpush(key string, value string) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("rpush", key, value)
	return err == nil
}
func (r RedisConnect) Lpop(key string, receiver interface{}) (interface{}, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("lpop", key))
	if err != nil {
		return nil, false
	}
	return r.format(reply, &receiver)
}

func (r RedisConnect) Rpop(key string, receiver interface{}) (interface{}, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("rpop", key))
	if err != nil {
		return nil, false
	}
	return r.format(reply, &receiver)
}

func (r RedisConnect) BRpop(key string, receiver interface{}) (interface{}, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := redis.Values(conn.Do("brpop", key, 30))
	if err != nil {
		return nil, false
	}
	return r.format(string(reply[1].([]byte)), &receiver)
}

func (r RedisConnect) Llen(key string) (int, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	listLen, err := redis.Int(conn.Do("llen", key))
	if err != nil {
		return 0, false
	}
	return listLen, true
}

//hash
func (r RedisConnect) Hget(key string, field string, receiver *interface{}) (interface{}, bool) {
	conn := r.Pool.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("hget", key, field))
	if err != nil {
		return nil, false
	}
	return r.format(reply, receiver)
}

func (r RedisConnect) Hset(key string, field string, value string) bool {
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("hset", key, field, value)
	return err == nil
}

func (r RedisConnect) Lock(key string, second int) bool {
	ok := r.Setnx(key, "1")
	if !ok {
		return false
	}
	r.Expired(key, second)
	return true
}

func (r RedisConnect) UnLock(key string) bool {
	ok := r.Expired(key, 0)
	return ok
}
