package routes

import (
	"haircut_api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	e.GET("/", handlers.Welcome)
	e.GET("/api/haircuts", handlers.GetHaircuts)
	e.GET("/api/haircuts/:id", handlers.GetHaircut)
	e.POST("/api/haircuts", handlers.CreteHairCut)
}
