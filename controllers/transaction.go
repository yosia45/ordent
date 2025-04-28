package controllers

import (
	"net/http"
	"ordent/dto"
	"ordent/models"
	"ordent/repositories"
	"ordent/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	itemRepo              repositories.ItemRepository
	transactionRepo       repositories.TransactionRepository
	transactionDetailRepo repositories.TransactionDetailRepository
}

// Creates a new TransactionController instance and initializes the required repositories
func NewTransactionController(itemRepo repositories.ItemRepository, transactionRepo repositories.TransactionRepository, transactionDetailRepo repositories.TransactionDetailRepository) *TransactionController {
	return &TransactionController{
		itemRepo:              itemRepo,
		transactionRepo:       transactionRepo,
		transactionDetailRepo: transactionDetailRepo,
	}
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction. This endpoint can only be accessed by users with isAdmin=false.
// @Tags transaction
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param transaction body dto.TransactionRequestBody true "Transaction details"
// @Success 201 {object} map[string]string "Transaction created successfully"
// @Failure 400 {object} utils.APIError "Bad Request"
// @Failure 401 {object} utils.APIError "Unauthorized"
// @Failure 403 {object} utils.APIError "Forbidden"
// @Failure 500 {object} utils.APIError "Internal Server Error"
// @Router /api/v1/transactions [post]
func (tc *TransactionController) CreateTransaction(c echo.Context) error {
	// Retrieve JWT payload from the request context to get the authenticated user's data
	userPayload := c.Get("userPayload").(*dto.JWTPayload)

	// Bind the request body to a struct for transaction data
	var transactionBody dto.TransactionRequestBody
	if err := c.Bind(&transactionBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	// Ensure that the paid amount is greater than or equal to 0
	if transactionBody.PaidAmount < 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Paid amount must be greater than or equal to 0"))
	}

	// Ensure that transaction details are provided (at least one item should be included)
	if len(transactionBody.TransactionDetailRequestBody) == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Transaction detail is required"))
	}

	var totalRequiredPrice float64

	// Iterate through each transaction detail to validate and calculate the price
	for _, detail := range transactionBody.TransactionDetailRequestBody {
		// Ensure that ItemID is present in the transaction detail
		if detail.ItemID == "" {
			return utils.HandlerError(c, utils.NewBadRequestError("Item ID is required"))
		}

		// Parse the ItemID into a UUID
		parsedItemID, err := uuid.Parse(detail.ItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewBadRequestError("Invalid Item ID format"))
		}

		// Ensure that the quantity is greater than 0
		if detail.Quantity == 0 {
			return utils.HandlerError(c, utils.NewBadRequestError("Quantity must be greater than 0"))
		}

		// Retrieve the item from the database using the ItemID
		item, err := tc.itemRepo.GetItemByID(parsedItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewNotFoundError("Item not found"))
		}

		// Ensure that the item stock is sufficient for the transaction
		if item.Stock < detail.Quantity {
			return utils.HandlerError(c, utils.NewBadRequestError("Insufficient stock"))
		}

		// Calculate the total price based on the price per unit and the quantity
		totalRequiredPrice += item.Price * float64(detail.Quantity)
	}

	// Ensure that the paid amount matches the total required price
	if transactionBody.PaidAmount != totalRequiredPrice {
		return utils.HandlerError(c, utils.NewBadRequestError("Paid amount does not match total price"))
	}

	// Create a new transaction with the calculated total price
	transaction := &models.Transaction{
		UserID:        userPayload.UserID,
		TotalPrice:    totalRequiredPrice,
		IsSuccessPaid: true,
	}

	// Save the transaction into the database
	transactionID, err := tc.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to create transaction"))
	}

	// Parse the newly created transaction ID into a UUID
	parsedTransactionID, err := uuid.Parse(transactionID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to parse transaction ID"))
	}

	// Iterate through each transaction detail to save it
	for _, detail := range transactionBody.TransactionDetailRequestBody {
		// Parse the ItemID into a UUID
		parsedItemID, err := uuid.Parse(detail.ItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewBadRequestError("Invalid Item ID format"))
		}

		// Retrieve the item from the database using the ItemID
		item, err := tc.itemRepo.GetItemByID(parsedItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewNotFoundError("Item not found"))
		}

		// Create a new transaction detail based on the item and quantity
		transactionDetail := &models.TransactionDetail{
			TransactionID: parsedTransactionID,
			ItemID:        item.ID,
			Quantity:      detail.Quantity,
			PricePerUnit:  item.Price,
			TotalPrice:    item.Price * float64(detail.Quantity),
		}

		// Save the transaction detail into the database
		if err := tc.transactionDetailRepo.CreateTransactionDetail(transactionDetail); err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to create transaction detail"))
		}

		// Decrease the item stock after a successful transaction
		item.Stock -= detail.Quantity
		if err := tc.itemRepo.EditItem(item, parsedItemID); err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to update item stock"))
		}
	}

	// Return a success response with HTTP status Created
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Transaction created successfully"})
}
