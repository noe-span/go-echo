package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("TestMain Method")
// 	fmt.Println("init...")
// }

func TestSetup1(t *testing.T) {

	//setup 1
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Greetings(c)) {
		// your assertions about the response etc
		fmt.Println("Test PASSED")
		assert.Equal(t, http.StatusOK, rec.Code)

		fmt.Println(rec.Body.String())

	} else {
		fmt.Println("Test FAILED")
	}

}

func TestSetup2(t *testing.T) {

	//setup 2
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("name")
	c.SetParamValues("eon")

	if assert.NoError(t, GreetingsWithParams(c)) {
		// your assertions about the response etc
		fmt.Println("Test PASSED")
		assert.Equal(t, http.StatusOK, rec.Code)

		fmt.Println(rec.Body.String())

	} else {
		fmt.Println("Test FAILED")
	}

}

func TestSetup3(t *testing.T) {

	//setup 3
	e := echo.New()
	query := "?name=eon"
	req := httptest.NewRequest(http.MethodGet, "/hello-queries"+query, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GreetingsWithQuery(c)) {
		// your assertions about the response etc
		fmt.Println("Test PASSED")
		assert.Equal(t, http.StatusOK, rec.Code)

		fmt.Println(rec.Body.String())

	} else {
		fmt.Println("Test FAILED")
	}

}

func TestSetup3QueryParams(t *testing.T) {

	//setup 3 (Query Params)

	q := make(url.Values)
	q.Set("name", "eon")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello-queries?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GreetingsWithQuery(c)) {
		// your assertions about the response etc
		fmt.Println("Test PASSED")
		assert.Equal(t, http.StatusOK, rec.Code)

		fmt.Println(rec.Body.String())

	} else {
		fmt.Println("Test FAILED")
	}

}
