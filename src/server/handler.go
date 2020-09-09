package main

import (
	"context"
	"fmt"
	"gen-go/qlnv"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var db *StringBigsetService.StringBigsetServiceIf

func GetDB() *StringBigsetService.StringBigsetServiceIf {
	if db == nil {
		*db = StringBigsetService.NewStringBigsetServiceModel("BigSetDemo",
			[]string{"0.0.0.0:2379"},
			GoEndpointBackendManager.EndPoint{
				Host:      "127.0.0.1",
				Port:      "18407",
				ServiceID: "BigSetDemo",
			})
	}
	return db
}

type UserServiceHandler struct {
}

func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{}
}

func (u UserServiceHandler) CreateUser(ctx context.Context, user *qlnv.User) (err error) {
	bskey := generic.TStringKey("user/" + user.UserId)
	var Db = GetDB()
	e := (*Db).BsMultiPut(bskey, []*generic.TItem{
		{Key: []byte("UserId"), Value: []byte(user.UserId)},
		{Key: []byte("Name"), Value: []byte(user.Name)},
		{Key: []byte("DateOfBirth"), Value: []byte(fmt.Sprintf("%d", user.DateOfBirth))},
		{Key: []byte("Address"), Value: []byte(user.Address)},
		{Key: []byte("UnitId"), Value: []byte(user.UnitId)},
	})
	if e != nil {
		return e
	}
	return nil
}

func (u UserServiceHandler) ReadUser(ctx context.Context, userId string) (r *qlnv.User, err error) {
	bskey := generic.TStringKey("user/" + userId)
	var Db = GetDB()
	info, e := (*Db).GetBigSetInfoByName(bskey)
	if e != nil {
		return nil, e
	}
	fmt.Println(info.String())
	return nil, nil
}
func (u UserServiceHandler) UpdateUser(ctx context.Context, user *qlnv.User) (err error) {
	return nil
}
func (u UserServiceHandler) DeleteUser(ctx context.Context, userId string) (err error) {
	return nil
}
func (u UserServiceHandler) GetUnitUser(ctx context.Context, userId string) (r *qlnv.Unit, err error) {
	return nil, nil
}
func (u UserServiceHandler) GetUserSortedByPage(ctx context.Context, userId string, numOfPages int32, sizeOfpage int32) (r []*qlnv.User, err error) {
	return nil, nil
}

// Unit Service

type UnitServiceHandler struct{}

func NewUnitServiceHandler() *UnitServiceHandler {
	return &UnitServiceHandler{}
}

func (unit UnitServiceHandler) CreateUnit(ctx context.Context, u *qlnv.Unit) (err error) {
	return nil
}

func (unit UnitServiceHandler) ReadUnit(ctx context.Context, unitId string) (r *qlnv.Unit, err error) {
	return nil, nil
}
func (unit UnitServiceHandler) UpdateUnit(ctx context.Context, u *qlnv.Unit) (err error) {
	return nil
}
func (unit UnitServiceHandler) DeleteUnit(ctx context.Context, userId string) (err error) {
	return nil
}
func (unit UnitServiceHandler) GetAllMemberOfUnit(ctx context.Context, unitId string) (r []*qlnv.User, err error) {
	return nil, nil
}
func (unit UnitServiceHandler) GetMembersByPage(ctx context.Context, unitId string, numOfPage int32, sizeOfPage int32) (r []*qlnv.User, err error) {
	return nil, nil
}
