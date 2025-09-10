package main

import (
	"context"
	"fmt"
	"log"
)

func handleSteamFetchBacklog(s *state, cmd command) error {
	// display games in your backlog (never played)
	backlog, err := s.dbQueries.GetTotalGamesNotPlayed(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, game := range backlog {
		fmt.Printf("%v, %v\n", game.Name, game.Appid)
	}
	return nil
}
