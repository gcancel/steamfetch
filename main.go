package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gcancel/steamfetch/internal/database"
	_ "github.com/mattn/go-sqlite3"

	"github.com/joho/godotenv"
)

type state struct {
	// db connection will go here
	steamID     string
	steamAPIKey string
	dbQueries   database.Queries
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env file\n")
		handleInitialStart()
	}

	db, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	applicationState := &state{
		steamID:     os.Getenv("STEAM_ID"),
		steamAPIKey: os.Getenv("STEAM_WEBAPI_KEY"),
		dbQueries:   *database.New(db),
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// registered commands:
	commands.register("steamfetch", handleSteamFetch) // default command if no arguments are passed
	commands.register("update", handleSteamFetchUpdate)

	var commandName string
	var commandArgs []string
	if len(os.Args) == 1 {
		commandName = "steamfetch"
	} else if len(os.Args) >= 2 {
		commandName = os.Args[1]
		commandArgs = os.Args[2:]
	}

	// if initial start of application
	if applicationState.steamAPIKey == "" || applicationState.steamID == "" {
		handleInitialStart()
	}

	err = commands.run(applicationState, command{name: commandName, arguments: commandArgs})
	if err != nil {
		log.Fatal(err)
	}

}
