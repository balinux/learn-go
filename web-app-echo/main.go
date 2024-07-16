package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// basic routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello word with echo")
	})

	e.GET("/about", func(c echo.Context) error {
		return c.String(http.StatusOK, "ini halaman aboout")
	})

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
