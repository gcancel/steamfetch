package main

type GetOwnedGames struct {
	Response struct {
		GameCount int `json:"game_count"`
		Games     []struct {
			Appid                  int64  `json:"appid"`
			Name                   string `json:"name"`
			PlaytimeForever        int64  `json:"playtime_forever"`
			ImgIconURL             string `json:"img_icon_url"`
			PlaytimeWindowsForever int64  `json:"playtime_windows_forever"`
			PlaytimeMacForever     int64  `json:"playtime_mac_forever"`
			PlaytimeLinuxForever   int64  `json:"playtime_linux_forever"`
			PlaytimeDeckForever    int64  `json:"playtime_deck_forever"`
			RtimeLastPlayed        int64  `json:"rtime_last_played"`
			PlaytimeDisconnected   int64  `json:"playtime_disconnected"`
			Playtime2Weeks         int64  `json:"playtime_2weeks,omitempty"`
		} `json:"games"`
	} `json:"response"`
}
