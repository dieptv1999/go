package controllers

import (
	"encoding/json"
	"fmt"
	"gen-go/qlnv"
	"organization2/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	qlnv.User	true		"body for user content"
// @Success 200 {int} qlnv.User.UserId
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	defer u.ServeJSON()
	var user qlnv.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Data["json"] = err.Error()
		return
	}
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} qlnv.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	defer u.ServeJSON()
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body	body 	qlnv.User	true		"body for user content"
// @Success 200 {string}
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	defer u.ServeJSON()
	uid := u.GetString(":uid")
	if uid != "" {
		var user qlnv.User
		e := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		if e != nil {
			u.Data["json"] = e.Error()
			return
		}
		err := models.UpdateUser(&user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "update success"
		}
	}
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	defer u.ServeJSON()
	uid := u.GetString(":uid")
	err := models.DeleteUser(uid)
	if err != nil {
		u.Data["json"] = "Delete fail!"
	} else {
		u.Data["json"] = "delete success!"
	}
}

// @Title UnitOfUser
// @Description get unit of the User
// @Param	uid		path 	string	true		"The uid you want to get Unit"
// @Success 200 {object} qlnv.Unit
// @Failure 403 uid is empty
// @router /unit/:uid [get]
func (u *UserController) GetUnitOfUser() {
	defer u.ServeJSON()
	uid := u.GetString(":uid")
	unit, err := models.GetUnitOfUser(uid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = unit
	}
}

// @Title get by page
// @Description get page of Users (age)
// @Param	page 	query 	string		true		"parameter require"
// @Param	size	query 	string		true		"parameter require"
// @Param	type	query 	string	true		"parameter require"
// @Success 200 {object} test
// @Failure 403 uid is empty
// @router /by-page [get]
func (u *UserController) GetUserSortedByPage() {
	defer u.ServeJSON()
	fmt.Println(u.GetString(":page"))
	numOfPage, eNum := strconv.Atoi(u.GetString("page"))
	sizeOfPage, eSize := strconv.Atoi(u.GetString("size"))
	if eNum != nil || eSize != nil {
		u.Data["json"] = "Syntax error"
		return
	}
	sortType := u.GetString("type")
	unit, err := models.GetUserSortedByPage(int32(numOfPage), int32(sizeOfPage), sortType)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = unit
	}
}
