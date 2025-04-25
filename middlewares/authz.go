package middlewares

import (
	"net/http"
	"ordent/dto"

	"github.com/labstack/echo/v4"
)

func AdminAuthz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userPayload := c.Get("userPayload").(*dto.JWTPayload)

		if userPayload.IsAdmin == false {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "You are not authorized to access this resource"})
		}

		return next(c)
	}
}

func ClientAuthz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userPayload := c.Get("userPayload").(*dto.JWTPayload)

		if userPayload.IsAdmin == true {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "You are not authorized to access this resource"})
		}

		return next(c)
	}
}
