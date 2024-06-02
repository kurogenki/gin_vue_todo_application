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

// func TaskIndex () (datas []Task) {
//   // var allTasks []Task
//   fmt.Println(Db)
//   result := Db.Find(&datas)
//   if result.Error != nil {
//     fmt.Println("errorが起きています")
// 		panic(result.Error)
// 	}
// 	return
// }