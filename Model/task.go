package model

import (
	"time"
)

type Task struct {
  Id  uint
  Title string
  Description *string
  CreatedAt time.Time
  updatedAt time.Time
}

// func Index () {
//   var allTasks []Task
//   model.Db.Find(&allTasks)
// }