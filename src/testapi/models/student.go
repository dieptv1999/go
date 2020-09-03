package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Students map[string]*Student
)

type Student struct {
	StudentId string
	Name      string
	Gender    string
	classes   []string
}

func init() {
	Students = make(map[string]*Student)
	Students["hjkhsbnmn123"] = &Student{"hjkhsbnmn123", "100", "Nam", []string{}}
	Students["mjjkxsxsaa23"] = &Student{"mjjkxsxsaa23", "101", "Nữ", []string{}}
}

func AddStudent(student Student) (StudentId string) {
	student.StudentId = strconv.FormatInt(time.Now().UnixNano(), 10)
	Students[student.StudentId] = &student
	return student.StudentId
}

func GetStudent(StudentId string) (student *Student, err error) {
	if v, ok := Students[StudentId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAllStudents() map[string]*Student {
	return Students
}

func UpdateStudent(StudentId string, sstudent *Student) (student *Student, err2 error) {
	if s, ok := Students[StudentId]; ok {
		if sstudent.Gender != "" {
			s.Gender = sstudent.Gender
		}
		if sstudent.Name != "" {
			s.Name = sstudent.Name
		}
		if sstudent.classes != nil {
			s.classes = sstudent.classes
		}
		return s, nil
	}
	return nil, errors.New("User Not Exist")
}

func DeleteStudent(StudentId string) {
	delete(Students, StudentId)
}

func AddClassToStudent(StudentId string, ClassId string) (result bool) {
	if u, ok := Students[StudentId]; ok {
		u.classes = append(u.classes, ClassId)
		return true
	}
	return false
}

func GetAllClassesOfStudent(StudentId string) []string {
	if u, ok := Students[StudentId]; ok {
		return u.classes
	}
	return nil
}
