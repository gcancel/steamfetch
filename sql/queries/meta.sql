-- name: SetDatabaseUpdateTime :one
UPDATE meta
SET last_update = (datetime(current_timestamp, 'localtime'))
RETURNING *;