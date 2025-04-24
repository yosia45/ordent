package repositories

import (
	"ordent/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *dto.RegisterBodyRequest) error
	GetUserByEmail(email string) (*dto.GetUserByEmailResponse, error)
	GetUserDetail(userID uuid.UUID) (*dto.GetUserDetailResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) CreateUser(user *dto.RegisterBodyRequest) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(email string) (*dto.GetUserByEmailResponse, error) {
	var user dto.GetUserByEmailResponse
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserDetail(userID uuid.UUID) (*dto.GetUserDetailResponse, error) {
	var user dto.GetUserDetailResponse
	if err := ur.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
