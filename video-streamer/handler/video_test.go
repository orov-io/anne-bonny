package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const testContentTypeHeaderKey = "Content-Type"
const testContentTypeHeaderMP4 = "video/mp4"
const testStorageServiceHostEnvKey = "STORAGE_SERVICE_HOST"

func TestNewVideo(t *testing.T) {
	storageServiceURL, err := getStorageServiceHost()
	video := NewVideoHandler(storageServiceURL)
	assert.NoError(t, err)
	assert.NotNil(t, video)
}

func TestGetVideoHandler(t *testing.T) {
	e, err := getEchoRouterWithVideoHandlers()
	if err != nil {
		t.Errorf("unable to intialize video handlers due to: %v", err)
		return
	}

	req := httptest.NewRequest(http.MethodGet, videoPath, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, testContentTypeHeaderMP4, res.Header.Get(testContentTypeHeaderKey))
}

func getEchoRouterWithVideoHandlers() (*echo.Echo, error) {
	e := echo.New()
	storageServiceURL, err := getStorageServiceHost()
	video := NewVideoHandler(storageServiceURL)
	video.AddHandlers(e)
	return e, err
}

func getStorageServiceHost() (*url.URL, error) {
	storageServiceHostRaw := os.Getenv(testStorageServiceHostEnvKey)
	return url.Parse(storageServiceHostRaw)
}
