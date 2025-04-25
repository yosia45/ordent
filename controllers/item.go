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

type ItemController struct {
	itemRepo repositories.ItemRepository
}

func NewItemController(itemRepo repositories.ItemRepository) *ItemController {
	return &ItemController{
		itemRepo: itemRepo,
	}
}

// CreateItem godoc
// @Summary Create new item
// @Description Create a new item. This endpoint can only be accessed by admin users (isAdmin=true).
// @Tags item
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param item body dto.ItemRequestBody true "Item details"
// @Success 201 {object} models.Item
// @Failure 400 {object} utils.APIError "Bad Request"
// @Failure 401 {object} utils.APIError "Unauthorized"
// @Failure 403 {object} utils.APIError "Forbidden"
// @Failure 500 {object} utils.APIError "Internal Server Error"
// @Router /api/v1/items [post]
func (ic *ItemController) CreateItem(c echo.Context) error {
	var itemBody dto.ItemRequestBody

	if err := c.Bind(&itemBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if itemBody.Name == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Name is required"))
	}

	if itemBody.Price == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Price is required"))
	}

	if itemBody.Stock == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Quantity is required"))
	}

	newItem := &models.Item{
		Name:  itemBody.Name,
		Price: itemBody.Price,
		Stock: itemBody.Stock,
	}

	if err := ic.itemRepo.CreateItem(newItem); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to create item"))
	}

	return c.JSON(http.StatusCreated, newItem)
}

// GetAllItems godoc
// @Summary Get all items
// @Description Get a list of all items. No authentication required.
// @Tags item
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Item
// @Failure 500 {object} utils.APIError "Internal Server Error"
// @Router /api/v1/items [get]
func (ic *ItemController) GetAllItems(c echo.Context) error {
	items, err := ic.itemRepo.GetAllItems()
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to fetch items"))
	}

	return c.JSON(http.StatusOK, items)
}

// EditItem godoc
// @Summary Edit an existing item
// @Description Edit an existing item. This endpoint can only be accessed by admin users (isAdmin=true).
// @Tags item
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Item ID"
// @Param item body dto.ItemRequestBody true "Item details"
// @Success 200 {object} models.Item
// @Failure 400 {object} utils.APIError "Bad Request"
// @Failure 401 {object} utils.APIError "Unauthorized"
// @Failure 403 {object} utils.APIError "Forbidden"
// @Failure 404 {object} utils.APIError "Not Found"
// @Failure 500 {object} utils.APIError "Internal Server Error"
// @Router /api/v1/items/{id} [put]
func (ic *ItemController) EditItem(c echo.Context) error {
	itemID := c.Param("id")

	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid item ID"))
	}

	var itemBody dto.ItemRequestBody
	if err := c.Bind(&itemBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if itemBody.Name == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Name is required"))
	}

	if itemBody.Price == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Price is required"))
	}

	if itemBody.Stock == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Quantity is required"))
	}

	item := &models.Item{
		Name:  itemBody.Name,
		Price: itemBody.Price,
		Stock: itemBody.Stock,
	}

	if err := ic.itemRepo.EditItem(item, parsedItemID); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to update item"))
	}

	return c.JSON(http.StatusOK, item)
}

// DeleteItem godoc
// @Summary Delete an existing item
// @Description Delete an existing item. This endpoint can only be accessed by admin users (isAdmin=true).
// @Tags item
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Item ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.APIError "Bad Request"
// @Failure 401 {object} utils.APIError "Unauthorized"
// @Failure 403 {object} utils.APIError "Forbidden"
// @Failure 404 {object} utils.APIError "Not Found"
// @Failure 500 {object} utils.APIError "Internal Server Error"
// @Router /api/v1/items/{id} [delete]
func (ic *ItemController) DeleteItem(c echo.Context) error {
	itemID := c.Param("id")

	parsedItemID, err := uuid.Parse(itemID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid item ID"))
	}

	if err := ic.itemRepo.DeleteItem(parsedItemID); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to delete item"))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Item success deleted",
	})
}
