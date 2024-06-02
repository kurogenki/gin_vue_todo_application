package model

import (
	"gorm.io/gorm"
)

type Task struct {
  gorm.Model
  Title string `gorm:"column:title"`
  Description string `gorm:"column:description"`
}

func TaskIndex () (datas []Task){
  result := Db.Find(&datas)
  if result.Error != nil {
		panic(result.Error.Error())
	}
	return
}

func (t *Task) CreateTask () {
  result := Db.Create(t)
	if result.Error != nil {
    panic(result.Error)
  }
}
