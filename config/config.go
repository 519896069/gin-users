package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

var (
	CONFIG = new(Config)
)

func init() {
	pwd, _ := os.Getwd()
	err := ini.MapTo(CONFIG, pwd+"/config/app.ini")
	if err != nil {
		log.Println(err)
		return
	}
}

type Config struct {
	App      App      `ini:"app"`
	Database Database `ini:"database"`
	Redis    Redis    `ini:"redis"`
}

type App struct {
	AppName string `ini:"app_name"`
	AppKey  string `ini:"app_key"`
}

type Database struct {
	Host     string `ini:"db_host"`
	Port     string `ini:"db_port"`
	Database string `ini:"db_database"`
	Username string `ini:"db_username"`
	Password string `ini:"db_password"`
}

type Redis struct {
	Host        string `ini:"redis_host"`
	Port        string `ini:"redis_port"`
	MaxIdle     int    `ini:"redis_maxidle"`
	Active      int    `ini:"redis_active"`
	IdleTimeout int    `ini:"redis_timeout"`
	Password    string `ini:"redis_password"`
	Database    int    `ini:"redis_database"`
}
