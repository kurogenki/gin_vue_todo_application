package model

import (
	"gorm.io/gorm"
)

type Task struct {
  gorm.Model
  Title string `gorm:"column:title"`
  Description string `gorm:"column:description"`
}

// タスクの一覧を取得
func TaskIndex () (datas []Task){
  result := Db.Find(&datas)
  if result.Error != nil {
		panic(result.Error.Error())
	}
	return
}

// タスクの詳細を取得
func ShowTask (id int) (datas Task){
  result := Db.First(&datas, id)
  if result.Error != nil {
		panic(result.Error.Error())
	}
	return
}

// タスクを新規で作成する
func (t *Task) CreateTask () {
  result := Db.Create(t)
	if result.Error != nil {
    panic(result.Error)
  }
}

// タスクを削除する
func DeleteTask (id int) {
  result := Db.Delete(&Task{}, id)
	if result.Error != nil {
    panic(result.Error)
  }
}
