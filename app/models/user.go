package models

import (
	"github.com/gin-gonic/gin"
	"time"
	"user/fzp"
	"user/fzp/helper"
)

type User struct {
	Model
	Username      string    `gorm:"type:varchar(32);comment:用户名" json:"username"`
	Password      string    `gorm:"type:varchar(32);comment:密码" json:"-"`
	Salt          string    `gorm:"type:varchar(32);comment:密码盐" json:"-"`
	Avatar        string    `gorm:"type:varchar(256);comment:头像地址;default:/static/img/default.png" json:"avatar"`
	Email         string    `gorm:"type:varchar(128);comment:邮箱地址" json:"email"`
	Mobile        string    `gorm:"type:varchar(11);comment:手机号;default:''" json:"mobile"`
	LastLoginTime time.Time `gorm:"comment:最后登录时间;default:'1970-01-01 00:00:00'" json:"-"`
	Status        int       `gorm:"comment:用户状态;default:0" json:"status"`
}

func GetUserModel() *User {
	return &User{
		Model: Model{
			Db: fzp.Runtime.Db,
		},
	}
}

func (user *User) CheckPassword(password string) bool {
	return EncryptPassword(password, user.Salt) == user.Password
}

func (user *User) GetToken() (*Token, bool) {
	var token Token
	tx := fzp.Runtime.Db.Where("uid=?", user.ID).Where("expired>?", time.Now()).First(&token)
	if tx.Error == nil {
		return nil, false
	}
	if token.ID == 0 {
		return nil, false
	}
	return &token, true
}

func EncryptPassword(password string, salt string) string {
	return helper.Md5(password + salt)
}

func Auth(ctx *gin.Context) *User {
	user, ok := ctx.Get("authUser")
	if !ok {
		panic("请重新登录")
	}
	return user.(*User)
}
