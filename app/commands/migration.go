package commands

import (
	"user/app/models"
	"user/lib/mysql"
)

func Migration(args ...string) {
	mysql.Mysql.Db.AutoMigrate(&models.User{})
	mysql.Mysql.Db.AutoMigrate(&models.Token{})
}
