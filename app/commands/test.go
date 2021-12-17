package commands

import (
	"encoding/json"
	"fmt"
	"user/lib"
)

type test struct {
	Ab string `json:"ab"`
	C  string `json:"c"`
	D  []int  `json:"d"`
}

func Test(args ...string) {
	//ok := lib.Redis.Set("test", map[string]interface{}{
	//	"ab": "b",
	//	"c":  "b",
	//	"d":  []int{1, 2, 3, 4, 5},
	//})
	re, ok := lib.Redis.Get("test")
	if ok {
		r := test{}
		json.Unmarshal(re, &r)
		fmt.Printf("%T,%s\n", r, r)
		fmt.Printf("%T,%v\n", r.D[:1], r.D)
		fmt.Printf("%T,%v\n", []int{1, 2, 3}[2], [3]int{1, 2, 3})
	} else {
		fmt.Println("err")
	}
}
