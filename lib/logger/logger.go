package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type Log struct {
	c      *gin.Context
	logger *logrus.Logger
}

func Logger(c *gin.Context) *Log {
	filename := "log-" + time.Now().Format("2006-01-02") + ".log"
	dir, _ := os.Getwd()
	// 日志文件
	filePath := path.Join(dir+"/storage/logs", filename)
	stat, err := os.Stat(filePath)
	var src *os.File
	if stat == nil {
		src, err = os.Create(filePath)
		if err != nil {
			fmt.Println("err", err)
		}
	} else {
		// 写入文件
		src, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("err", err)
		}
	}
	//实例化
	logger := logrus.New()
	logger.Out = src
	return &Log{
		c,
		logger,
	}
}

func (l Log) log(fields logrus.Fields, level string) {
	fields["level"] = level
	l.logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	l.logger.WithFields(fields).Info()
}

func (l Log) Error(fields logrus.Fields) {
	l.log(fields, "error")
}

func (l Log) Info(fields logrus.Fields) {
	l.log(fields, "info")
}
