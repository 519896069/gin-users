package api

import (
	"github.com/gin-gonic/gin"
	"user/app/business"
	"user/app/http/params"
	"user/fzp/helper"
)

func Login(ctx *gin.Context) {
	var query params.Login
	err := ctx.ShouldBind(&query)
	if err != nil {
		panic(err)
	}
	helper.Ok(ctx, business.Login(ctx, query))
}

func Register(ctx *gin.Context) {
	var query params.Register
	err := ctx.ShouldBind(&query)
	if err != nil {
		panic(err)
	}
	helper.Ok(ctx, business.Register(ctx, query))
}

// ChangePassword 修改密码
func ChangePassword(ctx *gin.Context) {
	var query params.ChangePassword
	err := ctx.Bind(&query)
	if err != nil {
		panic(err)
	}
	business.ChangePassword(ctx, query)
	helper.Ok(ctx, nil)
}
