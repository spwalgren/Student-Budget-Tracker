package database

import (
	"budget-tracker/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Initialize(dbname string) {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",models.Username,models.Password,models.Url, dbname)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	err = database.AutoMigrate(&models.UserInfo{}, &models.Transaction{}, &models.Budget{}, &models.Progress{})
        if err != nil {
                return
        }
	DB = database
}
