package main

import (
	"haircut-api-go-gin-gorm-mysql/config"
	"haircut-api-go-gin-gorm-mysql/controllers"
	"haircut-api-go-gin-gorm-mysql/repositories"
	"haircut-api-go-gin-gorm-mysql/routes"
	"haircut-api-go-gin-gorm-mysql/services"
)

func main() {
	// config inisialisasi database koneksi
	config.ConnectDatabase()

	// inisialisasi repository, service, controller
	haircutRepo := repositories.NewHaircutRepository(config.DB)
	haircutService := services.NewHaircutService(haircutRepo)
	haircutController := controllers.NewHaircutController(haircutService)

	// setup router
	router := routes.SetupRouter(haircutController)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
