package main

import (
	"os"
	bootstrap2 "user/fzp/bootstrap"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("请输入命令 ./exec mode args...")
	}
	switch args[1] {
	case "http":
		bootstrap2.StartRpc()
		bootstrap2.StartHttp()
	case "rpc":
		bootstrap2.StartRpc()
	default:
		bootstrap2.Exec(args[1:])
	}
}
