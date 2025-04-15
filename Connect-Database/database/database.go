package database

import (
	"fmt"

	"github.com/QuocB-HC/Gorm-Tutorials/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var user models.User

var db *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}

func GetAll() []models.User {
	var user_list []models.User

	result := db.Find(&user_list)

	if result.Error != nil {
		panic("failed to get all users")
	} else if len(user_list) == 0 {
		panic("no users found")
	}

	return user_list
}
