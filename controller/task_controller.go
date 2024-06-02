package controller

import (
	model "gin_vue_todo_application/Model"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
  datas := model.TaskIndex()
  c.JSON(http.StatusOK, gin.H{"datas": datas})
}

func CreateTask(c *gin.Context) {
  title := c.PostForm("title")
  description := c.PostForm("description")
  data := model.Task{Title: title, Description: description}
  data.CreateTask()
  c.JSON(http.StatusOK, gin.H{"datas": "タスクを作成しました"})
}