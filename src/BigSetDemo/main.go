package main

import (
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"log"
)

func main() {
	client := StringBigsetService.NewStringBigsetServiceModel("BigSetDemo",
		[]string{"0.0.0.0:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "127.0.0.1",
			Port:      "18407",
			ServiceID: "BigSetDemo",
		})
	bskey := generic.TStringKey("demo")
	client.BsPutItem(bskey, &generic.TItem{
		Key:   []byte("tvt"),
		Value: []byte("1234"),
	})
	item, err := client.BsGetItem(bskey, generic.TItemKey("tvd"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(item.String())
	lsItems, err := client.BsGetSliceR(bskey, 0, 10)
	if err != nil {
		log.Println("[ERROR] err", err)
	}
	//chanerr := make(chan error)
	//<-chanerr
	if lsItems != nil {
		for i := 0; i < len(lsItems); i++ {
			log.Println(i, string(lsItems[i].Value), "key", string(lsItems[i].Key))
		}
	}
}
