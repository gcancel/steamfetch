package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gcancel/steamfetch/internal/database"
	"github.com/schollz/progressbar/v3"
)

func handleSteamFetchUpdate(s *state, cmd command) error {
	forceUpdate := false
	if len(cmd.arguments) >= 1 {
		if cmd.arguments[0] == "-f" || cmd.arguments[0] == "-force" {
			forceUpdate = true
			err := s.dbQueries.ClearSteamDB(context.Background())
			if err != nil {
				log.Fatal("error clearing database", err)
			}
		}
	}

	steamClient := &http.Client{Timeout: 30 * time.Second}
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%v&steamid=%v&include_appinfo=1", s.steamAPIKey, s.steamID)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return fmt.Errorf("error during request")
	}
	res, err := steamClient.Do(req)
	if err != nil {
		return fmt.Errorf("error retrieving response from: %v", url)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error parsing response body")
	}

	var result GetOwnedGames
	err = json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %v", err)
	}

	// take this data and put it in the db for use instead of constantly calling api
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
			bar.Add(1)
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
	return nil
}
