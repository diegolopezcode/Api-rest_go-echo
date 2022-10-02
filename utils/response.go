package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message",omitempty`
	Data    interface{} `json:"data",omitempty`
}

func OkResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func BadResponse(c echo.Context, data interface{}) error {
	return nil
}
