package commands

import (
	"encoding/json"
	"fmt"
	"time"
	"user/fzp"
)

type test struct {
	Ab string `json:"ab"`
	C  string `json:"c"`
	D  []int  `json:"d"`
}

func testList(index int) {
	time.Sleep(time.Duration(index) * time.Second)
	t := test{
		Ab: "ab",
		C:  "c",
		D:  []int{},
	}
	t.D = append(t.D, index)
	marshal, _ := json.Marshal(t)
	fzp.Runtime.Redis.Lpush("tlist", string(marshal))
	t = test{}
	fmt.Printf("push:%d\n", index)
}

func Test(args ...string) {
	t := test{
		Ab: "ab",
		C:  "c",
		D:  []int{1, 2, 3},
	}
	marshal, _ := json.Marshal(t)
	fzp.Runtime.Redis.Lpush("tlist", string(marshal))
	t = test{}
	fzp.Runtime.Redis.Rpop("tlist", &t)
	fmt.Println(t)
}
