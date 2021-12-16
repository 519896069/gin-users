package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

func Uuid() string {
	date := strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1) //17位时间
	uniqueKey := "00000000"                                                      //8位机器唯一标识
	return date + uniqueKey + fmt.Sprintf("%0*d", 7, getId())
}

func getId() int64 {
	//获取redis连接
	client := Redis.Pool.Get()
	defer client.Close()
	//判断是否存在不存在则设置初始化1秒过期
	exists, _ := redis.Bool(client.Do("exists", "uid_count"))
	if !exists {
		client.Do("set", "uid_count", "0")
		client.Do("expire", "uid_count", "1")
	}
	//自增
	value, err := redis.Int64(client.Do("INCRBY", "uid_count", "1"))
	if err != nil {
		panic("创建用户ID失败")
	}
	return value
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "ok",
		"data": data,
	})
}

func Error(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": "-1",
		"msg":  msg,
		"data": data,
	})
}

func Mkdir(dir string) bool {
	//判断dir是否存在
	_, existErr := os.Stat(dir)
	if existErr == nil {
		//存在则不创建
		return true
	}
	//创建dir
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		//失败
		return false
	}
	//成功
	return true
}
