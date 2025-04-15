package main

import (
	"github.com/QuocB-HC/Gorm-Tutorials/database"
	// "github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	database.ConnectDB()

	// r := gin.Default()

	// r.Run()

	user_list := database.GetAll()
	for _, user := range user_list {
		fmt.Println("Name: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("===================")
	}
}
