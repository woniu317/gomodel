package database

import (
	"reflect"
	"strings"
)

func GetNameAndTagsT(t reflect.Type) (names, tags []string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		k := f.Type.Kind()
		if k == reflect.Struct && f.Anonymous {
			tempNames, tempTags := GetNameAndTagsT(f.Type)
			tags = append(tags, tempTags...)
			names = append(names, tempNames...)
		} else {
			tag, ok := f.Tag.Lookup("db")
			if ok {
				tags = append(tags, strings.Split(tag, ",")[0])
				names = append(names, f.Name)
			}
		}
	}
	return
}
