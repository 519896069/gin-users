package commands

import (
	"user/app/models"
	"user/lib"
)

func Migration(args ...string) {
	lib.Mysql.Db.AutoMigrate(&models.User{})
	lib.Mysql.Db.AutoMigrate(&models.Token{})
}
