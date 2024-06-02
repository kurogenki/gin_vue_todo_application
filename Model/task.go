package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
  Id  uint
  Title string
  Description *string
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt
}

var allTasks []Task

func TaskIndex () (datas []Task) {
  result := Db.Find(&allTasks)
  if result.Error != nil {
		panic(result.Error.Error())
	}
	return 
}