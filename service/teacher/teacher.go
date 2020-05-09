package teacher

import "model/model/teacher"

func GetBaseInfo(teacherId int64) (*teacher.BaseInfo, error) {
	return teacher.GetBaseInfo(teacherId)
}

func GetTechnology(technologyId int64) (*teacher.Technology, error) {
	return teacher.GetTechnology(technologyId)
}
