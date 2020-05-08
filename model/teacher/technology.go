package teacher

import (
	"database/sql"
	"github.com/didi/gendry/builder"
	"model/service/database"
	"reflect"
)

// 教师科研成果
type Technology struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
	//……
}

const technologyTable = "technology"

func technologyErrorProcess(t *Technology, err error) (*Technology, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return t, nil
}

func GetTechnology(id int64) (*Technology, error) {
	_, fields := database.GetNameAndTagsT(reflect.TypeOf(new(Technology)))
	where := map[string]interface{}{
		"id": id,
	}
	query, args, err := builder.BuildSelect(technologyTable, where, fields)
	var b Technology
	err = technologyEngine().Get(&b, query, args...)
	return technologyErrorProcess(&b, err)
}
