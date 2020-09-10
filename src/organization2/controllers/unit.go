package controllers

import (
	"encoding/json"
	"fmt"
	"gen-go/qlnv"
	"strconv"

	"github.com/astaxie/beego"
	"organization2/models"
)

// Operations about Units
type UnitController struct {
	beego.Controller
}

// @Title CreateUnit
// @Description create Units
// @Param	body		body 	qlnv.Unit	true		"body for Unit content"
// @Success 200 {int} qlnv.Unit.UnitId
// @Failure 403 body is empty
// @router / [post]
func (u *UnitController) Post() {
	var Unit qlnv.Unit
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &Unit)
	if err != nil {
		fmt.Println(err.Error() + " create Unit")
	}
	uid := models.AddUnit(Unit)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title Get
// @Description get Unit by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} qlnv.Unit
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UnitController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		Unit, err := models.GetUnit(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = Unit
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the Unit
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body	body 	qlnv.Unit	true		"body for Unit content"
// @Success 200 {string}
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UnitController) Put() {
	uid := u.GetString(":uid")
	defer u.ServeJSON()
	if uid != "" {
		var Unit qlnv.Unit
		e := json.Unmarshal(u.Ctx.Input.RequestBody, &Unit)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		err := models.UpdateUnit(&Unit)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "update success"
		}
	}
}

// @Title Delete
// @Description delete the Unit
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UnitController) Delete() {
	defer u.ServeJSON()
	uid := u.GetString(":uid")
	err := models.DeleteUnit(uid)
	if err != nil {
		u.Data["json"] = "Delete fail!"
	} else {
		u.Data["json"] = "delete success!"
	}
}

// @Title getAllMember
// @Description get all members of the Unit
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {object} list-user
// @Failure 403 uid is empty
// @router /all-member/:uid [get]
func (u *UnitController) GetAllMemberOfUnit() {
	defer u.ServeJSON()
	uid := u.GetString(":uid")
	listUser, err := models.GetAllMemberOfUnit(uid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = listUser
	}
}

// @Title get member  by page
// @Description get members page of Unit
// @Param	unitId	query 	string	true		"parameter require"
// @Param	page 	query 	string	true		"parameter require"
// @Param	size	query 	string	true		"parameter require"
// @Success 200 {object} members
// @Failure 403 unitId page size is empty
// @router /member-by-page [get]
func (u *UnitController) GetMembersByPage() {
	defer u.ServeJSON()
	fmt.Println(u.GetString("unitId"))
	unitId := u.GetString("unitId")
	numOfPage, eNum := strconv.Atoi(u.GetString("page"))
	sizeOfPage, eSize := strconv.Atoi(u.GetString("size"))
	if eNum != nil || eSize != nil {
		u.Data["json"] = "Syntax error"
		return
	}
	unit, err := models.GetMembersByPage(unitId, int32(numOfPage), int32(sizeOfPage))
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = unit
	}
}

// @Title get unit  by page
// @Description get unit page
// @Param	page 	query 	string	true		"parameter require"
// @Param	size	query 	string	true		"parameter require"
// @Param	type	query 	string	true		"parameter require"
// @Success 200 {object} members
// @Failure 403 page size type is empty
// @router /by-page [get]
func (u UnitController) GetUnitsByPage() {
	defer u.ServeJSON()
	numOfPage, eNum := strconv.Atoi(u.GetString("page"))
	sizeOfPage, eSize := strconv.Atoi(u.GetString("size"))
	if eNum != nil || eSize != nil {
		u.Data["json"] = "Syntax error"
		return
	}
	sortType := u.GetString("type")
	unit, err := models.GetUnitByPage(int32(numOfPage), int32(sizeOfPage),sortType)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = unit
	}
}
