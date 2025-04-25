package routes

import (
	"ordent/configs"
	"ordent/controllers"
	"ordent/middlewares"
	"ordent/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Echo) {
	transactionRepo := repositories.NewTransactionRepository(configs.DB)
	itemRepo := repositories.NewItemRepository(configs.DB)
	transactionDetailRepo := repositories.NewTransactionDetailRepository(configs.DB)

	transactionController := controllers.NewTransactionController(itemRepo, transactionRepo, transactionDetailRepo)

	e.POST("/api/v1/transactions", transactionController.CreateTransaction, middlewares.JWTAuth, middlewares.ClientAuthz)
}
