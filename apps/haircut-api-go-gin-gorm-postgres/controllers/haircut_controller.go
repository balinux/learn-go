package controllers

import (
	"haircut-api-go-gin-gorm-mysql/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HaircutController struct {
	service services.HaircutService
}

func NewHaircutController(service services.HaircutService) *HaircutController {
	return &HaircutController{service}
}

func (ctrl *HaircutController) GetAllHaircuts(c *gin.Context) {
	haircuts, err := ctrl.service.GetAllHaircuts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": haircuts})
}
