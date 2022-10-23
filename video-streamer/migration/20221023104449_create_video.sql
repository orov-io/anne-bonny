-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE video (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    path VARCHAR NOT NULL
);

CREATE UNIQUE INDEX idx_video_uuid ON video(uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP INDEX idx_video_uuid;

DROP TABLE video;

DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd
