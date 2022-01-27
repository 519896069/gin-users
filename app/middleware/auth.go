package middleware

import (
	"github.com/gin-gonic/gin"
	"user/app/business"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		panic("登录信息失效，请重新登录~")
	} else {
		if !business.CheckToken(ctx, token[7:]) {
			panic("请先登录")
		} else {
			ctx.Next()
		}
	}
}
