package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Basemodel struct {
	ID        uuid.UUID       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}
