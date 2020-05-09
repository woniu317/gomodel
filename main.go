package main

import (
	"fmt"

	"model/service/teacher"

	"github.com/YunzhanghuOpen/glog"
)

func main() {
	var teacherId int64
	teacher, err := teacher.GetBaseInfo(teacherId)
	if err != nil {
		glog.Info("Query teacher error!")
		return
	}
	if teacher == nil {
		glog.Info("No teacher")
		return
	}
	fmt.Println(teacher)
}
