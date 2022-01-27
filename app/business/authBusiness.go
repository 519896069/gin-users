package business

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user/app/http/params"
	"user/app/models"
	"user/fzp"
	"user/fzp/helper"
)

func Login(ctx *gin.Context, params params.Login) gin.H {
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

func Register(ctx *gin.Context, params params.Register) gin.H {
	user := models.GetUserModel()
	query := user.Db.Where("email=?", params.Email).First(&user)
	if query.Error == nil {
		panic("邮箱已存在")
	}
	//验证
	//验证成功
	salt := helper.Md5(helper.Uuid())
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

func ChangePassword(ctx *gin.Context, params params.ChangePassword) {
	salt := helper.Uuid()
	user := models.Auth(ctx)
	user.Salt = salt
	user.Password = models.EncryptPassword(params.Password, salt)
	tx := fzp.Runtime.Db.Updates(user)
	if tx.Error != nil {
		panic(tx.Error.Error())
	}
}

func CheckToken(ctx *gin.Context, token string) bool {
	user := models.GetTokenModel().GetUserByToken(token)
	if user != nil {
		ctx.Set("authUser", user)
		return true
	}
	return false
}
