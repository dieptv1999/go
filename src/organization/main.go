package main

import (
	"github.com/astaxie/beego"
)



//func runUnitClient(typeMethod string) error {
//	addr := "localhost:9090"
//	transportFactory := thrift.NewTBufferedTransportFactory(8192)
//	var transport thrift.TTransport
//	var err error
//	transport, err = thrift.NewTSocket(addr)
//	if err != nil {
//		fmt.Println("Error opening socket:", err)
//		return err
//	}
//	transport, err = transportFactory.GetTransport(transport)
//	if err != nil {
//		return err
//	}
//	defer transport.Close()
//	if err := transport.Open(); err != nil {
//		return err
//	}
//
//	protocol := thrift.NewTBinaryProtocolTransport(transport)
//	iprot := thrift.NewTMultiplexedProtocol(protocol, "UnitService")
//	oprot := thrift.NewTMultiplexedProtocol(protocol, "UnitService")
//	return handleClient(qlnv.NewUserServiceClient(thrift.NewTStandardClient(iprot, oprot)), typeMethod)
//}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
