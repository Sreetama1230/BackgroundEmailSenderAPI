package db

import (
	"new_start/BackgroundEmailSenderAPI/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and migrates the Email model.
// It panics if the connection or migration fails.
func InitDB() {

	var err error
	dsn := "sreetama:password@tcp(localhost:3306)/emails?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = DB.AutoMigrate(&models.Email{})
	if err != nil {
		panic("failed to migrate database")
	}

}
