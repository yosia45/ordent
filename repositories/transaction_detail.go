package repositories

import (
	"ordent/models"

	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	CreateTransactionDetail(transactionDetail *models.TransactionDetail) error
}

type transactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepository{db: db}
}

func (tr *transactionDetailRepository) CreateTransactionDetail(transactionDetail *models.TransactionDetail) error {
	if err := tr.db.Create(transactionDetail).Error; err != nil {
		return err
	}
	return nil
}
