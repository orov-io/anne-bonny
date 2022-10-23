package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/orov-io/anne-bonny/video-streamer/model"
	m "github.com/orov-io/maryread/middleware"
	"github.com/stretchr/testify/assert"
)

const (
	testFactoryMiddlewareSQLDriver    = "postgres"
	testFactoryMiddlewareMockEndpoint = "/factoryMiddleware"
)

func TestInjectFactory(t *testing.T) {
	db, _, _ := sqlmock.New()
	e := echo.New()
	e.Use(m.NewSQLX().WithConfig(m.SQLXConfig{
		DB:     db,
		Driver: testFactoryMiddlewareSQLDriver,
	}))
	assert.NotPanics(t, func() {
		e.Use(InjectFactory())
	})

	e.GET(testFactoryMiddlewareMockEndpoint, testFactoryMiddlewareHandler(t, "Inject Factory"))

	req := httptest.NewRequest(http.MethodGet, testFactoryMiddlewareMockEndpoint, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
}

func testFactoryMiddlewareHandler(t *testing.T, test string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var factory *model.Factory
		assert.NotPanics(t, func() {
			factory = MustGetFactory(c)
		})

		assert.NotNil(t, factory)
		return c.NoContent(http.StatusNoContent)
	}
}
