package controller

import (
	model "gin_vue_todo_application/Model"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// タスクの一覧を取得
func GetAllTasks(c *gin.Context) {
  datas := model.TaskIndex()
  c.JSON(http.StatusOK, gin.H{"datas": datas})
}

// タスクの詳細を取得
func ShowTask(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  datas := model.ShowTask(id)
  c.JSON(http.StatusOK, gin.H{"datas": datas})
}

// タスクを新規作成
func CreateTask(c *gin.Context) {
  title := c.PostForm("title")
  description := c.PostForm("description")
  data := model.Task{Title: title, Description: description}
  data.CreateTask()
  c.JSON(http.StatusOK, gin.H{"datas": "タスクを作成しました"})
}

// タスクを削除
func DeleteTask(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  model.DeleteTask(id)
  c.JSON(http.StatusOK, gin.H{"datas": "タスクを削除しました"})
}