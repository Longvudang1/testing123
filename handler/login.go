package handler

import (
	"net/http"

	"github.com/Longvudang1/testing123/models"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: "123456",
	})
}
