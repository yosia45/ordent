package routes

import (
	"ordent/configs"
	"ordent/controllers"
	"ordent/middlewares"
	"ordent/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository(configs.DB)

	userController := controllers.NewUserController(userRepo)

	e.GET("/api/v1/myprofiles", userController.MyProfile, middlewares.JWTAuth, middlewares.ClientAuthz)
	e.POST("/api/v1/register", userController.RegisterUser)
	e.POST("/api/v1/login", userController.LoginUser)
}
