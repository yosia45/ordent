package repositories

import (
	"ordent/dto"
	"ordent/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ItemRepository interface defines the methods for interacting with Item data
type ItemRepository interface {
	CreateItem(item *models.Item) error
	GetAllItems() ([]dto.GetAllItemResponse, error)
	GetItemByID(itemID uuid.UUID) (*models.Item, error)
	EditItem(item *models.Item, itemID uuid.UUID) error
	DeleteItem(itemID uuid.UUID) error
}

// itemRepository struct is the concrete implementation of ItemRepository
type itemRepository struct {
	db *gorm.DB // holds a reference to the database connection
}

// NewItemRepository is a constructor function that returns a new instance of itemRepository
// It initializes the repository with a database connection
func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db} // Return the repository instance with the db reference
}

// CreateItem inserts a new item into the database
func (ir *itemRepository) CreateItem(item *models.Item) error {
	// Try to insert the item into the database
	if err := ir.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}

// GetItemByID fetches an item from the database by its unique ID
func (ir *itemRepository) GetItemByID(itemID uuid.UUID) (*models.Item, error) {
	var item models.Item

	// Look for an item with the given itemID
	if err := ir.db.Where("id = ?", itemID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// GetAllItems fetches all items from the database and returns them in a specific DTO format
func (ir *itemRepository) GetAllItems() ([]dto.GetAllItemResponse, error) {
	var items []models.Item

	// Fetch all items from the database
	if err := ir.db.Find(&items).Error; err != nil {
		return nil, err
	}

	// Loop through all fetched items and prepare the response DTO
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

// EditItem updates an existing item in the database with new data
func (ir *itemRepository) EditItem(item *models.Item, itemID uuid.UUID) error {
	// Update the item with the new data where the ID matches the itemID
	if err := ir.db.Model(&models.Item{}).Where("id = ?", itemID).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

// DeleteItem deletes an item from the database by its ID
func (ir *itemRepository) DeleteItem(itemID uuid.UUID) error {
	// Delete the item where the ID matches the given itemID
	if err := ir.db.Delete(&models.Item{}, itemID).Error; err != nil {
		return err
	}
	return nil
}
