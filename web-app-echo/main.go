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

	// query parameter
	// http://localhost:8080/show?name=rio%20jp&age=22
	e.GET("/show", showQueryParam)

	// form aplication/x-www-form-urlencoded
	//  run: curl -d "name=Joe Smith" -d "email=joe@yhotie.com" http://localhost:8080/save
	e.POST("/save", save)

	// handling request
	// test: curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"john@example.com"}' http://localhost:1323/user/example
	type User struct {
		Name  string `json:"name" xml:"name" form:"name" query:"name"`
		Email string `json:"email" xml:"email" form:"email" query:"email"`
	}
	e.POST("/user/example", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}

func save(c echo.Context) error {
	// get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}

func showQueryParam(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	return c.String(http.StatusOK, "name: "+name+", age: "+age)
}

func handlerName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}
