package middlewares

import (
	"ordent/dto"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(payload dto.JWTPayload) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  payload.UserID,
		"is_admin": payload.IsAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
