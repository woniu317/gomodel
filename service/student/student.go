package student

import (
	"model/model/student"
)

func GetStudent(studentId int64) (*student.Student, error) {
	return student.GetStudent(studentId)
}

func GetFamily(familyId int64) (*student.Family, error) {
	return student.GetFamily(familyId)
}
