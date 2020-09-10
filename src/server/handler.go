package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gen-go/qlnv"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var db = StringBigsetService.NewStringBigsetServiceModel("BigSetDemo",
	[]string{"0.0.0.0:2379"},
	GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "18407",
		ServiceID: "BigSetDemo",
	})

type UserServiceHandler struct {
}

func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{}
}

func (u UserServiceHandler) CreateUser(ctx context.Context, user *qlnv.User) (err error) {
	bskey := generic.TStringKey("user")
	jsonByte, eJson := json.Marshal(*user)
	if eJson != nil {
		fmt.Println("don't convert jsonByte")
		return eJson
	}
	e := db.BsPutItem(bskey, &generic.TItem{
		Key:   []byte(user.UserId),
		Value: jsonByte,
	})
	eUnit := db.BsPutItem(generic.TStringKey(user.UnitId), &generic.TItem{
		Key:   []byte(user.UserId),
		Value: []byte(user.UserId),
	})
	if eUnit != nil {
		return eUnit
	}
	if e != nil {
		return e
	}
	return nil
}

func (u UserServiceHandler) ReadUser(ctx context.Context, userId string) (r *qlnv.User, err error) {
	bskey := generic.TStringKey("user")
	a, e := db.BsGetItem(bskey, generic.TItemKey(userId))
	if e != nil {
		return nil, e
	}
	var userTmp qlnv.User
	eJson := json.Unmarshal(a.Value, &userTmp)
	if eJson != nil {
		return nil, errors.New("Don't convert json to User")
	}
	return &userTmp, nil
}
func (u UserServiceHandler) UpdateUser(ctx context.Context, user *qlnv.User) (err error) {
	bskey := generic.TStringKey("user")
	info, _ := db.BsGetItem(bskey, generic.TItemKey(user.UserId))
	if info == nil {
		return errors.New("User not exist!")
	}
	jsonByte, eJson := json.Marshal(*user)
	if eJson != nil {
		return errors.New("Don't convert jsonByte")
	}
	e := db.BsPutItem(bskey, &generic.TItem{
		Key:   []byte(user.UserId),
		Value: jsonByte,
	})
	if e != nil {
		return e
	}
	return nil
}
func (u UserServiceHandler) DeleteUser(ctx context.Context, userId string) (err error) {
	fmt.Println("delete")
	bskey := generic.TStringKey("user")
	e := db.BsRemoveItem(bskey, generic.TItemKey(userId))
	if e != nil {
		return e
	}
	return nil
}

func (u UserServiceHandler) GetUnitUser(ctx context.Context, userId string) (r *qlnv.Unit, err error) {
	bskey := generic.TStringKey("user")
	infoUser, _ := db.BsGetItem(bskey, generic.TItemKey(userId))
	if infoUser == nil {
		return nil, errors.New("User not exist!")
	}
	var userTmp qlnv.User
	eJson := json.Unmarshal(infoUser.Value, &userTmp)
	if eJson != nil {
		return nil, errors.New("Don't convert json to User")
	}
	infoUnit, _ := db.BsGetItem("unit", generic.TItemKey(userTmp.UnitId))
	if infoUnit == nil {
		return nil, errors.New("Unit of this User don't exist")
	}
	var unitTmp qlnv.Unit
	eJson = json.Unmarshal(infoUnit.Value, &unitTmp)
	if eJson != nil {
		return nil, errors.New("Don't convert json to Unit")
	}
	return &unitTmp, nil
}
func (u UserServiceHandler) GetUserSortedByPage(ctx context.Context, numOfPages int32, sizeOfpage int32, sortType string) (r []*qlnv.User, err error) {
	bskey := generic.TStringKey("user")
	var listItem []*generic.TItem
	var e error
	if sortType == "inc" {
		listItem, e = db.BsGetSlice(bskey, (numOfPages-1)*sizeOfpage, sizeOfpage)
	} else {
		listItem, e = db.BsGetSliceR(bskey, (numOfPages-1)*sizeOfpage, sizeOfpage)
	}
	if e != nil {
		return nil, e
	}
	if listItem == nil {
		return nil, errors.New("Page don't exist")
	}
	var listUser []*qlnv.User
	for _, item := range listItem {
		var userTmp qlnv.User
		eJson := json.Unmarshal(item.Value, &userTmp)
		if eJson != nil {
			fmt.Println("error in sort by page")
			return nil, eJson
		} else {
			listUser = append(listUser, &userTmp)
		}
	}
	return listUser, nil
}

// Unit Service

type UnitServiceHandler struct{}

func NewUnitServiceHandler() *UnitServiceHandler {
	return &UnitServiceHandler{}
}

func (unit UnitServiceHandler) CreateUnit(ctx context.Context, u *qlnv.Unit) (err error) {
	bskey := generic.TStringKey("unit")
	jsonByte, eJson := json.Marshal(*u)
	if eJson != nil {
		fmt.Println("don't convert jsonByte")
	}
	e := db.BsPutItem(bskey, &generic.TItem{
		Key:   []byte(u.UnitId),
		Value: jsonByte,
	})
	if e != nil {
		return e
	}
	return nil
}

func (unit UnitServiceHandler) ReadUnit(ctx context.Context, unitId string) (r *qlnv.Unit, err error) {
	bskey := generic.TStringKey("unit")
	a, e := db.BsGetItem(bskey, generic.TItemKey(unitId))
	if e != nil {
		return nil, e
	}
	var unitTmp qlnv.Unit
	eJson := json.Unmarshal(a.Value, &unitTmp)
	if eJson != nil {
		return nil, errors.New("Don't convert json to Unit")
	}
	return &unitTmp, nil
}
func (unit UnitServiceHandler) UpdateUnit(ctx context.Context, u *qlnv.Unit) (err error) {
	bskey := generic.TStringKey("unit")
	info, _ := db.BsGetItem(bskey, generic.TItemKey(u.UnitId))
	if info == nil {
		return errors.New("User not exist!")
	}
	jsonByte, eJson := json.Marshal(*u)
	if eJson != nil {
		return errors.New("Don't convert jsonByte")
	}
	e := db.BsPutItem(bskey, &generic.TItem{
		Key:   []byte(u.UnitId),
		Value: jsonByte,
	})
	if e != nil {
		return e
	}
	return nil
}
func (unit UnitServiceHandler) DeleteUnit(ctx context.Context, unitId string) (err error) {
	bskey := generic.TStringKey("unit")
	e := db.BsRemoveItem(bskey, generic.TItemKey(unitId))
	if e != nil {
		return e
	}
	return nil
}
func (unit UnitServiceHandler) GetAllMemberOfUnit(ctx context.Context, unitId string) (r []*qlnv.User, err error) {
	bskey := generic.TStringKey("user")
	len, e := db.GetTotalCount(bskey)
	if e != nil {
		return nil, e
	}
	var listUser []*qlnv.User
	var i int64
	var sizePage int64 = 10
	for i = 0; i < len; i += sizePage {
		listItem, eItem := db.BsGetSlice(bskey, int32(i), int32(sizePage))
		if eItem != nil {
			fmt.Println("get list item error")
		} else {
			for _, item := range listItem {
				var userTmp qlnv.User
				eJson := json.Unmarshal(item.Value, &userTmp)
				if eJson != nil {
					fmt.Println("error in sort by page")
				} else {
					if userTmp.UnitId == unitId {
						listUser = append(listUser, &userTmp)
					}
				}
			}
		}
	}
	return listUser, nil
}
func (unit UnitServiceHandler) GetMembersByPage(ctx context.Context, unitId string, numOfPage int32, sizeOfPage int32) (r []*qlnv.User, err error) {
	bskey := generic.TStringKey(unitId)
	info, e := db.GetBigSetInfoByName(bskey)
	if info == nil {
		return nil, e
	}
	listItem, e := db.BsGetSlice(bskey, (numOfPage-1)*sizeOfPage, sizeOfPage)
	if e != nil {
		return nil, e
	}
	if listItem == nil || len(listItem) == 0 {
		return nil, errors.New("Page don't exist")
	}
	var listUser []*qlnv.User
	for _, item := range listItem {
		userData, _ := db.BsGetItem("user", item.Value)
		if userData == nil {
			continue
		}
		var userTmp qlnv.User
		eJson := json.Unmarshal(userData.Value, &userTmp)
		if eJson != nil {
			fmt.Println("error in sort by page")
			return nil, eJson
		} else {
			listUser = append(listUser, &userTmp)
		}
	}
	return listUser, nil
}
func (unit UnitServiceHandler) GetUnitsByPage(ctx context.Context, numOfPages int32, sizeOfpage int32, sortType string) (r []*qlnv.Unit, err error) {
	bskey := generic.TStringKey("unit")
	var listItem []*generic.TItem
	var e error
	if sortType == "inc" {
		listItem, e = db.BsGetSlice(bskey, (numOfPages-1)*sizeOfpage, sizeOfpage)
	} else {
		listItem, e = db.BsGetSliceR(bskey, (numOfPages-1)*sizeOfpage, sizeOfpage)
	}
	if e != nil {
		return nil, e
	}
	if listItem == nil {
		return nil, errors.New("Page don't exist")
	}
	var listUnit []*qlnv.Unit
	for _, item := range listItem {
		var unitTmp qlnv.Unit
		eJson := json.Unmarshal(item.Value, &unitTmp)
		if eJson != nil {
			fmt.Println("error in sort by page")
			return nil, eJson
		} else {
			listUnit = append(listUnit, &unitTmp)
		}
	}
	return listUnit, nil
}
