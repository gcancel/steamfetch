-- +goose Up
CREATE TABLE meta(
    last_updated TEXT DEFAULT (datetime(current_timestamp, 'localtime'))
);
INSERT INTO meta(last_updated)
VALUES(
    (datetime(current_timestamp, 'localtime'))
);
-- +goose Down
DROP TABLE meta;