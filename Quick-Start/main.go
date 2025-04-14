package main

import (
	"fmt"

	"github.com/QuocB-HC/Gorm-Tutorials/database"
)

func main() {
	database.Init()

	// fmt.Println(database.DeleteById(1))
	
	products, message := database.GetAll()

	if products == nil {
		fmt.Println(message)
	}

	for _, product := range products {
		fmt.Println("Product: " + product.Code + " with price: " + fmt.Sprint(product.Price))
	}
}
