package main

import (
	"fmt"
	"github.com/Hotchilli-Cloud/go-sdk/cmd/hcloud"
)

func main() {
	c := hcloud.NewClient("a6b29b7bbcea43928726bf1cdc5287cc")

	//res, err := c.CreateZone("example.com")
	//res, err := c.DeleteZone("example.com")

	//res, err := c.CreateRecord("example.com", &hcloud.CreateRecordRequest{
	//	Priority:    "0",
	//	RecordName:  "example.com",
	//	RecordType:  "A",
	//	RecordValue: "46.17.220.150",
	//	TTL:         "300",
	//})

	//res, err := c.DeleteRecord("example.com", "136")
	res, err := c.GetRecord("example.com", "137")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(res.RecordName)

	//res, err := c.CreateCheckoutSessionItem("cs_yx6ddpcKqNDuYbhFyYkAv3yqKWwnxFJrx4P85Xp1OLQVktZQ", hcloud.ItemType{
	//	ID:            "",
	//	Object:        "",
	//	CreatedTime:   0,
	//	SessionId:     "",
	//	Type:          "",
	//	AutoTopup:     false,
	//	AutoRenew:     false,
	//	BillingPeriod: "",
	//	CartName:      "",
	//	Period:        0,
	//	Item:          nil,
	//	Subtotal:      0,
	//})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println(res)
}
