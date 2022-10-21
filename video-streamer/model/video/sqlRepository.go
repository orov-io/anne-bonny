package videoModel

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VideoSQL struct {
	ID   uint   `db:"id"`
	UUID string `db:"uuid"`
	Path string `db:"path"`
}

func (v VideoSQL) ToEntity() (*Video, error) {
	parsedUUID, err := uuid.Parse(v.UUID)
	if err != nil {
		return nil, err
	}

	return &Video{
		ID:   parsedUUID,
		Path: v.Path,
	}, nil
}

type VideoSQLRepository struct {
	// TODO: Review this when code the test to know if we can inject a mock database.
	db  *sqlx.DB
	ctx context.Context
}

func NewVideoSQLRepository(ctx context.Context, db *sqlx.DB) *VideoSQLRepository {
	return &VideoSQLRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *VideoSQLRepository) GetVideo(id string) (*Video, error) {
	dbVideo := VideoSQL{}
	// TODO: Use squirrel
	err := r.db.Select(&dbVideo, "SELECT * FROM video WHERE uuid=$1", id)
	if err != nil {
		return nil, err
	}

	return dbVideo.ToEntity()
}
