package kernel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user/routes"
)

func StartHttp() {
	//获取gin实例
	engine := gin.New()
	engine.Static("/static", "./storage/static")
	//初始化路由
	routes.New(engine)
	err := engine.Run(":8888")
	if err != nil {
		fmt.Println("http error")
		return
	}
}
