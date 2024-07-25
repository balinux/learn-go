package handlers

import (
	"haircut_api/models"
	"haircut_api/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetHaircuts(c echo.Context) error {
	return c.JSON(http.StatusOK, services.GetAllHaircut())
}

// handle untuk mengakses data api haricut berdasarkan id
func GetHaircut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	haircut := services.GetHaircutById(id)
	if haircut == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Haircut not found"})
	}
	return c.JSON(http.StatusOK, haircut)
}

// create haircut
func CreteHairCut(c echo.Context) error {
	newHaircut := new(models.Haircut)
	if err := c.Bind(newHaircut); err != nil {
		return err
	}
	services.AddHaircut(*newHaircut)
	return c.JSON(http.StatusCreated, newHaircut)
}
