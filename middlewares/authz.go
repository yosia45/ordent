package middlewares

import (
	"net/http"
	"ordent/dto"

	"github.com/labstack/echo/v4"
)

func AdminAuthz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Retrieve the user payload that was set in the JWTAuth middleware
		userPayload := c.Get("userPayload").(*dto.JWTPayload)

		// Check if the user is NOT an admin
		if !userPayload.IsAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "You are not authorized to access this resource"})
		}

		// If the user is an admin, proceed to the next handler
		return next(c)
	}
}

func ClientAuthz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Retrieve the user payload that was set in the JWTAuth middleware
		userPayload := c.Get("userPayload").(*dto.JWTPayload)

		// Check if the user IS an admin
		if userPayload.IsAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "You are not authorized to access this resource"})
		}

		// If the user is a client (non-admin), proceed to the next handler
		return next(c)
	}
}
