package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	ClientID    string `json:"CLIENT_ID"`
	RedirectURL string `json:"REDIRECT_URL"`
	Secret      string `json:"SECRET"`
	State       string `json:"STATE"`
}

func LoadConfig(filename string) *Config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config, ", err)
		return nil
	}
	var config Config
	json.Unmarshal(body, &config)
	return &config
}

func LoadConfigFromEnv() *Config {
	return &Config{
		ClientID: os.Getenv("SPOTIFY_CLIENT_ID"),
		RedirectURL: os.Getenv("SPOTIFY_REDIRECT_URL"),
		Secret: os.Getenv("SPOTIFY_SECRET"),
		State: os.Getenv("SPOTIFY_STATE"),
	}
}