package routes

import (
	"ordent/configs"
	"ordent/controllers"
	"ordent/middlewares"
	"ordent/repositories"

	"github.com/labstack/echo/v4"
)

// TransactionRoutes sets up the transaction-related routes in the Echo framework
func TransactionRoutes(e *echo.Echo) {

	// Create an instance of the transaction repository, using the database connection from the configs package
	transactionRepo := repositories.NewTransactionRepository(configs.DB)

	// Create an instance of the item repository, used by the transaction controller
	itemRepo := repositories.NewItemRepository(configs.DB)

	// Create an instance of the transaction detail repository, used by the transaction controller
	transactionDetailRepo := repositories.NewTransactionDetailRepository(configs.DB)

	// Create an instance of the transaction controller, passing the required repositories to it
	transactionController := controllers.NewTransactionController(itemRepo, transactionRepo, transactionDetailRepo)

	e.POST("/api/v1/transactions", transactionController.CreateTransaction, middlewares.JWTAuth, middlewares.ClientAuthz)
}
