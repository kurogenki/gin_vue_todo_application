package model

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
  // .envの設定
  env_err := godotenv.Load(".env")
  if env_err != nil {
    fmt.Println("読み込み出来ませんでした: %v", env_err)
  }

  user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
  db_host := os.Getenv("DB_HOST")

  // アクセスするデータベースの設定
	var path string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pw, db_host, db_name)

  dialector := mysql.Open(path)

  var err error
	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
  fmt.Println("db connected!!")
  Db.AutoMigrate(Task{})
  
  // if err != nil {
  //   panic("failed to connect database")
  // } else {
  //   fmt.Println("db connected!!")
  //   Db.AutoMigrate(Task{})
  // }
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}