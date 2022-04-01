package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ORMdemo() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate schema
	// db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find by primary key
	db.First(&product, "code = ?", "D42") // find record "Code == D42"

	// Update
	db.Model(&product).Update("Price", 200)
	// Update - multiple words once
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete
	db.Delete(&product, 1)
}
