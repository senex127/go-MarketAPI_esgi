package main

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	//"github.com/gin-gonic/gin"
)

func main(){
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/go-exam?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(db)

	//db.AutoMigrate(&task.Task{})

	fmt.Println("Hello")

}