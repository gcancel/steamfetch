package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func handleSteamFetchUpdate(s *state, cmd command) error {
	steamClient := &http.Client{Timeout: 30 * time.Second}
	fmt.Println(s.steamAPIKey, s.steamID)
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%v&steamid=%v&include_appinfo=true", s.steamAPIKey, s.steamID)

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
	// will contain logic to update the ~/steamfetch_data.json file with your data

	fmt.Printf("results: %v", result)
	return nil
}
