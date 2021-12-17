package models

import (
	"encoding/json"
	"fmt"
	"time"
	"user/lib"
)

type Token struct {
	Model
	Uid     uint      `gorm:"comment:用户id" json:"uid"`
	Token   string    `gorm:"type:varchar(32);comment:token值" json:"token"`
	Expired time.Time `gorm:"default:'1970-01-01 00:00:00'" json:"expired"`
}

//func init() {
//lib.Mysql.Db.AutoMigrate(&Token{})
//}

func GetTokenModel() *Token {
	return &Token{
		Model: Model{
			Db: lib.Mysql.Db,
		},
	}
}

func (t Token) GetUserByToken(tokenStr string) *User {
	var (
		uid   uint64
		cache = true
	)
	//先从缓存中获取UID
	tokenJsonStr, ok := lib.Redis.Get(tokenStr)
	if !ok {
		cache = false
	}
	//解析缓存中的数据
	err := json.Unmarshal(tokenJsonStr, &t)
	if err != nil {
		cache = false
	}
	if t.Expired.Unix() < time.Now().Unix() {
		cache = false
	}
	//判断缓存是否成功
	if cache {
		uid = uint64(t.Uid)
	} else {
		//缓存失效则从数据库中查找uid
		var token Token
		tx := t.Db.Where("token=?", tokenStr).Where("expired>?", time.Now()).First(&token)
		if tx.Error != nil {
			return nil
		}
		uid = uint64(token.Uid)
	}
	//获取User返回
	var user User
	utx := t.Db.Where("id=?", uid).First(&user)
	if utx.Error != nil {
		return nil
	}
	return &user
}

func (t Token) LoginSuccess(user *User) string {
	var (
		token   = lib.Md5(lib.Uuid())
		expired = time.Now().Add(3600 * 30 * 6 * time.Second)
	)
	//设置用户最后登录时间
	user.LastLoginTime = time.Now()
	t.Db.Where("id=?", user.ID).Updates(user)
	//设置用户最近一次token失效
	if existToken, ok := user.GetToken(); ok {
		existToken.setExpired()
	}
	//创建token
	t.Uid = user.ID
	t.Token = token
	t.Expired = expired
	t.Db.Create(&t)

	//把token加入缓存
	json, err := json.Marshal(t)
	if err == nil {
		lib.Redis.Set(token, json)
	} else {
		panic(fmt.Sprintf("%.v", err))
	}
	//返回token
	return token
}

func (t *Token) setExpired() {
	conn := lib.Redis.Pool.Get()
	defer conn.Close()
	conn.Do("HDEL", "user_token", t.Token)
	t.Db.Unscoped().Delete(&t)
}
