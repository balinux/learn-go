package routes

import (
	"haircut-api-go-gin-gorm-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(haircutController *controllers.HaircutController) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/haricuts", haircutController.GetAllHaircuts)
	}
	return router
}
