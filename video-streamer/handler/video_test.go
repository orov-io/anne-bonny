package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	vsMiddleware "github.com/orov-io/anne-bonny/video-streamer/middleware"
	"github.com/orov-io/maryread"
	m "github.com/orov-io/maryread/middleware"
	"github.com/stretchr/testify/assert"
)

const (
	testHandlerUUID        = "f0fd7ef3-86f9-499c-9853-fac245946e0c"
	testHandlerInvalidUUID = "invalid"
	testHandlerEmptyUUID   = ""
	testHandlerPath        = "/SampleVideo_1280x720_1mb.mp4"
	testHandlerID          = 1
	testHandlerSQLDriver   = "postgres"

	testContentTypeHeaderKey     = "Content-Type"
	testContentTypeHeaderMP4     = "video/mp4"
	testStorageServiceHostEnvKey = "STORAGE_SERVICE_HOST"
)

func TestNewVideo(t *testing.T) {
	storageServiceURL, err := getStorageServiceHost()
	video := NewVideoHandler(storageServiceURL)
	assert.NoError(t, err)
	assert.NotNil(t, video)
}

func TestGetVideoHandler_HappyPath(t *testing.T) {
	db, mock := getMockDBWithHappyPathExpectations()
	e, err := getEchoRouterWithVideoHandlers(db)
	if err != nil {
		t.Errorf("unable to intialize video handlers due to: %v", err)
		return
	}

	URI, err := url.Parse(fmt.Sprintf("%s?id=%s", videoPath, testHandlerUUID))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, URI.String(), nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, testContentTypeHeaderMP4, res.Header.Get(testContentTypeHeaderKey))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetVideoHandler_IDNotFound(t *testing.T) {
	db, mock := getMockDBWithBadIDExpectations()
	e, err := getEchoRouterWithVideoHandlers(db)
	if err != nil {
		t.Errorf("unable to intialize video handlers due to: %v", err)
		return
	}

	URI, err := url.Parse(fmt.Sprintf("%s?id=%s", videoPath, testHandlerUUID))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, URI.String(), nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetVideoHandler_IDEmpty(t *testing.T) {
	db, mock := getMockDBWithNoExpectations()
	e, err := getEchoRouterWithVideoHandlers(db)
	if err != nil {
		t.Errorf("unable to intialize video handlers due to: %v", err)
		return
	}

	URI, err := url.Parse(fmt.Sprintf("%s?id=%s", videoPath, testHandlerEmptyUUID))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, URI.String(), nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetVideoHandler_IDNotUUID(t *testing.T) {
	db, mock := getMockDBWithNoExpectations()
	e, err := getEchoRouterWithVideoHandlers(db)
	if err != nil {
		t.Errorf("unable to intialize video handlers due to: %v", err)
		return
	}

	URI, err := url.Parse(fmt.Sprintf("%s?id=%s", videoPath, testHandlerInvalidUUID))
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, URI.String(), nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func getMockDBWithHappyPathExpectations() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	parsedID, _ := uuid.Parse(testHandlerUUID)
	row := sqlmock.NewRows([]string{"id", "uuid", "path"}).AddRow(testHandlerID, parsedID, testHandlerPath)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM video WHERE uuid = $1")).WithArgs(testHandlerUUID).WillReturnRows(row)

	return db, mock
}

func getMockDBWithBadIDExpectations() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	row := sqlmock.NewRows([]string{"id", "uuid", "path"})
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM video WHERE uuid = $1")).WithArgs(testHandlerUUID).WillReturnRows(row)

	return db, mock
}

func getMockDBWithNoExpectations() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	return db, mock
}

func getEchoRouterWithVideoHandlers(db *sql.DB) (*echo.Echo, error) {
	e := echo.New()
	e.Use(m.NewSQLX().WithConfig(m.SQLXConfig{
		DB:     db,
		Driver: testHandlerSQLDriver,
	}))
	e.Use(vsMiddleware.InjectFactory())
	e.Validator = maryread.NewValidator()
	storageServiceURL, err := getStorageServiceHost()
	video := NewVideoHandler(storageServiceURL)
	video.AddHandlers(e)
	return e, err
}

func getStorageServiceHost() (*url.URL, error) {
	storageServiceHostRaw := os.Getenv(testStorageServiceHostEnvKey)
	return url.Parse(storageServiceHostRaw)
}
