package configs

import (
	"fmt"
	"log"
	"ordent/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare a global variable DB to hold the database connection
var DB *gorm.DB

func InitDB() {
	var err error

	// Read database configuration from environment variables
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Create the Data Source Name (DSN) string for connecting to MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, password, host, port, name)

	// Open a connection to the database using GORM with MySQL driver
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	// Automatically migrate the schema, creating or updating tables based on models
	DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Transaction{},
		&models.TransactionDetail{},
	)

	// Log success message if database connection and migration are successful
	log.Println("Success connecting to DB")
}
