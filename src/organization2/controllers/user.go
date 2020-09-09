package controllers

import (
	"encoding/json"
	"fmt"
	"gen-go/qlnv"
	"organization2/models"

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
	var user qlnv.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println(err.Error() + " create user")
	}
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} qlnv.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body	body 	qlnv.User	true		"body for user content"
// @Success 200 {string}
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user qlnv.User
		e := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		if e != nil {
			fmt.Println(e.Error())
		}
		err := models.UpdateUser(&user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "update success"
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	err := models.DeleteUser(uid)
	if err != nil {
		u.Data["json"] = "Delete fail!"
	} else {
		u.Data["json"] = "delete success!"
	}
	u.ServeJSON()
}

// @Title UnitOfUser
// @Description get unit of the User
// @Param	uid		path 	string	true		"The uid you want to get Unit"
// @Success 200 {object} qlnv.Unit
// @Failure 403 uid is empty
// @router /get-unit/:uid [get]
func (u *UserController) GetUnitOfUser() {
	uid := u.GetString(":uid")
	unit, err := models.GetUnitOfUser(uid)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = unit
	}
	u.ServeJSON()
}

// @Title get by page
// @Description get page of Users (age)
// @Param	page 	query 	int		true		"parameter require"
// @Param	size	query 	int		true		"parameter require"
// @Param	type	query 	string	true		"parameter require"
// @Success 200 {object}
// @Failure 403 uid is empty
// @router /get-by-page/:page/:size/:type [get]
func (u *UserController) GetUserSortedByPage() {
	numOfPage, eNum := u.GetInt32(":page")
	sizeOfPage, eSize := u.GetInt32(":size")
	if eNum != nil || eSize != nil {
		u.Data["json"] = "Syntax error"
		u.ServeJSON()
		return
	}
	sortType := u.GetString(":type")
	unit, err := models.GetUserSortedByPage(numOfPage, sizeOfPage, sortType)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = unit
	}
	u.ServeJSON()
}
