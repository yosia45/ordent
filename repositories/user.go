package repositories

import (
	"ordent/dto"
	"ordent/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define UserRepository interface that outlines the methods a User repository must implement
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*dto.GetUserByEmailResponse, error)
	GetUserDetail(userID uuid.UUID) (*dto.GetUserDetailResponse, error)
}

// Concrete struct that implements UserRepository interface, holds a gorm.DB instance
type userRepository struct {
	db *gorm.DB // holds a reference to the database connection
}

// Constructor function that returns a new instance of userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser inserts a new User record into the database
func (ur *userRepository) CreateUser(user *models.User) error {
	// Create the user in the database
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user by their email address
func (ur *userRepository) GetUserByEmail(email string) (*dto.GetUserByEmailResponse, error) {
	var user models.User

	// Query the database for a user with the matching email
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// Map the found user to a DTO response struct
	return &dto.GetUserByEmailResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}, nil
}

// GetUserDetail retrieves a user's detailed information including their transactions and transaction details
func (ur *userRepository) GetUserDetail(userID uuid.UUID) (*dto.GetUserDetailResponse, error) {
	var user models.User

	// Query the user by ID and preload related transactions and transaction details, including the related Item
	if err := ur.db.Preload("Transactions.TransactionDetails.Item").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	var transactions []dto.TransactionResponse

	// Loop through each transaction of the user
	for _, trx := range user.Transactions {
		var trxDetails []dto.TransactionDetailResponse

		// Loop through each transaction detail inside the transaction
		for _, detail := range trx.TransactionDetails {

			// Append each transaction detail with item info into a slice
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

		// Append the transaction with its details into the transactions slice
		transactions = append(transactions, dto.TransactionResponse{
			ID:                 trx.ID,
			TotalPrice:         trx.TotalPrice,
			IsSuccessPaid:      trx.IsSuccessPaid,
			CreatedAt:          trx.CreatedAt,
			TransactionDetails: trxDetails,
		})
	}

	// Create the final user detail response including transactions
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
