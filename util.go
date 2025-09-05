package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

func writeEnv(key, value string) error {
	// append the values to the .env file in the project's root directory
	// check if the key already exists, if it does, return
	if os.Getenv(key) == key {
		fmt.Println("key already exists.")
		os.Exit(0)
	}
	if os.Getenv(key) == "" {
		entry := fmt.Sprintf("%v=\"%v\"\n", key, value)
		file, err := os.OpenFile("./.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(entry)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func handleInitialStart() {
	// function will take input and update the environment variables.
	fmt.Println("Initial Start:")
	fmt.Println("steamfetch requires your steam_id and an active steam web api key to function.")

	var input string
	fmt.Printf("Enter steam id: ")
	// handle input...

	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	writeEnv("STEAM_ID", input)

	fmt.Printf("Enter steam web api key: ")
	// handle input...
	_, err = fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	writeEnv("STEAM_WEBAPI_KEY", input)
	fmt.Println("settings updated. please restart steamfetch to use...")
	os.Exit(0)
}

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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}
