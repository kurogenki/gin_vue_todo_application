package main

import (
  // "net/http"
  controller "gin_vue_todo_application/controllers"
)

func main() {
  router := controller.GetRouter()
  router.Run("localhost:8080")
}