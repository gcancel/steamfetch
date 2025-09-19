package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func getOwnedGames(url string) (GetOwnedGames, error) {
	steamClient := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return GetOwnedGames{}, fmt.Errorf("error during request")
	}
	res, err := steamClient.Do(req)
	if err != nil {
		return GetOwnedGames{}, fmt.Errorf("error retrieving response from: %v", url)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return GetOwnedGames{}, fmt.Errorf("error parsing response body")
	}

	var result GetOwnedGames
	err = json.Unmarshal(data, &result)
	if err != nil {
		return GetOwnedGames{}, fmt.Errorf("error unmarshalling data: %v", err)
	}
	return result, nil
}
