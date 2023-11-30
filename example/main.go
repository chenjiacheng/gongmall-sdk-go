package main

import (
	"fmt"

	"github.com/chenjiacheng/gongmall-sdk-go/gongmall"
)

func main() {
	client := gongmall.NewClient("Your AppKey", "Your AppSecret", "Your ServiceId")
	// client := gongmall.NewSandboxClient("Your AppKey", "Your AppSecret", "Your ServiceId") // 沙箱环境

	res, err := client.Account.QueryBalance(gongmall.QueryBalanceReq{
		ServiceID: "111",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Success)
	fmt.Println(res.Data)
}
