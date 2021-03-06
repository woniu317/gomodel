package main

import (
	"fmt"
	"log"

	"model/service/teacher"
)

func main() {
	var teacherId int64
	teacher, err := teacher.GetBaseInfo(teacherId)
	if err != nil {
		log.Println("Query teacher error!")
		return
	}
	if teacher == nil {
		log.Println("No teacher")
		return
	}
	fmt.Println(teacher)
}
