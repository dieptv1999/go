package main

import (
	"fmt"
	"gen-go/qlnv"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	addr := "localhost:9090"
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		fmt.Println(err.Error() + "server")
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	handlerUser := NewUserServiceHandler()
	processorUser := qlnv.NewUserServiceProcessor(handlerUser)
	handlerUnit := NewUnitServiceHandler()
	processorUnit := qlnv.NewUnitServiceProcessor(handlerUnit)
	processor := thrift.TMultiplexedProcessor{}
	processor.RegisterProcessor("UserService", processorUser)
	processor.RegisterProcessor("UnitService", processorUnit)
	server := thrift.NewTSimpleServer4(&processor, transport, transportFactory, protocolFactory)
	errServe := server.Serve()
	if errServe != nil {
		fmt.Println(errServe)
	}
}
