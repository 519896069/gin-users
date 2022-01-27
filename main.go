package main

import (
	"os"
	"user/fzp/bootstrap"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("请输入命令 ./exec mode args...")
	}
	switch args[1] {
	case "http":
		bootstrap.StartRpc()
		bootstrap.StartHttp()
	case "rpc":
		bootstrap.StartRpc()
	default:
		bootstrap.Exec(args[1:])
	}
}
