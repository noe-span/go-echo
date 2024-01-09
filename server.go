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

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
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

	//http://localhost:3000/hello-user-post
	e.POST("/hello-user-post", UserPost)

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

func UserPost(c echo.Context) error {

	u := new(User)

	u.Name = "eon"
	u.Age = 99

	//[]byte
	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error:", err)

	}

	return c.JSONBlob(
		http.StatusOK,
		[]byte(jsonData),
	)
}

/*
===============================================================================
json
===============================================================================
*/

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
