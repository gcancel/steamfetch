package config

type Config struct {
	SteamID string `json:"steam_id"`
}

const configFileName = "/.blfetchconfig.json"

func Read() (Config, error) {
	return Config{}, nil
}

func (c Config) write() error {
	return nil
}

func (c *Config) SetSteamID(steamID string) error {
	return nil
}

func getConfigFilePath() (string, error) {
	return "", nil
}
