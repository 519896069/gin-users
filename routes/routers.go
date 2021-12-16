package routes

import (
	"github.com/gin-gonic/gin"
	api2 "user/app/http/controllers/api"
	"user/app/middleware"
)

func New(engine *gin.Engine) {
	engine.Use(middleware.Logger)
	engine.Use(middleware.Error)
	auth := engine.Group("auth")
	{
		auth.POST("/login", api2.Login)
		auth.POST("/register", api2.Register)
	}

	user := engine.Group("user")
	user.Use(middleware.Auth)
	{
		user.GET("/info", api2.Info)
		user.POST("/updateInfo", api2.UpdateInfo)
		user.POST("/uploadAvatar", api2.UploadAvatar)
		user.POST("/changePassword", api2.ChangePassword)
	}
}
