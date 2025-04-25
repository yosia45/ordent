package repositories

import (
	"ordent/dto"
	"ordent/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *models.Item) error
	GetAllItems() ([]dto.GetAllItemResponse, error)
	GetItemByID(itemID uuid.UUID) (*models.Item, error)
	EditItem(item *models.Item, itemID uuid.UUID) error
	DeleteItem(itemID uuid.UUID) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (ir *itemRepository) CreateItem(item *models.Item) error {
	if err := ir.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) GetItemByID(itemID uuid.UUID) (*models.Item, error) {
	var item models.Item
	if err := ir.db.Where("id = ?", itemID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (ir *itemRepository) GetAllItems() ([]dto.GetAllItemResponse, error) {
	var items []models.Item
	if err := ir.db.Find(&items).Error; err != nil {
		return nil, err
	}

	var itemResponses []dto.GetAllItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, dto.GetAllItemResponse{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return itemResponses, nil
}

func (ir *itemRepository) EditItem(item *models.Item, itemID uuid.UUID) error {
	if err := ir.db.Model(&models.Item{}).Where("id = ?", itemID).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) DeleteItem(itemID uuid.UUID) error {
	if err := ir.db.Delete(&models.Item{}, itemID).Error; err != nil {
		return err
	}
	return nil
}
