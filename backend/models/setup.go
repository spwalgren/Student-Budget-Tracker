package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/budget_tracker?charset=utf8mb4",Username,Password,Url)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	err = database.AutoMigrate(&UserInfo{})
        if err != nil {
                return
        }
	DB = database
}