package main

import (
	"log"
	"ordent/configs"
	"ordent/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

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

	e.Logger.Fatal(e.Start(":" + port))
}
