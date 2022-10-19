package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const testContentTypeHeaderKey = "Content-Type"
const testContentTypeHeaderMP4 = "video/mp4"

var testStorageAccountName = os.Getenv("STORAGE_ACCOUNT_NAME")
var testStorageContainer = os.Getenv("STORAGE_CONTAINER")
var testStorageAccessKey = os.Getenv("STORAGE_ACCESS_KEY")

func TestNewVideo(t *testing.T) {
	video, err := NewVideoHandler(testStorageAccountName, testStorageAccessKey, testStorageContainer)
	assert.NoError(t, err)
	assert.NotNil(t, video)
}

func TestGetVideoHandler(t *testing.T) {
	e, err := getEchoRouterWithVideoHandlers()
	if err != nil {
		t.Errorf("unable to intialize video handler due to: %v", err)
		return
	}

	req := httptest.NewRequest(http.MethodGet, "/video/SampleVideo_1280x720_1mb.mp4", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, testContentTypeHeaderMP4, res.Header.Get(testContentTypeHeaderKey))
}

func getEchoRouterWithVideoHandlers() (*echo.Echo, error) {
	e := echo.New()
	video, err := NewVideoHandler(testStorageAccountName, testStorageAccessKey, testStorageContainer)
	video.AddHandlers(e)
	return e, err
}
