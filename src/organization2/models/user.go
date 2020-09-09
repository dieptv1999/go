package models

import (
	"context"
	"errors"
	"fmt"
	"gen-go/qlnv"
	"github.com/apache/thrift/lib/go/thrift"
	"strconv"
	"time"
)

var defaultCtx = context.Background()

func getTransport() thrift.TTransport {
	var transport thrift.TTransport
	addr := "localhost:9090"
	//transportFactory := thrift.NewTBufferedTransportFactory(8192)
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil
	}
	//transport, err = transportFactory.GetTransport(transport)
	//if err != nil {
	//	return nil
	//}
	//defer transport.Close()
	if err := transport.Open(); err != nil {
		return nil
	}
	return transport
}
func GetUserClient(transport thrift.TTransport) *qlnv.UserServiceClient {

	protocol := thrift.NewTBinaryProtocolTransport(transport)
	iprot := thrift.NewTMultiplexedProtocol(protocol, "UserService")
	oprot := thrift.NewTMultiplexedProtocol(protocol, "UserService")
	return qlnv.NewUserServiceClient(thrift.NewTStandardClient(iprot, oprot))
}

func AddUser(u qlnv.User) string {
	u.UserId = string(u.DateOfBirth) + strconv.FormatInt(time.Now().UnixNano(), 10)
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUserClient(transport)
	err := client.CreateUser(defaultCtx, &u)
	if err != nil {
		fmt.Println(err.Error())
	}
	return u.UserId
}

func GetUser(uid string) (u *qlnv.User, err error) {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUserClient(transport)
	if u, ok := client.ReadUser(defaultCtx, uid); ok == nil {
		fmt.Println("ok")
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func UpdateUser(uu *qlnv.User) error {
	fmt.Println(uu.UserId)
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUserClient(transport)
	e := client.UpdateUser(defaultCtx, uu)
	if e != nil {
		return e
	}
	return nil
}

func DeleteUser(uid string) error {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUserClient(transport)
	err := client.DeleteUser(defaultCtx, uid)
	if err != nil {
		return err
	}
	return nil
}

func GetUnitOfUser(uid string) (unit *qlnv.Unit, err error) {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUserClient(transport)
	unit, e := client.GetUnitUser(defaultCtx, uid)
	if e == nil {
		return unit, nil
	}
	return nil, e
}

func GetUserSortedByPage(numOfPages int32, sizeOfpage int32, sortType string) (r []*qlnv.User, err error) {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUserClient(transport)
	r, e := client.GetUserSortedByPage(defaultCtx, numOfPages, sizeOfpage, sortType)
	if e == nil {
		return r, nil
	}
	return nil, e
}
