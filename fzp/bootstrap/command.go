package bootstrap

import (
	"user/app/commands"
)

type CommandHandler func(args ...string)

var Commands = map[string]CommandHandler{
	"test":      commands.Test,
	"migration": commands.Migration,
}

func getCommand(cmd string) CommandHandler {
	doFunc, ok := Commands[cmd]
	if !ok {
		panic("命令不存在")
	}
	return doFunc
}

func Exec(args []string) {
	handler := getCommand(args[0])
	handler(args[1:]...)
}
