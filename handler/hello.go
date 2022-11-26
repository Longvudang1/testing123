package handler

import (
	"net/http"

	"github.com/Longvudang1/testing123/models"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	x := &models.X{
		Text: "Hello from echo",
	}
	return c.JSON(http.StatusOK, x)
}
