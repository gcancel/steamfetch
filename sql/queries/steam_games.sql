-- name: GetCountGames :one
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

-- name: GetTotalGameTimeForever :one
SELECT SUM(playtime_forever) FROM steam_games;

-- name: GetTotalGameTime2Weeks :one
SELECT SUM(playtime_2weeks) FROM steam_games;

-- name: GetTopPlayedGames :many
SELECT appid, name, playtime_forever FROM steam_games
ORDER BY playtime_forever DESC LIMIT ?;

-- name: ClearSteamDB :exec
DELETE FROM steam_games;

-- name: GetTotalGamesNotPlayed :many
SELECT name, appid from steam_games
WHERE playtime_forever = 0
ORDER BY name;