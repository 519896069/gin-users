package commands

import (
	"user/app/models"
	"user/fzp"
)

func Migration(args ...string) {
	fzp.Runtime.Db.AutoMigrate(&models.User{})
	fzp.Runtime.Db.AutoMigrate(&models.Token{})
}
