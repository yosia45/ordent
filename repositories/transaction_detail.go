package repositories

import (
	"ordent/models"

	"gorm.io/gorm"
)

// TransactionDetailRepository interface defines the operations related to transaction details
type TransactionDetailRepository interface {
	CreateTransactionDetail(transactionDetail *models.TransactionDetail) error
}

// transactionDetailRepository is the concrete implementation of TransactionDetailRepository
type transactionDetailRepository struct {
	db *gorm.DB // holds a reference to the database connection
}

// NewTransactionDetailRepository is a constructor function that returns a new TransactionDetailRepository instance
func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepository{db: db}
}

// CreateTransactionDetail inserts a new transaction detail record into the database
func (tr *transactionDetailRepository) CreateTransactionDetail(transactionDetail *models.TransactionDetail) error {
	// Try to create (insert) the transaction detail into the database
	if err := tr.db.Create(transactionDetail).Error; err != nil {
		return err
	}
	return nil
}
