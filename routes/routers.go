package routes

import (
	"github.com/gin-gonic/gin"
	"user/app/http/controllers/api"
	"user/app/middleware"
)

func New(engine *gin.Engine) {
	engine.Use(middleware.Logger)
	engine.Use(middleware.Error)
	auth := engine.Group("auth")
	{
		auth.POST("/login", api.Login)
		auth.POST("/register", api.Register)
	}

	user := engine.Group("user")
	user.Use(middleware.Auth)
	{
		user.GET("/info", api.Info)
		user.POST("/updateInfo", api.UpdateInfo)
		user.POST("/uploadAvatar", api.UploadAvatar)
		user.POST("/changePassword", api.ChangePassword)
	}
}
