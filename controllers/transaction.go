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
	userPayload := c.Get("userPayload").(*dto.JWTPayload)

	var transactionBody dto.TransactionRequestBody
	if err := c.Bind(&transactionBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if transactionBody.PaidAmount < 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Paid amount must be greater than or equal to 0"))
	}

	if len(transactionBody.TransactionDetailRequestBody) == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Transaction detail is required"))
	}

	var totalRequiredPrice float64

	for _, detail := range transactionBody.TransactionDetailRequestBody {
		if detail.ItemID == "" {
			return utils.HandlerError(c, utils.NewBadRequestError("Item ID is required"))
		}

		parsedItemID, err := uuid.Parse(detail.ItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewBadRequestError("Invalid Item ID format"))
		}

		if detail.Quantity == 0 {
			return utils.HandlerError(c, utils.NewBadRequestError("Quantity must be greater than 0"))
		}

		item, err := tc.itemRepo.GetItemByID(parsedItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewNotFoundError("Item not found"))
		}

		if item.Stock < detail.Quantity {
			return utils.HandlerError(c, utils.NewBadRequestError("Insufficient stock"))
		}

		totalRequiredPrice += item.Price * float64(detail.Quantity)
	}

	if transactionBody.PaidAmount != totalRequiredPrice {
		return utils.HandlerError(c, utils.NewBadRequestError("Paid amount does not match total price"))
	}

	transaction := &models.Transaction{
		UserID:        userPayload.UserID,
		TotalPrice:    totalRequiredPrice,
		IsSuccessPaid: true,
	}

	transactionID, err := tc.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to create transaction"))
	}

	parsedTransactionID, err := uuid.Parse(transactionID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to parse transaction ID"))
	}

	for _, detail := range transactionBody.TransactionDetailRequestBody {
		parsedItemID, err := uuid.Parse(detail.ItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewBadRequestError("Invalid Item ID format"))
		}

		item, err := tc.itemRepo.GetItemByID(parsedItemID)
		if err != nil {
			return utils.HandlerError(c, utils.NewNotFoundError("Item not found"))
		}

		transactionDetail := &models.TransactionDetail{
			TransactionID: parsedTransactionID,
			ItemID:        item.ID,
			Quantity:      detail.Quantity,
			PricePerUnit:  item.Price,
			TotalPrice:    item.Price * float64(detail.Quantity),
		}

		if err := tc.transactionDetailRepo.CreateTransactionDetail(transactionDetail); err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to create transaction detail"))
		}

		item.Stock -= detail.Quantity
		if err := tc.itemRepo.EditItem(item, parsedItemID); err != nil {
			return utils.HandlerError(c, utils.NewInternalError("Failed to update item stock"))
		}
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Transaction created successfully"})
}
