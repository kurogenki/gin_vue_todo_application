package main

import (
	// "net/http"
	model "gin_vue_todo_application/Model"
	router "gin_vue_todo_application/controller"
)

func main() {
  model.Init()

  router := router.GetRouter()
  router.Run("localhost:8080")
}