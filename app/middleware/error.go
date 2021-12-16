package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user/lib"
)

func Error(ctx *gin.Context) {
	defer catchError(ctx)

	ctx.Next()
}

func catchError(ctx *gin.Context) {
	if err := recover(); err != nil {
		ctx.Abort()
		lib.Error(ctx, gin.H{}, fmt.Sprintf("%v", err))
	}
}
