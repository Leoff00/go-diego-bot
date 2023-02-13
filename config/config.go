package config

import (
	"encoding/json"
	"os"

	"github.com/Leoff00/go-diego-bot/envs"
)

var (
	Token     string
	BotPrefix string

	config *configProps
)

type configProps struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	file, err := os.ReadFile("./config.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		return err
	}

	Token = envs.Getenv("AUTH_TOKEN")
	BotPrefix = config.BotPrefix

	return nil
}
