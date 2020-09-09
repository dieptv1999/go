package controllers

import (
	"encoding/json"
	"fmt"
	"gen-go/qlnv"
	"organization2/models"

	"github.com/astaxie/beego"
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
	if uid != "" {
		var Unit qlnv.Unit
		e := json.Unmarshal(u.Ctx.Input.RequestBody, &Unit)
		if e != nil {
			fmt.Println(e.Error())
		}
		err := models.UpdateUnit(&Unit)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "update success"
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the Unit
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UnitController) Delete() {
	uid := u.GetString(":uid")
	err := models.DeleteUnit(uid)
	if err != nil {
		u.Data["json"] = "Delete fail!"
	} else {
		u.Data["json"] = "delete success!"
	}
	u.ServeJSON()
}

// @Title getAllMember
// @Description get all members of the Unit
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {object} list user
// @Failure 403 uid is empty
// @router /get-all-member/:uid [get]
func (u *UnitController) GetAllMemberOfUnit() {
	uid := u.GetString(":uid")
	listUser, err := models.GetAllMemberOfUnit(uid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = listUser
	}
	u.ServeJSON()
}
