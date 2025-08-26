package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type state struct {
	// db connection will go here
	steamID     string
	steamAPIKey string
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env file\n")
		os.Exit(1)
	}

	applicationState := &state{
		steamID:     os.Getenv("STEAM_ID"),
		steamAPIKey: os.Getenv("STEAM_WEBAPI_KEY"),
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

func handleInitialStart() {
	// function will take input and update the environment variables.
	fmt.Println("Initial Start:")
	fmt.Println("steamfetch requires your steam_id and an active steam web api key to function.")

	var input string
	fmt.Printf("Enter steam id: ")
	// handle input...
	fmt.Scan(&input)
	fmt.Println(input)

	fmt.Printf("Enter steam web api key: ")
	// handle input...
	fmt.Scan(&input)
	fmt.Println(input)
	os.Exit(0)
}
