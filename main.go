package main

import (
	"fmt"
	"github.com/Hotchilli-Cloud/go-sdk/cmd/hcloud"
)

func main() {
	c := hcloud.NewClient("hcloud_test_nwkrhiu2hwfsafwq")

	res, err := c.CreateCheckoutSession()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(res)
	fmt.Println(res.ID)
}
