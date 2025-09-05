-- +goose Up
CREATE TABLE steam_games(
    appid INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    playtime_forever INTEGER NOT NULL,
    img_icon_url TEXT NOT NULL,
    playtime_windows_forever INTEGER NOT NULL,
    playtime_mac_forever INTEGER NOT NULL,
    playtime_linux_forever INTEGER NOT NULL,
    playtime_deck_forever INTEGER NOT NULL,
    rtime_last_played INTEGER NOT NULL,
    playtime_disconnected INTEGER NOT NULL,
    playtime_2weeks INTEGER NOT NULL
    );
-- +goose Down
DROP TABLE games;