package bootstrap

import "os"

func Bootstrap() {
	args := os.Args
	if len(args) == 1 {
		panic("请输入命令 ./exec mode args...")
	}
	switch args[1] {
	case "http":
		StartRpc()
		StartHttp()
	case "rpc":
		StartRpc()
	default:
		Exec(args[1:])
	}
}
