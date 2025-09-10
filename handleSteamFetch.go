package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
)

func handleSteamFetch(s *state, cmd command) error {
	// will fetch data from the db to aggregate data and display in terminal
	mostPlayedLimit := 5
	if len(cmd.arguments) >= 1 {
		limit, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			log.Fatal("error parsing mostplayed limit", err)
		}
		if limit <= 0 || limit > 50 {
			fmt.Println("limit must be set between 0 and 50. using default (5)")
		} else {
			mostPlayedLimit = limit
		}

	}
	lastDBUpdate, err := s.dbQueries.GetLastDBUpdate(context.Background())
	if err != nil {
		log.Fatal("error retrieving last DB timestamp", err)
	}

	gameCount, err := s.dbQueries.GetCountGames(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	totalGamesNotPlayed, err := s.dbQueries.GetTotalGamesNotPlayed(context.Background())
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
	arrow := setANSIText("->|", Blue)
	logo := `
	⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
	⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠋⠉⠉⠉⠉⠻⣿⣿⣿⣿
	⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⢀⣴⣶⣶⣄⠀⠈⢻⣿⣿
	⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠀⠀⢸⣿⣿⣿⣿⡇⠀⢸⣿⣿
	⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠀⠀⠈⠻⠿⠿⠛⠁⠀⣸⣿⣿
	⠿⣿⣿⣿⣿⡿⠟⠛⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣼⣿⣿⣿
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣾⣿⣿⣿⣿⣿⣿
	⣶⣶⣤⡀⠀⠀⣴⣿⣦⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿
	⣿⣿⣿⣧⡀⠀⠙⠿⠋⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
	⣿⣿⣿⣿⣷⣤⣀⣀⣀⣀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
	⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
`
	fmt.Println(setANSIText(logo, Blue))
	fmt.Printf("%v %v\n", setANSIText("Your steam id:", Yellow), s.steamID)
	fmt.Printf("%v %v\n", setANSIText("🎮 Total number of games:", Yellow), gameCount)
	fmt.Printf("%v %v\n", setANSIText("😵‍💫 Total games backlog(not played):", Yellow), len(totalGamesNotPlayed))
	fmt.Printf("%v\n%s%v minutes\n%s%v hours\n%s%v days\n%s%v month/s\n%s%v...\n",
		setANSIText("⌛ Total steam gameplay time:", Yellow),
		arrow,
		totalGameTimeForeverMins,
		arrow,
		totalGameTimeAllHours,
		arrow,
		totalGameTimeAllDays,
		arrow,
		totalGameTimeAllMonths,
		arrow,
		totalGameTimeAllYears,
	)

	fmt.Printf("%v %v minutes\n", setANSIText("Total steam gameplay time (2 week):", Yellow), int(totalGameTime2Weeks.Float64))
	fmt.Println(setANSIText("Most played games:", Yellow))
	fmt.Println(setANSIText("----------------------------------", Blue))

	mostPlayedGames, err := s.dbQueries.GetTopPlayedGames(context.Background(), int64(mostPlayedLimit))
	if err != nil {
		log.Fatal(err)
	}
	for _, game := range mostPlayedGames {
		hrs, min := minutesToHours(game.PlaytimeForever)
		fmt.Printf("%s%v: %vhrs, %vmins\n", arrow, game.Name, hrs, min)
	}
	fmt.Println(setANSIText("----------------------------------", Blue))
	fmt.Printf("%v: %v", setANSIText("Last updated", Cyan), lastDBUpdate.String)
	return nil
}
