package main

import (
	"context"
	"database/sql"
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
		file, err := os.OpenFile("./.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
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
	err = writeEnv("STEAM_ID", input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter steam web api key: ")
	// handle input...
	_, err = fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	err = writeEnv("STEAM_WEBAPI_KEY", input)
	if err != nil {
		log.Fatal(err)
	}
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
    );
		CREATE TABLE IF NOT EXISTS meta(
		last_updated TEXT DEFAULT (datetime(current_timestamp, 'localtime'))
	);
		INSERT INTO meta(last_updated)
		VALUES(
			(datetime(current_timestamp, 'localtime'))
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

/* func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
} */

// ANSI Escape Color Codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func setANSIText(s, code string) string {
	return code + s + Reset
}

func minutesToMinutes(mins int) (float64, string) {
	return float64(mins), "mins"
}

func minutesToHours(mins int) (float64, string) {
	return float64(mins) / 60, "hrs"
}

func minutesToDays(mins int) (float64, string) {
	hrs, _ := minutesToHours(mins)
	return float64(hrs) / 24, "days"
}

func minutesToMonths(mins int) (float64, string) {
	days, _ := minutesToDays(mins)
	return float64(days) / 30, "months"
}

func minutesToYears(mins int) (float64, string) {
	months, _ := minutesToMonths(mins)
	return float64(months) / 12, "years"
}
