package controllers

import (
	"encoding/json"
	"testapi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type StudentController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *StudentController) Post() {
	var student models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &student)
	uid := models.AddStudent(student)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *StudentController) GetAll() {
	students := models.GetAllStudents()
	u.Data["json"] = students
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *StudentController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		student, err := models.GetStudent(uid)
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
func (u *StudentController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var student models.Student
		json.Unmarshal(u.Ctx.Input.RequestBody, &student)
		uu, err := models.UpdateStudent(uid, &student)
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
func (u *StudentController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteStudent(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title AddClass
// @Description add class to user
// @Param	StudentId ClassId path 	string	true		"The classId"
// @Success 200 {string} add success!
// @Failure 403 uid is empty
// @router /addClass/:StudentId/:ClassId [get]
func (u *StudentController) AddClass() {
	StudentId := u.GetString(":StudentId")
	ClassId := u.GetString(":ClassId")
	a := models.AddClassToStudent(StudentId, ClassId)
	if a == true {
		u.Data["json"] = "Add success!"
	} else {
		u.Data["json"] = "Fail!"
	}
	u.ServeJSON()
}

// @Title GetAllClassesOfStudent
// @Description get all class of student
// @Param	StudentId path 	string	true	"The classId"
// @Success 200 {string} get success!
// @Failure 403 uid is empty
// @router /GetAllClassesOfStudent/:StudentId [get]
func (u *StudentController) GetAllClassesOfStudent() {
	StudentId := u.GetString(":StudentId")
	classes := models.GetAllClassesOfStudent(StudentId)
	u.Data["json"] = classes
	u.ServeJSON()
}