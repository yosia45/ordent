package routes

import (
	"ordent/configs"
	"ordent/controllers"
	"ordent/middlewares"
	"ordent/repositories"

	"github.com/labstack/echo/v4"
)

// ItemRoutes sets up the item-related routes in the Echo framework
func ItemRoutes(e *echo.Echo) {
	// Create an instance of the item repository, using the database connection from the configs package
	itemRepo := repositories.NewItemRepository(configs.DB)

	// Create an instance of the item controller, passing the item repository to it
	itemController := controllers.NewItemController(itemRepo)

	e.POST("/api/v1/items", itemController.CreateItem, middlewares.JWTAuth, middlewares.AdminAuthz)
	e.GET("/api/v1/items", itemController.GetAllItems)
	e.PUT("/api/v1/items/:id", itemController.EditItem, middlewares.JWTAuth, middlewares.AdminAuthz)
	e.DELETE("/api/v1/items/:id", itemController.DeleteItem, middlewares.JWTAuth, middlewares.AdminAuthz)
}
