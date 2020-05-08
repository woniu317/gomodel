package main

import (
	"fmt"

	"model/model/teacher"

	"github.com/YunzhanghuOpen/glog"
)

func main() {
	var teacherId int64
	teacher, err := teacher.GetBaseInfo(teacherId)
	if err != nil {
		glog.Info("query teacher error!")
		return
	}
	if teacher == nil {
		glog.Info("no teacher")
		return
	}
	fmt.Println(teacher)
}
