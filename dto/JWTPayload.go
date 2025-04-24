package dto

import "github.com/google/uuid"

type JWTPayload struct {
	UserID  uuid.UUID `json:"user_id"`
	IsAdmin bool      `json:"is_admin"`
}
