package controller

import (
	"fmt"
	model "gin_vue_todo_application/Model"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetRouter() *gin.Engine {
  r := gin.Default()

  // ここからCorsの設定
  r.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
        "http://localhost:3000",
    },
    // アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
    AllowMethods: []string{
        "POST",
        "GET",
        "OPTIONS",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: true,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
  }))


  err := godotenv.Load(".env")
  if err != nil {
    fmt.Println("読み込み出来ませんでした: %v", err)
  }

  user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
  db_host := os.Getenv("DB_HOST")

	var path string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pw, db_host, db_name)
  Db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  } else {
    fmt.Println("db connected!!")
    Db.AutoMigrate(model.Task{})
  }

  r.GET("/test", func(c *gin.Context) {
    fmt.Println("test")
    c.JSON(200, gin.H{"message": "test",})
  })

  r.GET("/tasks", GetAllTasks)

  return r
}