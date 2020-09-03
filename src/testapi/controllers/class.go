package controllers

import (
	"encoding/json"
	"testapi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type ClassController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ClassController) Post() {
	var class models.Class
	json.Unmarshal(u.Ctx.Input.RequestBody, &class)
	uid := models.AddClass(class)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *ClassController) GetAll() {
	students := models.GetAllClasses()
	u.Data["json"] = students
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *ClassController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		student, err := models.GetClass(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = student
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *ClassController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var class models.Class
		json.Unmarshal(u.Ctx.Input.RequestBody, &class)
		uu, err := models.UpdateClass(uid, &class)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *ClassController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteClass(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

//APIVersion 1.0.0
// @Title AddStudent
// @Description add student to class
// @Param	StudentId ClassId path 	string	true		"The classId"
// @Success 200 {string} add success!
// @Failure 403 uid is empty
// @router /AddStudent/:ClassId/:StudentId [get]
func (u *ClassController) AddStudent() {
	ClassId := u.GetString(":ClassId")
	StudentId := u.GetString(":StudentId")
	a := models.AddStudentToClass(ClassId, StudentId)
	if a == true {
		u.Data["json"] = "Add success!"
	} else {
		u.Data["json"] = "Fail!"
	}
	u.ServeJSON()
}

//APIVersion 1.0.0
// @Title GetAllStudentOfClass
// @Description get all students of class
// @Param	ClassId path 	string	true		"The classId"
// @Success 200 {string} get success!
// @Failure 403 uid is empty
// @router /GetAllStudentOfClass/:ClassId [get]
func (u *ClassController) GetAllStudentOfClass() {
	ClassId := u.GetString(":ClassId")
	Students := models.GetAllStudentsOfClass(ClassId)
	u.Data["json"] = Students
	u.ServeJSON()
}
