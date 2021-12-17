package models

import (
	"time"
	"user/lib"
)

var (
	AuthUser *User
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
			Db: lib.Mysql.Db,
		},
	}
}

func (user *User) CheckPassword(password string) bool {
	return EncryptPassword(password, user.Salt) == user.Password
}

func (user *User) GetToken() (*Token, bool) {
	var token Token
	tx := lib.Mysql.Db.Where("uid=?", user.ID).Where("expired>?", time.Now()).First(&token)
	if tx.Error == nil {
		return nil, false
	}
	return &token, true
}

func EncryptPassword(password string, salt string) string {
	return lib.Md5(password + salt)
}
