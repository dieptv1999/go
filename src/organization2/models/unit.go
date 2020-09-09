package models

import (
	"errors"
	"fmt"
	"gen-go/qlnv"
	"github.com/apache/thrift/lib/go/thrift"
	"strconv"
	"time"
)

func GetUnitClient(transport thrift.TTransport) *qlnv.UnitServiceClient {

	protocol := thrift.NewTBinaryProtocolTransport(transport)
	iprot := thrift.NewTMultiplexedProtocol(protocol, "UnitService")
	oprot := thrift.NewTMultiplexedProtocol(protocol, "UnitService")
	return qlnv.NewUnitServiceClient(thrift.NewTStandardClient(iprot, oprot))
}

func AddUnit(u qlnv.Unit) string {
	u.UnitId = "unit_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUnitClient(transport)
	err := client.CreateUnit(defaultCtx, &u)
	if err != nil {
		return err.Error()
	}
	return u.UnitId
}

func GetUnit(uid string) (u *qlnv.Unit, err error) {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUnitClient(transport)
	u, ok := client.ReadUnit(defaultCtx, uid)
	if ok == nil {
		fmt.Println("ok")
		return u, nil
	}
	return nil, ok
}

func UpdateUnit(uu *qlnv.Unit) error {
	fmt.Println(uu.UnitId)
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUnitClient(transport)
	e := client.UpdateUnit(defaultCtx, uu)
	if e != nil {
		return e
	}
	return nil
}

func DeleteUnit(uid string) error {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUnitClient(transport)
	err := client.DeleteUnit(defaultCtx, uid)
	if err != nil {
		return err
	}
	return nil
}

func GetAllMemberOfUnit(uid string) (unit []*qlnv.User, err error) {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUnitClient(transport)
	if unit, err := client.GetAllMemberOfUnit(defaultCtx, uid); err != nil {
		return unit, nil
	}
	return nil, errors.New("Unit not exist")
}

func GetMembersByPage(UnitId string, numOfPages int32, sizeOfpage int32) (r []*qlnv.User, err error) {
	var transport = getTransport()
	if transport != nil {
		defer transport.Close()
	}
	client := GetUnitClient(transport)
	if r, err := client.GetMembersByPage(defaultCtx, UnitId, numOfPages, sizeOfpage); err != nil {
		return r, nil
	}
	return nil, errors.New("Unit not exist")
}
