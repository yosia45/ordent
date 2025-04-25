package middlewares

import (
	"ordent/dto"
	"ordent/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GenerateJWT(payload dto.JWTPayload) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  payload.UserID,
		"is_admin": payload.IsAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.HandlerError(c, utils.NewForbiddenError("Token not found"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return utils.HandlerError(c, utils.NewForbiddenError("Token not found"))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, utils.NewForbiddenError("Invalid Signing Method")
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			return utils.HandlerError(c, utils.NewForbiddenError("Invalid token"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userIDstr, ok := claims["user_id"].(string)
			if !ok {
				return utils.HandlerError(c, utils.NewForbiddenError("Invalid claim user_id"))
			}

			isAdmin, ok := claims["is_admin"].(bool)
			if !ok {
				return utils.HandlerError(c, utils.NewForbiddenError("Invalid claim is_admin"))
			}

			userID, err := uuid.Parse(userIDstr)
			if err != nil {
				return utils.HandlerError(c, utils.NewBadRequestError("Invalid UUID Claim Format"))
			}

			c.Set("userPayload", &dto.JWTPayload{
				UserID:  userID,
				IsAdmin: isAdmin,
			})

			return next(c)
		}

		return utils.HandlerError(c, utils.NewForbiddenError("Invalid token"))
	}
}
