package controllers

import (
	"net/http"
	"test-echo/models"

	"github.com/labstack/echo/v4"
)

func FecthAllUsers(c echo.Context) error {
	result, err := models.FecthAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateUsers(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	telp := c.FormValue("telp")

	result, err := models.CreateUsers(name, address, telp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
