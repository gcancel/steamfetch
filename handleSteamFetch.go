package main

import (
	"context"
	"fmt"
	"log"
)

func handleSteamFetch(s *state, cmd command) error {
	// will fetch data from the db to aggregate data and display in terminal
	gameCount, err := s.dbQueries.GetCountGames(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	totalGameTimeForever, err := s.dbQueries.GetTotalGameTimeForever(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	totalGameTime2Weeks, err := s.dbQueries.GetTotalGameTime2Weeks(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	totalGameTimeForeverMins := int(totalGameTimeForever.Float64)
	totalGameTimeAllHours := totalGameTimeForeverMins / 60
	totalGameTimeAllDays := totalGameTimeAllHours / 24
	totalGameTimeAllMonths := totalGameTimeAllDays / 30
	totalGameTimeAllYears := fmt.Sprintf("%.1f year/s", float64(totalGameTimeAllMonths)/12)
	// Print out
	fmt.Printf("Your steam id: %v\n", s.steamID)
	fmt.Printf("Total number of games: %v\n", gameCount)
	fmt.Printf("Total steam gameplay time:\n%v minutes\n->| %v hours\n->| %v days\n->| %v month/s\n->| %v...\n",
		totalGameTimeForeverMins,
		totalGameTimeAllHours,
		totalGameTimeAllDays,
		totalGameTimeAllMonths,
		totalGameTimeAllYears,
	)

	fmt.Printf("Total steam gameplay time (2 week period): %v minutes\n", int(totalGameTime2Weeks.Float64))
	fmt.Println("Most played games:")
	limit := 5
	mostPlayedGames, err := s.dbQueries.GetTopPlayedGames(context.Background(), int64(limit))
	if err != nil {
		log.Fatal(err)
	}
	for _, game := range mostPlayedGames {
		fmt.Printf("->|%v: %v mins\n", game.Name, game.PlaytimeForever)
	}
	// create query for most played games
	fmt.Printf("...")

	return nil
}
