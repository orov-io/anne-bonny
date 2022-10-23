package model

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/orov-io/anne-bonny/video-streamer/model/video"
)

type Factory struct {
	Video video.Repository
	ctx   context.Context
	dbx   *sqlx.DB
}

func NewFactory(ctx context.Context, dbx *sqlx.DB) *Factory {
	return &Factory{
		Video: video.NewVideoSQLRepository(ctx, dbx),
		ctx:   ctx,
		dbx:   dbx,
	}
}
