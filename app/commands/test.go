package commands

import "fmt"

func Test(args ...string)  {
	fmt.Println("test")
	fmt.Println(args)
}