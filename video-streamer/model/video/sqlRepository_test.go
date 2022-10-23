package video

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

const (
	videoSQLRepoTestValidUUID   = "f0fd7ef3-86f9-499c-9853-fac245946e0c"
	videoSQLRepoTestInvalidUUID = "cd1ad06b-e338-4ba4-9879-9e9f04f9aa9e"
	videoSQLRepoTestPath        = "/path"
	videoSQLRepoTestID          = 1
	videoSQLRepoTestSQLDriver   = "postgres"
)

func TestSQLRepository_GetVideo(t *testing.T) {
	db, mock, _ := sqlmock.New()
	dbx := sqlx.NewDb(db, videoSQLRepoTestSQLDriver)
	repo := NewVideoSQLRepository(context.Background(), dbx)
	parsedID, _ := uuid.Parse(videoSQLRepoTestValidUUID)
	row := sqlmock.NewRows([]string{"id", "uuid", "path"}).AddRow(videoSQLRepoTestID, parsedID, videoSQLRepoTestPath)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM video WHERE uuid = $1")).WithArgs(videoSQLRepoTestValidUUID).WillReturnRows(row)
	video, err := repo.GetVideo(videoSQLRepoTestValidUUID)
	assert.NoError(t, err)

	assert.Equal(t, parsedID, video.ID)
	assert.Equal(t, videoSQLRepoTestPath, video.Path)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLRepository_GetVideo_DBError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	dbx := sqlx.NewDb(db, videoSQLRepoTestSQLDriver)
	repo := NewVideoSQLRepository(context.Background(), dbx)
	row := sqlmock.NewRows([]string{"id", "uuid", "path"})

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM video WHERE uuid = $1")).WithArgs(videoSQLRepoTestInvalidUUID).WillReturnRows(row)
	video, err := repo.GetVideo(videoSQLRepoTestInvalidUUID)
	assert.Error(t, err)
	assert.Nil(t, video)
	assert.Equal(t, ErrVideoNotFound, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
