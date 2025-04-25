package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterBodyRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginBodyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserByEmailResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	IsAdmin  bool      `json:"is_admin"`
	Password string    `json:"password"`
}

type GetUserDetailResponse struct {
	ID           uuid.UUID             `json:"id"`
	FullName     string                `json:"full_name"`
	Email        string                `json:"email"`
	Username     string                `json:"username"`
	IsAdmin      bool                  `json:"is_admin"`
	CreatedAt    time.Time             `json:"created_at"`
	Transactions []TransactionResponse `json:"transactions"`
}
