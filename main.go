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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.InitDB()

	port := os.Getenv("PORT")

	e := echo.New()

	routes.UserRoutes(e)
	routes.ItemRoutes(e)
	routes.TransactionRoutes(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
