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
// APIError is a custom error type that represents an API error response.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// Error implements the `Error()` method for the `APIError` struct,
// providing a string representation of the error.
func (e *APIError) Error() string {
	// Return the error in a formatted string that includes the code, message, and detail
	return fmt.Sprintf("Code:%d, Message: %s, Detail: %s", e.Code, e.Message, e.Detail)
}

// NewNotFoundError creates a new APIError with a 404 Not Found status code.
func NewNotFoundError(message string) *APIError {
	return &APIError{
		Code:    http.StatusNotFound,
		Message: message,
		Detail:  "Resource not found",
	}
}

// NewBadRequestError creates a new APIError with a 400 Bad Request status code.
func NewBadRequestError(message string) *APIError {
	return &APIError{
		Code:    http.StatusBadRequest,
		Message: message,
		Detail:  "Invalid Request Data",
	}
}

// NewInternalError creates a new APIError with a 500 Internal Server Error status code.
func NewInternalError(message string) *APIError {
	return &APIError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Detail:  "Internal Server Error",
	}
}

// NewUnauthorizedError creates a new APIError with a 401 Unauthorized status code.
func NewUnauthorizedError(message string) *APIError {
	return &APIError{
		Code:    http.StatusUnauthorized,
		Message: message,
		Detail:  "Unauthorized Access",
	}
}

// NewForbiddenError creates a new APIError with a 403 Forbidden status code.
func NewForbiddenError(message string) *APIError {
	return &APIError{
		Code:    http.StatusForbidden,
		Message: message,
		Detail:  "Forbidden Access",
	}
}

// HandlerError is a helper function that returns the APIError as a JSON response.
func HandlerError(c echo.Context, err *APIError) error {
	// Return the APIError as a JSON response with the error code and message
	return c.JSON(err.Code, err)
}
