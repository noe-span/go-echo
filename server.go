package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	//http://localhost:3000/hello
	e.GET("/hello", Greetings)

	//localhost:3000/hello/eon
	e.GET("/hello/:name", GreetingsWithParams)

	//http://localhost:3000/hello-queries?name=eon
	e.GET("/hello-queries", GreetingsWithQuery)

	e.Logger.Fatal(e.Start(":3000"))
}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

func GreetingsWithParams(c echo.Context) error {
	params := c.Param("name")
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World, my name is " + params,
	})
}

func GreetingsWithQuery(c echo.Context) error {
	query := c.QueryParam("name")
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World i'm using queries and my name is " + query,
	})
}
