package model

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/orov-io/anne-bonny/video-streamer/model/video"
)

const factoryTestSQLDriver = "postgres"

func TestNewFactory(t *testing.T) {
	db, _, _ := sqlmock.New()
	dbx := sqlx.NewDb(db, factoryTestSQLDriver)
	ctx := context.Background()
	type args struct {
		ctx context.Context
		dbx *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *Factory
	}{
		{
			name: "Generates a valid factory",
			args: args{
				ctx: ctx,
				dbx: dbx,
			},
			want: &Factory{
				Video: video.NewVideoSQLRepository(ctx, dbx),
				ctx:   ctx,
				dbx:   dbx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFactory(tt.args.ctx, tt.args.dbx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
