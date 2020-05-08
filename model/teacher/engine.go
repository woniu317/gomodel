package teacher

import (
	"model/model"
	"model/service/database/dbclient"
)

func baseEngine() model.DbInterface {
	return dbclient.GetTeacherEngine()
}

func technologyEngine() model.DbInterface {
	return dbclient.GetTeacherEngine()
}
