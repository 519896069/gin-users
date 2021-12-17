package middleware

import (
	"github.com/gin-gonic/gin"
	"user/app/models"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		panic("登录信息失效，请重新登录~")
	} else {
		if !checkToken(token[7:]) {
			panic("请先登录")
		} else {
			ctx.Next()
		}
	}
}

func checkToken(token string) bool {
	user := models.GetTokenModel().GetUserByToken(token)
	if user != nil {
		models.AuthUser = user
		return true
	}
	return false
}
