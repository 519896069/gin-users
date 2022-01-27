package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"user/app/business"
	"user/app/http/params"
	"user/fzp/helper"
)

func Info(ctx *gin.Context) {
	time.Sleep(5 * time.Second)
	user, _ := ctx.Get("authUser")
	helper.Ok(ctx, user)
}

// UpdateInfo 改个人信息
func UpdateInfo(ctx *gin.Context) {
	var query params.UserInfo
	err := ctx.Bind(&query)
	if err != nil {
		panic(err)
	}
	helper.Ok(ctx, business.UpdateUserInfo(ctx, query))
}

// UploadAvatar 上传头像
func UploadAvatar(ctx *gin.Context) {
	helper.Ok(ctx, business.UploadAvatar(ctx))
}
