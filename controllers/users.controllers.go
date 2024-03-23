package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	// name := c.FormValue("name")
	// address := c.FormValue("address")
	// telp := c.FormValue("telp")

	payload := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&payload)

	if err != nil {
		return err
	}

	name := payload["name"].(string)
	address := payload["address"].(string)
	telp := payload["telp"].(string)

	result, err := models.CreateUsers(name, address, telp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUsers(c echo.Context) error {
	id := c.Param("id")
	// name := c.FormValue("name")
	// address := c.FormValue("address")
	// telp := c.FormValue("telp")

	payload := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&payload)

	if err != nil {
		return err
	}

	name := payload["name"].(string)
	address := payload["address"].(string)
	telp := payload["telp"].(string)

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUsers(conv_id, name, address, telp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteUsers(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUsers(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DetailUsers(c echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DetailUsers(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}
