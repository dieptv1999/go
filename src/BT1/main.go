package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

func main() {
	isServer := false
	addr := "localhost:9000"
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	if isServer {
		if err := runServer(transportFactory, protocolFactory, addr); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		time.Sleep(1 * time.Second)
		if err := runClient(transportFactory, protocolFactory, addr); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
