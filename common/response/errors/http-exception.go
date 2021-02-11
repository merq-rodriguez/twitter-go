package errors

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type HttpException struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error, omitempty"`
}

func NotFoundError(c echo.Context, message string, payload interface{}) error {
	var _message string
	status := http.StatusNotFound
	if strings.TrimSpace(message) == "" {
		_message = http.StatusText(status)
	} else {
		_message = message
	}
	return c.JSON(status, HttpException{
		Message:    _message,
		StatusCode: status,
		Error:      payload,
	})
}

func BadRequestError(c echo.Context, message string, payload interface{}) error {
	var _message string
	status := http.StatusBadRequest
	if strings.TrimSpace(message) == "" {
		_message = http.StatusText(status)
	} else {
		_message = message
	}
	return c.JSON(status, HttpException{
		Message:    _message,
		StatusCode: status,
		Error:      payload,
	})
}

func NotAcceptableError(c echo.Context, message string, payload interface{}) error {
	var _message string
	status := http.StatusNotAcceptable
	if strings.TrimSpace(message) == "" {
		_message = http.StatusText(status)
	} else {
		_message = message
	}
	return c.JSON(status, HttpException{
		Message:    _message,
		StatusCode: status,
		Error:      payload,
	})
}

func UnauthorizedError(c echo.Context, message string, payload interface{}) error {
	var _message string
	status := http.StatusUnauthorized
	if strings.TrimSpace(message) == "" {
		_message = http.StatusText(status)
	} else {
		_message = message
	}
	return c.JSON(status, HttpException{
		Message:    _message,
		StatusCode: status,
		Error:      payload,
	})
}
func InternalServerError(c echo.Context, message string, payload interface{}) error {
	var _message string
	status := http.StatusInternalServerError
	if strings.TrimSpace(message) == "" {
		_message = http.StatusText(status)
	} else {
		_message = message
	}
	return c.JSON(status, HttpException{
		Message:    _message,
		StatusCode: status,
		Error:      payload,
	})
}
