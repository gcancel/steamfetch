package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gcancel/steamfetch/internal/database"
	"github.com/schollz/progressbar/v3"
)

func handleSteamFetchUpdate(s *state, cmd command) error {
	forceUpdate := false
	if len(cmd.arguments) >= 1 {
		if cmd.arguments[0] == "-f" || cmd.arguments[0] == "--force" {
			forceUpdate = true
			err := s.dbQueries.ClearSteamDB(context.Background())
			if err != nil {
				log.Fatal("error clearing database", err)
			}
		}
	}

	// calling "GetOwnedGames" api
	result, err := getOwnedGames(s.getOwnedGamesAPIURL)
	if err != nil {
		log.Fatal(err)
	}

	gameCount, err := s.dbQueries.GetCountGames(context.Background())
	if err != nil {
		log.Fatal("error counting games in database", err)
	}
	if gameCount <= 1 || forceUpdate {
		steamGames := result.Response.Games

		bar := progressbar.NewOptions(len(steamGames),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(15),
			progressbar.OptionSetDescription("[cyan]updating database...[reset]"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[blue]o[reset]",
				SaucerHead:    "[cyan]0[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}))

		for _, game := range steamGames {
			err = bar.Add(1)
			if err != nil {
				log.Fatal(err)
			}
			_, err := s.dbQueries.InsertGame(
				context.Background(),
				database.InsertGameParams{
					Appid:                  game.Appid,
					Name:                   game.Name,
					PlaytimeForever:        game.PlaytimeForever,
					ImgIconUrl:             game.ImgIconURL,
					PlaytimeWindowsForever: game.PlaytimeWindowsForever,
					PlaytimeMacForever:     game.PlaytimeMacForever,
					PlaytimeLinuxForever:   game.PlaytimeLinuxForever,
					PlaytimeDeckForever:    game.PlaytimeDeckForever,
					RtimeLastPlayed:        game.RtimeLastPlayed,
					PlaytimeDisconnected:   game.PlaytimeDisconnected,
					Playtime2weeks:         game.Playtime2Weeks,
				},
			)
			if err != nil {
				log.Fatal("error populating database", err)
			}
			// add results print out here.
			//fmt.Printf("adding game: %v | appid: %v\n", game.Name, game.Appid)
		}
	}

	//fmt.Printf("results: %#v", result)
	dbmeta, err := s.dbQueries.SetDatabaseUpdateTime(context.Background())
	if err != nil {
		log.Fatal("error setting update timestamp ", err)
	}
	fmt.Printf("\nLast updated: %v\n", dbmeta.String)
	return nil
}
