package main

import (
	"haircut_api/databases"
	"haircut_api/middleware"
	"haircut_api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	databases.InitialDatabase()

	e := echo.New()

	// register middleware
	middleware.RegisterMiddleware(e)

	// register routes
	routes.RegisterRouter(e)

	// start server
	e.Logger.Fatal(e.Start(":8889"))
}
