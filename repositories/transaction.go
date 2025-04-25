package repositories

import (
	"ordent/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) (transactionID string, err error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (tr *transactionRepository) CreateTransaction(transaction *models.Transaction) (transactionID string, err error) {
	if err := tr.db.Create(transaction).Error; err != nil {
		return "", err
	}

	return transaction.ID.String(), nil
}
