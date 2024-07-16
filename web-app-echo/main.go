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

	// path parameter
	e.GET("/users/:id", func(c echo.Context) error {
		// user id from path /users/:id
		id := c.Param("id")
		return c.String(http.StatusOK, "user id:"+id)
	})

	// user path with external handler
	e.GET("/user/:name", handlerName)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}

func handlerName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}
