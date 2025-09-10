package main

import (
	"context"
	"fmt"
	"log"
)

func handleSteamFetchBacklog(s *state, cmd command) error {
	// returns comma delimited list of games you have never touched
	backlog, err := s.dbQueries.GetTotalGamesNotPlayed(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, game := range backlog {
		fmt.Printf("%v, %v\n", game.Name, game.Appid)
	}
	return nil
}
