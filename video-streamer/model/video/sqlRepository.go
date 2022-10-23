package video

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type videoSQL struct {
	ID   uint   `db:"id"`
	UUID string `db:"uuid"`
	Path string `db:"path"`
}

func (v videoSQL) ToEntity() (*Video, error) {
	parsedUUID, err := uuid.Parse(v.UUID)
	if err != nil {
		return nil, err
	}

	return &Video{
		ID:   parsedUUID,
		Path: v.Path,
	}, nil
}

type SQLRepository struct {
	// TODO: Review this when code the test to know if we can inject a mock database.
	dbx *sqlx.DB
	ctx context.Context
}

func NewVideoSQLRepository(ctx context.Context, db *sqlx.DB) *SQLRepository {
	return &SQLRepository{
		dbx: db,
		ctx: ctx,
	}
}

func (r *SQLRepository) GetVideo(id string) (*Video, error) {
	dbVideo := videoSQL{}
	// TODO: Use squirrel
	err := r.dbx.Get(&dbVideo, "SELECT * FROM video WHERE uuid=$1", id)
	if err != nil {
		return nil, err
	}

	return dbVideo.ToEntity()
}
