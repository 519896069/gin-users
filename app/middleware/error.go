package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user/fzp/helper"
)

func Error(ctx *gin.Context) {
	defer catchError(ctx)

	ctx.Next()
}

func catchError(ctx *gin.Context) {
	if err := recover(); err != nil {
		ctx.Abort()
		helper.Error(ctx, gin.H{}, fmt.Sprintf("%v", err))
	}
}
