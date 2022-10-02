package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
}

const helloPath = "/hello"

func NewHelloHandler() *Hello {
	return &Hello{}
}

func (v *Hello) GetHelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, I'm Truman Capote"})
}

func (v *Hello) AddHandlers(e *echo.Echo) {
	e.GET(helloPath, v.GetHelloHandler)
}
