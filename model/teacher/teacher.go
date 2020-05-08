package teacher

import (
	"database/sql"
	"reflect"

	"model/service/database"

	"github.com/didi/gendry/builder"
)

type BaseInfo struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int32  `db:"age"`
	Sex  int8   `db:"sex"`
}

const teacherBaseInfo = "teacher_base_info"

type BaseInfoQuery struct {
	Id int64
	//........
}

func baseErrorProcess(b *BaseInfo, err error) (*BaseInfo, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return b, nil
}

// 根据teacherId获取教师信息
func GetBaseInfo(teacherId int64) (*BaseInfo, error) {
	_, fields := database.GetNameAndTagsT(reflect.TypeOf(new(BaseInfo)))
	where := map[string]interface{}{
		"id": teacherId,
	}
	query, args, err := builder.BuildSelect(teacherBaseInfo, where, fields)
	var b BaseInfo
	err = baseEngine().Get(&b, query, args...)
	return baseErrorProcess(&b, err)
}
