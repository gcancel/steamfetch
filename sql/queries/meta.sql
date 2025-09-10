-- name: SetDatabaseUpdateTime :one
UPDATE meta
SET last_updated = (datetime(current_timestamp, 'localtime'))
RETURNING *;

-- name: GetLastDBUpdate :one
SELECT last_updated FROM meta;