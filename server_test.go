package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain Method")
	fmt.Println("init...")
}

func TestSetup1(t *testing.T) {

	//setup 1
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello")
	// c.Param("eon")
	// c.QueryParam("eon")

	fmt.Println(rec.Body.String())

	Greetings(c)

	if assert.NoError(t, Greetings(c)) {
		// your assertions about the response etc
		fmt.Println("Test PASSED")
	} else {
		fmt.Println("Test FAILED")
	}

	// tr := &http.Transport{}
	// client := &http.Client{Transport: tr}

	// // Call the api
	// respApi, err := client.Get(
	// 	"http://localhost:3000/hello/eon",
	// )

	// fmt.Println("respAPI", respApi.Body)

	// fmt.Println("error", err)

}

// func testHanlder1(c echo.Context) {

// }

func TestSetup2(t *testing.T) {

	//setup 2
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello/eon", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello/eon")

	fmt.Println(rec.Body)

}

// func testHanlder2(c echo.Context) {

// }

func TestSetup3(t *testing.T) {

	//setup 3
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello-queries", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("hello-queries")
	c.SetParamNames("name")
	c.SetParamValues("eon")

	fmt.Println(rec.Body)

}

// func testHanlder3(c echo.Context) {

// 	c.Get()

// }
