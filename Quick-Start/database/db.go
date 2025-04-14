package database

import (
	"github.com/QuocB-HC/Gorm-Tutorials/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

var product model.Product
var db *gorm.DB


func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("myDB.db"), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}

	// fmt.Println("Connected to database successfully")

	// Migrate the schema
	db.AutoMigrate(&model.Product{})
}

func CreateNew(code string, price uint) string {
	result := db.Create(&model.Product{Code: code, Price: price})

	if result.Error != nil {
		return "Failed to create product: " + code
	}

	return "Product is created successfully with Code: " + code + " and Price: " + fmt.Sprint(price)
}

func GetAll() ([]model.Product, string) {
	var product_list []model.Product

	result := db.Find(&product_list)

	if result.Error != nil {
		return nil, "Failed to get product list"
	} else if len(product_list) == 0 {
		return nil, "No products found"
	}

	return product_list, ""
}

func GetById(id uint) string {
	result := db.First(&product, id) // find product with integer primary key

	if result.Error != nil {
		return "Product with id " + fmt.Sprint(id) + " not found"
	}

	return "Product: " + product.Code + " with price: " + fmt.Sprint(product.Price)
}

func GetByCode(code string) string {
	result := db.First(&product, "code = ?", code) // find product with code

	if result.Error != nil {
		return "Product: " + code + " not found"
	}

	return "Product: " + product.Code + " with price: " + fmt.Sprint(product.Price)
}

func UpdateByID(id, price uint, code string) string {
	result := db.Model(&product).Where("id = ?", id).Updates(map[string]any{"Price": price, "Code": code})

	if result.Error != nil {
		return "Failed to update product: " + product.Code
	}

	return "Product with id " + fmt.Sprint(id) + " is updated successfully"
}

func DeleteById(id uint) string {
	result := db.Delete(&product, id)

	if result.Error != nil {
		return "Failed to delete product: " + product.Code
	}

	return "Product with id " + fmt.Sprint(id) + " is deleted successfully"
}