package routes

import (
	"ordent/configs"
	"ordent/controllers"
	"ordent/middlewares"
	"ordent/repositories"

	"github.com/labstack/echo/v4"
)

// UserRoutes sets up the user-related routes in the Echo framework
func UserRoutes(e *echo.Echo) {
	// Create an instance of the user repository, using the database connection from the configs package
	userRepo := repositories.NewUserRepository(configs.DB)

	// Create an instance of the user controller, passing the user repository to it
	userController := controllers.NewUserController(userRepo)

	e.GET("/api/v1/myprofiles", userController.MyProfile, middlewares.JWTAuth, middlewares.ClientAuthz)
	e.POST("/api/v1/register", userController.RegisterUser)
	e.POST("/api/v1/login", userController.LoginUser)
}
