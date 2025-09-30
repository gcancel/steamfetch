package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gcancel/steamfetch/internal/database"
)

func handleSteamFetchBacklog(s *state, cmd command) error {
	// returns comma delimited list of games you have never touched
	//--all -a and --played -p flags will display out all games and played games respectively
	if len(cmd.arguments) >= 1 {
		flag := cmd.arguments[0]
		if flag == "--all" || flag == "-a" {
			games, err := s.dbQueries.GetTotalGamesAll(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			printGameResults(games)
			return nil
		}
		if flag == "--played" || flag == "-p" {
			playedGames, err := s.dbQueries.GetTotalGamesPlayed(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			printGameResults(playedGames)
			return nil
		}

	}
	backlog, err := s.dbQueries.GetTotalGamesNotPlayed(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	printGameResults(backlog)
	return nil
}

func printGameResults(games interface{}) {
	switch s := games.(type) {
	case []database.GetTotalGamesNotPlayedRow:
		for _, game := range s {
			fmt.Printf("%v, %v\n", game.Name, game.Appid)
		}
		fmt.Printf("Total: %v", len(s))
	case []database.GetTotalGamesAllRow:
		for _, game := range s {
			fmt.Printf("%v, %v\n", game.Name, game.Appid)
		}
		fmt.Printf("Total: %v", len(s))
	case []database.GetTotalGamesPlayedRow:
		for _, game := range s {
			fmt.Printf("%v, %v\n", game.Name, game.Appid)
		}
		fmt.Printf("Total: %v", len(s))
	}

}
