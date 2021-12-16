package kernel

import (
	"user/app/commands"
)

type CommandHandler func(args ...string)

var Commands = map[string]CommandHandler{
	"test" : commands.Test,
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
