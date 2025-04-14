package main

import (
	"fmt"

	"github.com/QuocB-HC/Gorm-Tutorials/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Kết nối đến SQLite (file test.db sẽ được tạo nếu chưa có)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to database!")

	// Auto migrate
	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		panic("Failed to migrate schema")
	}

	fmt.Println("Migrated successfully!")

	// Tạo dữ liệu
	post := models.Post{
		Title:     "Hello GORM",
		Content:   "Đây là bài post mẫu.",
		AuditInfo: models.AuditInfo{CreatedBy: "admin", UpdatedBy: "admin"},
	}

	result := db.Create(&post)
	if result.Error != nil {
		panic("Failed to create post")
	}

	fmt.Println("Post created with ID:", post.ID)

	// Truy vấn dữ liệu
	var posts []models.Post
	db.Find(&posts)
	fmt.Println("All Posts:")
	for _, p := range posts {
		fmt.Printf("- %s (created by: %s)\n", p.Title, p.CreatedBy)
	}
}
