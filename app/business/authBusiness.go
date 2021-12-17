package business

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user/app/http/params"
	"user/app/models"
	"user/lib"
)

func Login(params params.Login) gin.H {
	user := models.GetUserModel()
	query := user.Db.Table("users").Where("email=?", params.Email).First(&user)
	if query.Error != nil {
		fmt.Println(query.Error.Error())
		panic("邮箱或密码错误")
	}

	if !user.CheckPassword(params.Password) {
		panic("邮箱或密码错误")
	}
	//登录成功
	return gin.H{
		"token": models.GetTokenModel().LoginSuccess(user),
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"mobile":   user.Mobile,
		},
	}
}

func Register(params params.Register) gin.H {
	user := models.GetUserModel()
	query := user.Db.Where("email=?", params.Email).First(&user)
	if query.Error == nil {
		panic("邮箱已存在")
	}
	//验证
	//验证成功
	salt := lib.Md5(lib.Uuid())
	user.Username = params.Username
	user.Email = params.Email
	user.Password = models.EncryptPassword(params.Password, salt)
	user.Salt = salt
	tx := user.Db.Create(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		panic("用户创建失败，联系肥肥~")
	}

	return gin.H{
		"token": models.GetTokenModel().LoginSuccess(user),
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"mobile":   user.Mobile,
		},
	}
}

func ChangePassword(params params.ChangePassword) {
	salt := lib.Uuid()
	models.AuthUser.Salt = salt
	models.AuthUser.Password = models.EncryptPassword(params.Password, salt)
	tx := lib.Mysql.Db.Updates(&models.AuthUser)
	if tx.Error != nil {
		panic(tx.Error.Error())
	}
}
