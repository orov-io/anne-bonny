package video

import sq "github.com/Masterminds/squirrel"

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	videoTable = "video"
	uuidColumn = "uuid"
)

func getVideoByUUIDSQL(uuid string) sq.SelectBuilder {
	videos := allVideosSQL()
	return videos.Where(sq.Eq{uuidColumn: uuid})
}

func allVideosSQL() sq.SelectBuilder {
	return psql.Select("*").From(videoTable)
}
