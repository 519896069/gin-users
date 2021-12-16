package middleware

import (
	"github.com/gin-gonic/gin"
	"user/lib"
)

func Verify(ctx *gin.Context) {
	ctx.Next()
}

func checkApiSign(params map[string]string) bool {
	sign, ok := params["sign"]
	if !ok {
		return false
	}
	delete(params, "sign")
	return getApiSign(params) == sign
}

func getApiSign(params map[string]string) string {
	signStr := ""
	for key, value := range params {
		signStr += key + "=" + value
	}
	return lib.Md5(signStr)
}
