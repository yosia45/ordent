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
	// Define claims (data embedded inside the token)
	claims := jwt.MapClaims{
		"user_id":  payload.UserID,
		"is_admin": payload.IsAdmin,
	}

	// Create a new JWT token using HS256 signing method and attach claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key from environment variable
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.HandlerError(c, utils.NewForbiddenError("Token not found"))
		}

		// Remove "Bearer " prefix from the token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return utils.HandlerError(c, utils.NewForbiddenError("Token not found"))
		}

		// Parse the token string
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Check that the signing method is HMAC (HS256)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, utils.NewForbiddenError("Invalid Signing Method")
			}

			// Provide the secret key to verify the signature
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		// If parsing or verification fails, reject the request
		if err != nil {
			return utils.HandlerError(c, utils.NewForbiddenError("Invalid token"))
		}

		// If the token is valid and claims are present
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// Extract user_id from claims
			userIDstr, ok := claims["user_id"].(string)
			if !ok {
				return utils.HandlerError(c, utils.NewForbiddenError("Invalid claim user_id"))
			}

			// Extract is_admin from claims
			isAdmin, ok := claims["is_admin"].(bool)
			if !ok {
				return utils.HandlerError(c, utils.NewForbiddenError("Invalid claim is_admin"))
			}

			// Parse the user_id string into UUID type
			userID, err := uuid.Parse(userIDstr)
			if err != nil {
				return utils.HandlerError(c, utils.NewBadRequestError("Invalid UUID Claim Format"))
			}

			// Set the extracted payload into the context, so the next handler can access user information
			c.Set("userPayload", &dto.JWTPayload{
				UserID:  userID,
				IsAdmin: isAdmin,
			})

			// Proceed to the next handler
			return next(c)
		}

		// If token is invalid or claims are missing, reject the request
		return utils.HandlerError(c, utils.NewForbiddenError("Invalid token"))
	}
}
