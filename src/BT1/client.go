package main

import (
	"context"
	"fmt"
	"gen-go/add"
	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client *add.CalculateClient) (err error) {
	var a int32 = 10
	var b int32 = 20
	sum, _ := client.Add(defaultCtx, a, b)
	fmt.Printf("%d+%d=%d\n", a, b, sum)
	return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(add.NewCalculateClient(thrift.NewTStandardClient(iprot, oprot)))
}
