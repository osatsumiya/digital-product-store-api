package database

import (
	"digital-product-store-api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123123 dbname=digital_product_store port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully!")

	DB = db

	DB.AutoMigrate(
		&models.Product{},
		&models.Customer{},
		&models.Order{},
		&models.License{},
	)
}
