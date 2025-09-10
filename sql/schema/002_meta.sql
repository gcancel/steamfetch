-- +goose Up
CREATE TABLE meta(
    steam_id TEXT PRIMARY KEY,
    last_update TEXT DEFAULT (datetime(current_timestamp, 'localtime'))
);
-- +goose Down
DROP TABLE meta;