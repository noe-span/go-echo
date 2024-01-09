package main

import (
	"encoding/json"
	"fmt"
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

	//http://localhost:3000/hello/eon
	e.GET("/hello/:name", GreetingsWithParams)

	//http://localhost:3000/hello-queries?name=eon
	e.GET("/hello-queries", GreetingsWithQuery)

	//http://localhost:3000/hello-post
	e.POST("/hello-post", GreetingsWithPost)

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

func GreetingsWithPost(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

// Encode a data structure into JSON
func Marshal(h HelloWorld) string {

	//[]byte
	jsonData, err := json.Marshal(h)
	if err != nil {
		fmt.Println("Error:", err)

	}

	return string(jsonData)

}

// Decoding json string into a go data structure
func Unmarshal(s string) HelloWorld {

	var h HelloWorld

	err := json.Unmarshal([]byte(s), &h)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return h

}
