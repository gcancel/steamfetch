package main

import (
	"context"
	"database/sql"
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

	// check if updating the current database is needed
	err = integrityCheck(s, totalGameTimeForever)
	if err != nil {
		log.Fatal(err)
	}

	// bullet
	arrow := setANSIText("->|", Blue)

	totalGameTimeForeverMins := int(totalGameTimeForever.Float64)
	totalGameTimeAllMins := printSteamGameTime(totalGameTimeForeverMins, arrow, minutesToMinutes)
	totalGameTimeAllHours := printSteamGameTime(totalGameTimeForeverMins, arrow, minutesToHours)
	totalGameTimeAllDays := printSteamGameTime(totalGameTimeForeverMins, arrow, minutesToDays)
	totalGameTimeAllMonths := printSteamGameTime(totalGameTimeForeverMins, arrow, minutesToMonths)
	totalGameTimeAllYears := printSteamGameTime(totalGameTimeForeverMins, arrow, minutesToYears)
	// Print out

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
	fmt.Printf("%v", printBacklogRatio(int(gameCount), len(totalGamesNotPlayed)))

	//fmt.Printf("%v %v\n", setANSIText("🎮 Total number of games:", Yellow), gameCount)
	//fmt.Printf("%v %v\n", setANSIText("😵‍💫 Total games backlog(not played):", Yellow), len(totalGamesNotPlayed))

	fmt.Printf("%v\n", totalGameTimeAllMins)
	fmt.Printf("%v\n", totalGameTimeAllHours)
	fmt.Printf("%v\n", totalGameTimeAllDays)
	fmt.Printf("%v\n", totalGameTimeAllMonths)
	fmt.Printf("%v\n", totalGameTimeAllYears)

	fmt.Printf("%v %v mins\n", setANSIText("Total steam gameplay time (2 week):", Yellow), int(totalGameTime2Weeks.Float64))
	fmt.Println(setANSIText("Most played games: ", Yellow))
	fmt.Println(setANSIText("----------------------------------", Blue))

	mostPlayedGames, err := s.dbQueries.GetTopPlayedGames(context.Background(), int64(mostPlayedLimit))
	if err != nil {
		log.Fatal(err)
	}
	for _, game := range mostPlayedGames {
		hrs, time := minutesToHours(game.PlaytimeForever)
		fmt.Printf("%v: %.2f %s\n", game.Name, hrs, time)
	}
	fmt.Println(setANSIText("----------------------------------", Blue))
	fmt.Printf("%v: %v", setANSIText("Last updated", Cyan), lastDBUpdate.String)
	return nil
}

func printSteamGameTime(mins int, bullet string, f func(m int) (float64, string)) string {
	time, measurement := f(mins)
	if bullet == "" {
		return fmt.Sprintf("%.2f %v", time, measurement)
	}
	return fmt.Sprintf("%v%.2f %v", bullet, time, measurement)

}

func printBacklogRatio(total, notPlayed int) string {
	percentage := float64(notPlayed) / float64(total) * 100
	labelText := setANSIText("🎮 Games Played/Not Played: ", Yellow)
	output := fmt.Sprintf("%v%v / %v (%.2f%%)\n", labelText, notPlayed, total, percentage)
	return output
}

func integrityCheck(s *state, gameTime sql.NullFloat64) error {
	result, err := getOwnedGames(s.getOwnedGamesAPIURL)
	if err != nil {
		log.Fatal(err)
	}

	allGames := result.Response.Games
	var currentTotal int
	for _, game := range allGames {
		currentTotal += int(game.PlaytimeForever)
	}

	//fmt.Printf("current: %v in database: %v\n", currentTotal, int(gameTime.Float64))
	if currentTotal != int(gameTime.Float64) {
		fmt.Printf("Steam game time has been recently accrued. Consider performing update for accuracy... %v mins\n", currentTotal)

		// turn into option
		// err := handleSteamFetchUpdate(s, command{name: "update", arguments: []string{"-f"}})
		// if err != nil {
		//	log.Fatal(err)
		// }

		// fmt.Println("update completed. please run steamfetch again.")
		// os.Exit(0)
	}
	return nil
}
