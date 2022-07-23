package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID int
	UserName string
	CreateDate time.Time
	Gender string

}


func main() {
	db, err := gorm.Open(
		"mysql",
		"root:123@(127.0.0.1:3306)/python")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	user := User{1,"King",time.Now(),"男"}
	db.Create(&user)

}
