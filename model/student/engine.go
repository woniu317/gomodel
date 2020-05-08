package student

import (
	"model/model"
	"model/service/database/dbclient"
)

func familyEngine() model.DbInterface {
	return dbclient.GetStudentEngine()
}

func baseEngine() model.DbInterface {
	return dbclient.GetStudentEngine()
}
