package business

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
	"user/app/http/params"
	"user/app/models"
	"user/fzp"
	"user/fzp/helper"
)

func UpdateUserInfo(params params.UserInfo) *models.User {
	models.AuthUser.Username = params.Username
	models.AuthUser.Email = params.Email
	models.AuthUser.Mobile = params.Mobile
	fzp.Runtime.Db.Updates(models.AuthUser)
	return models.AuthUser
}

func UploadAvatar(ctx *gin.Context) string {
	var (
		rootPath, _ = os.Getwd()
		date        = time.Now().Format("20060102")
		userDir     = strconv.Itoa(int(models.AuthUser.ID % 100))
		filename    = helper.Md5(strconv.Itoa(int(models.AuthUser.ID))) + ".png"
	)
	var (
		root         = rootPath + "/storage"
		uploadPath   = fmt.Sprintf("/static/img/avatar/%s/%s/", date, userDir)
		completePath = uploadPath + filename
	)
	avatar, getErr := ctx.FormFile("avatar")
	if getErr != nil {
		panic("头像上传失败")
	}
	//判断目录是否存在
	helper.Mkdir(root + uploadPath)
	//地址：/日期Ymd/用户id取余100/md5(uid).png
	uploadErr := ctx.SaveUploadedFile(avatar, root+completePath)
	if uploadErr != nil {
		panic(fmt.Sprintf("%v", uploadErr))
	}
	models.AuthUser.Avatar = completePath
	fzp.Runtime.Db.Updates(models.AuthUser)
	return completePath
}
