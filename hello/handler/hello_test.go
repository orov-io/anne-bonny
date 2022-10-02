package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHello(t *testing.T) {
	hello := NewHelloHandler()
	assert.NotNil(t, hello)
}

type getHelloResponse struct {
	Message string
}

func TestGetHelloHandler(t *testing.T) {
	e := getEchoRouterWithHelloHandlers()

	req := httptest.NewRequest(http.MethodGet, helloPath, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	var helloResponse getHelloResponse
	rawBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, rawBody)

	t.Logf("Tha raw body: %v", string(rawBody))

	err = json.Unmarshal(rawBody, &helloResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, helloResponse.Message)
}

func getEchoRouterWithHelloHandlers() *echo.Echo {
	e := echo.New()
	hello := NewHelloHandler()
	hello.AddHandlers(e)
	return e
}
