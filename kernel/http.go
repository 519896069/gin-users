package kernel

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//优雅重启
	quit := make(chan os.Signal)
	//获取信号
	signal.Notify(quit, os.Interrupt)
	<-quit
	//新建5秒超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
