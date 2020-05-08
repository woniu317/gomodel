package student

import (
	"database/sql"
	"reflect"

	"model/service/database"

	"github.com/didi/gendry/builder"
)

type Student struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

var studentName = "student"

func studentErrorProcess(s *Student, err error) (*Student, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return s, nil
}

func GetStudent(id int64) (*Student, error) {
	_, fields := database.GetNameAndTagsT(reflect.TypeOf(new(Student)))
	where := map[string]interface{}{
		"id": id,
	}
	query, args, err := builder.BuildSelect(studentName, where, fields)
	var b Student
	err = baseEngine().Get(&b, query, args...)
	return studentErrorProcess(&b, err)
}
