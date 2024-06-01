package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "github.com/joho/godotenv"
)

var Db *gorm.DB

func Init() {
  err := godotenv.Load(".env")
  if err != nil {
    fmt.Println("読み込み出来ませんでした: %v", err)
  }

  user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
  db_host := os.Getenv("DB_HOST")

	var path string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pw, db_host, db_name)
  db, err := gorm.Open(mysql.Open(path))

  if err != nil {
  panic("failed to connect database")
  } else {
    fmt.Println("db connected!!")
    db.AutoMigrate(Task{})
  }
}