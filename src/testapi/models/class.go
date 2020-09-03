package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Classes map[string]*Class
)

type Class struct {
	ClassId  string
	Name     string
	Teacher  string
	Students []string
}

func init() {
	Classes = make(map[string]*Class)
	Classes["hjkhsbnmn123"] = &Class{"hjkhsbnmn123", "100", "TVD", []string{}}
	Classes["mjjkxsxsaa23"] = &Class{"mjjkxsxsaa23", "101", "TVT", []string{}}
}

func AddClass(Class Class) (ClassId string) {
	Class.ClassId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Classes[Class.ClassId] = &Class
	return Class.ClassId
}

func GetClass(ClassId string) (Class *Class, err error) {
	if v, ok := Classes[ClassId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAllClasses() map[string]*Class {
	return Classes
}

func UpdateClass(ClassId string, sClass *Class) (Class *Class, err2 error) {
	if s, ok := Classes[ClassId]; ok {
		if sClass.Teacher != "" {
			s.Teacher = sClass.Teacher
		}
		if sClass.Name != "" {
			s.Name = sClass.Name
		}
		if sClass.Students != nil {
			s.Students = sClass.Students
		}
		return s, nil
	}
	return nil, errors.New("ClassId Not Exist")
}

func DeleteClass(ClassId string) {
	delete(Classes, ClassId)
}

func AddStudentToClass(ClassId string, StudentId string) (result bool) {
	if u, ok := Classes[ClassId]; ok {
		u.Students = append(u.Students, ClassId)
		return true
	}
	return false
}

func GetAllStudentsOfClass(ClassId string) []string {
	if u, ok := Classes[ClassId]; ok {
		return u.Students
	}
	return nil
}
