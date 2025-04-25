// utils/api_error.go
package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIError godoc
// @Description Represents a standard API error response
// @Property code int "The error code"
// @Property message string "A brief message explaining the error"
// @Property detail string "Detailed explanation of the error"
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Code:%d, Message: %s, Detail: %s", e.Code, e.Message, e.Detail)
}

func NewNotFoundError(message string) *APIError {
	return &APIError{
		Code:    http.StatusNotFound,
		Message: message,
		Detail:  "Resource not found",
	}
}

func NewBadRequestError(message string) *APIError {
	return &APIError{
		Code:    http.StatusBadRequest,
		Message: message,
		Detail:  "Invalid Request Data",
	}
}

func NewInternalError(message string) *APIError {
	return &APIError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Detail:  "Internal Server Error",
	}
}

func NewUnauthorizedError(message string) *APIError {
	return &APIError{
		Code:    http.StatusUnauthorized,
		Message: message,
		Detail:  "Unauthorized Access",
	}
}

func NewForbiddenError(message string) *APIError {
	return &APIError{
		Code:    http.StatusForbidden,
		Message: message,
		Detail:  "Forbidden Access",
	}
}

func HandlerError(c echo.Context, err *APIError) error {
	return c.JSON(err.Code, err)
}
