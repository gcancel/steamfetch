-- name: GetCountGames :exec
SELECT COUNT(*) FROM steam_games;

-- name: InsertGame :one
INSERT INTO steam_games(appid, name, playtime_forever, img_icon_url, playtime_windows_forever, playtime_mac_forever, playtime_linux_forever, playtime_deck_forever, rtime_last_played, playtime_disconnected, playtime_2weeks)
VALUES(
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)RETURNING *;
