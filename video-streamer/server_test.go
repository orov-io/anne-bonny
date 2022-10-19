package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/orov-io/maryread"
	"github.com/stretchr/testify/assert"
)

func TestParsePort(t *testing.T) {
	expectedPort := os.Getenv(portEnvKey)
	assert.NotEmpty(t, port)
	assert.Equal(t, ":"+expectedPort, port)
}

func TestAddHandlers(t *testing.T) {
	app := maryread.Default()
	addHandlers(app)
	expectedRoutes := []*echo.Route{
		{
			Method: http.MethodGet,
			Path:   "/ping",
			Name:   "handler.(*Ping).GetPingHandler",
		},
		{
			Method: http.MethodGet,
			Path:   "/video",
			Name:   "handler.(*Video).GetVideoHandler",
		},
	}
	routes := app.Router().Routes()
	for _, route := range expectedRoutes {
		assert.True(t, isResgisteredRoute(route, routes), fmt.Sprintf("%+v", route))
	}

}

func isResgisteredRoute(expectedRoute *echo.Route, routes []*echo.Route) bool {
	for _, route := range routes {
		if expectedRoute.Method == route.Method && expectedRoute.Path == route.Path && strings.Contains(route.Name, expectedRoute.Name) {
			return true
		}
	}
	return false
}
