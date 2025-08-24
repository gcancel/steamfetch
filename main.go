package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gcancel/blfetch/internal/config"
)

type state struct {
	config      *config.Config
	steamAPIKey string
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config file...", err)
	}

	applicationState := &state{
		config:      &cfg,
		steamAPIKey: os.Getenv("STEAM_WEBAPI_KEY"),
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("setid", handleSetSteamID)

	if len(os.Args) > 2 {
		log.Fatal("Usage: blfetch [args...]")
		return
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	fmt.Println("fetching Steam stats...")
	err = commands.run(applicationState, command{name: commandName, arguments: commandArgs})
	if err != nil {
		log.Fatal(err)
	}

}
