package main

import (
	"log"
	"ordent/configs"
	"ordent/routes"
	"os"

	_ "ordent/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Ordent API
// @version 1.0
// @description This is the API documentation for Ordent backend.
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	configs.InitDB()

	// Get the server port from environment variables
	port := os.Getenv("PORT")

	// Create a new Echo instance
	e := echo.New()

	// Register User, Item, and Transaction related routes
	routes.UserRoutes(e)
	routes.ItemRoutes(e)
	routes.TransactionRoutes(e)

	// Set up Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start the Echo server on the specified port
	e.Logger.Fatal(e.Start(":" + port))
}
