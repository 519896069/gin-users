package middleware

import (
	"github.com/gin-gonic/gin"
	"user/app/models"
)

func Auth(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		panic("token is required")
	} else {
		if !checkToken(token) {
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
