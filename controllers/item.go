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

func (ic *ItemController) GetAllItems(c echo.Context) error {
	items, err := ic.itemRepo.GetAllItems()
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to fetch items"))
	}

	return c.JSON(http.StatusOK, items)
}

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
