package main

import (
	"context"
	"database/sql"
	"log"
)

/* func writeEnv(key, value string) error {
	// append the values to the .env file in the project's root directory
	// check if the key already exists, if it does, return
	return nil
} */

func initDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./steam_games.db")
	if err != nil {
		log.Fatal("error connecting to database ", err)
	}
	initialQuery :=
		`CREATE TABLE IF NOT EXISTS steam_games(
		appid INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		playtime_forever INTEGER NOT NULL,
		img_icon_url TEXT NOT NULL,
		playtime_windows_forever INTEGER NOT NULL,
		playtime_mac_forever INTEGER NOT NULL,
		playtime_linux_forever INTEGER NOT NULL,
		playtime_deck_forever INTEGER NOT NULL,
		rtime_last_played INTEGER NOT NULL,
		playtime_disconnected INTEGER NOT NULL,
		playtime_2weeks INTEGER NOT NULL
    )`

	_, err = db.ExecContext(
		context.Background(),
		initialQuery,
	)
	if err != nil {
		log.Fatal("error creating database ", err)
	}
	return db, nil
}
