package routes

import (
	"ordent/configs"
	"ordent/controllers"
	"ordent/middlewares"
	"ordent/repositories"

	"github.com/labstack/echo/v4"
)

func ItemRoutes(e *echo.Echo) {
	itemRepo := repositories.NewItemRepository(configs.DB)

	itemController := controllers.NewItemController(itemRepo)

	e.POST("/api/v1/items", itemController.CreateItem, middlewares.JWTAuth, middlewares.AdminAuthz)
	e.GET("/api/v1/items", itemController.GetAllItems)
	e.PUT("/api/v1/items/:id", itemController.EditItem, middlewares.JWTAuth, middlewares.AdminAuthz)
	e.DELETE("/api/v1/items/:id", itemController.DeleteItem, middlewares.JWTAuth, middlewares.AdminAuthz)
}
