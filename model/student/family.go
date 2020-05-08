package student

import (
	"database/sql"
	"github.com/didi/gendry/builder"
	"model/service/database"
	"reflect"
)

type Family struct {
	Id      int64
	Address string
	//……
}

const familyTable = "family"

func familyErrorProcess(f *Family, err error) (*Family, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return f, nil
}

func GetFamily(id int64) (*Family, error) {
	_, fields := database.GetNameAndTagsT(reflect.TypeOf(new(Family)))
	where := map[string]interface{}{
		"id": id,
	}
	query, args, err := builder.BuildSelect(familyTable, where, fields)
	var b Family
	err = familyEngine().Get(&b, query, args...)
	return familyErrorProcess(&b, err)
}
