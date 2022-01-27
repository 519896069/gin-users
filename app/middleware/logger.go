package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"user/lib/logger"
)

type WriterInterceptor struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w WriterInterceptor) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w WriterInterceptor) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	interceptor := &WriterInterceptor{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = interceptor
	c.Next()
	if c.Request.Method == http.MethodOptions {
		return
	}
	// 结束时间
	endTime := time.Now()
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()
	// 执行时间
	latencyTime := endTime.Sub(startTime).String()
	// 返回日志
	logger.Logger(c).Info(logrus.Fields{
		"uri":         reqUri,
		"statusCode":  statusCode,
		"latencyTime": latencyTime,
		"method":      reqMethod,
		"req":         c.Request.PostForm,
		"resp":        interceptor.body.String(),
	})
}
