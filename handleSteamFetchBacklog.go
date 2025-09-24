package main

import (
	"context"
	"fmt"
	"log"
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
			for _, game := range games {
				fmt.Printf("%v, %v\n", game.Name, game.Appid)
			}
			fmt.Printf("Total: %v", len(games))
			return nil
		}
		if flag == "--played" || flag == "-p" {
			playedGames, err := s.dbQueries.GetTotalGamesPlayed(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			for _, game := range playedGames {
				fmt.Printf("%v, %v\n", game.Name, game.Appid)
			}
			fmt.Printf("Total: %v", len(playedGames))
			return nil
		}

	}
	backlog, err := s.dbQueries.GetTotalGamesNotPlayed(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, game := range backlog {
		fmt.Printf("%v, %v\n", game.Name, game.Appid)
	}
	fmt.Printf("Total: %v", len(backlog))
	return nil
}
