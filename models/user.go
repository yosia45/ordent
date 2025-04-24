package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Basemodel
	FullName     string        `json:"full_name" gorm:"not null"`
	Email        string        `json:"email" gorm:"not null;unique"`
	Username     string        `json:"username" gorm:"not null;unique"`
	Password     string        `json:"password" gorm:"not null"`
	IsAdmin      bool          `json:"is_admin" gorm:"default:false"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(hashedPassword)
	return
}
