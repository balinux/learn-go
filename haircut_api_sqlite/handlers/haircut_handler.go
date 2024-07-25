package handlers

import (
	"haircut_api/models"
	"haircut_api/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetHaircuts(c echo.Context) error {
	haircuts, err := services.GetAllHaircut()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve haircuts"})
	}
	return c.JSON(http.StatusOK, haircuts)
}

// handle untuk mengakses data api haricut berdasarkan id
func GetHaircut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	haircut, err := services.GetHaircutByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve haircut"})
	}
	if haircut == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Haircut not found"})
	}
	return c.JSON(http.StatusOK, haircut)
}

// create haircut
func CreteHairCut(c echo.Context) error {
	newHaircut := new(models.Haircut)
	if err := c.Bind(newHaircut); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid input"})
	}
	id, err := services.AddHaircut(*newHaircut)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create haircut"})
	}
	newHaircut.ID = int(id)

	return c.JSON(http.StatusCreated, newHaircut)
}
