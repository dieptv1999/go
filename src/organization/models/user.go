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

func getUserClient() *qlnv.UserServiceClient {
	addr := "localhost:9090"
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return nil
	}

	protocol := thrift.NewTBinaryProtocolTransport(transport)
	iprot := thrift.NewTMultiplexedProtocol(protocol, "UserService")
	oprot := thrift.NewTMultiplexedProtocol(protocol, "UserService")
	return qlnv.NewUserServiceClient(thrift.NewTStandardClient(iprot, oprot))
}

func AddUser(u qlnv.User) string {
	u.UserId = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	client := getUserClient()
	err := client.CreateUser(defaultCtx, &u)
	if err != nil {
		fmt.Println(err.Error())
	}
	return u.UserId
}

func GetUser(uid string) (u *qlnv.User, err error) {
	client := getUserClient()
	if u, ok := client.ReadUser(defaultCtx, uid); ok != nil {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func UpdateUser(uu *qlnv.User) error {
	client := getUserClient()
	e := client.UpdateUser(defaultCtx, uu)
	if e != nil {
		return errors.New("Not found exception!")
	}
	//if u, ok := UserList[uid]; ok {
	//	if uu.Name != "" {
	//		u.Name = uu.Name
	//	}
	//	if uu.DateOfBirth != "" {
	//		u.DateOfBirth = uu.DateOfBirth
	//	}
	//	if uu.Address != "" {
	//		u.Address = uu.Address
	//	}
	//	return u, nil
	//}
	return nil
}

func DeleteUser(uid string) error {
	client := getUserClient()
	err := client.DeleteUser(defaultCtx, uid)
	if err != nil {
		return errors.New("User not exist")
	}
	return nil
}

func GetUnitOfUser(uid string) (unit *qlnv.Unit, err error) {
	client := getUserClient()
	if unit, err := client.GetUnitUser(defaultCtx, uid); err != nil {
		return unit, nil
	}
	return nil, errors.New("Unit not exist")
}

func GetUserSortedByPage(userId string, numOfPages int32, sizeOfpage int32) (r []*qlnv.User, err error) {
	client := getUserClient()
	if r, err := client.GetUserSortedByPage(defaultCtx, userId, numOfPages, sizeOfpage); err != nil {
		return r, nil
	}
	return nil, errors.New("User not exist")
}
