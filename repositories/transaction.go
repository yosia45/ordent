package repositories

import (
	"ordent/models"

	"gorm.io/gorm"
)

// TransactionRepository interface defines the contract for transaction-related database operations
type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) (transactionID string, err error)
}

// transactionRepository is a concrete implementation of the TransactionRepository interface
type transactionRepository struct {
	db *gorm.DB // holds a reference to the database connection
}

// NewTransactionRepository is a constructor function that returns a new TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// CreateTransaction inserts a new transaction record into the database
func (tr *transactionRepository) CreateTransaction(transaction *models.Transaction) (transactionID string, err error) {
	// Attempt to create the transaction record in the database
	if err := tr.db.Create(transaction).Error; err != nil {
		return "", err
	}

	// Return the created transaction's ID as a string and no error
	return transaction.ID.String(), nil
}
