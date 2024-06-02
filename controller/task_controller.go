package controller

import (
	model "gin_vue_todo_application/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
  datas := model.TaskIndex()
  c.JSON(http.StatusOK, gin.H{"datas": datas})
}