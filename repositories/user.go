package repositories

import (
	"fmt"
	"ordent/dto"
	"ordent/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*dto.GetUserByEmailResponse, error)
	GetUserDetail(userID uuid.UUID) (*dto.GetUserDetailResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(email string) (*dto.GetUserByEmailResponse, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &dto.GetUserByEmailResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}, nil
}

func (ur *userRepository) GetUserDetail(userID uuid.UUID) (*dto.GetUserDetailResponse, error) {
	var user models.User
	if err := ur.db.Preload("Transactions.TransactionDetails.Item").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	var transactions []dto.TransactionResponse
	for _, trx := range user.Transactions {
		var trxDetails []dto.TransactionDetailResponse
		for _, detail := range trx.TransactionDetails {
			trxDetails = append(trxDetails, dto.TransactionDetailResponse{
				Item: dto.GetItemDetailTransactionResponse{
					ID:   detail.Item.ID,
					Name: detail.Item.Name,
				},
				Quantity:     detail.Quantity,
				PricePerUnit: detail.PricePerUnit,
				TotalPrice:   detail.TotalPrice,
			})
		}

		transactions = append(transactions, dto.TransactionResponse{
			ID:                 trx.ID,
			TotalPrice:         trx.TotalPrice,
			IsSuccessPaid:      trx.IsSuccessPaid,
			CreatedAt:          trx.CreatedAt,
			TransactionDetails: trxDetails,
		})
	}

	fmt.Println(transactions)

	response := &dto.GetUserDetailResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		Email:        user.Email,
		Username:     user.Username,
		IsAdmin:      user.IsAdmin,
		CreatedAt:    user.CreatedAt,
		Transactions: transactions,
	}

	return response, nil
}
