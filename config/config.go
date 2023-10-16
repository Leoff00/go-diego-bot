package config

import (
	"github.com/Leoff00/go-diego-bot/envs"
)

var (
	Token          string
	BotPrefix      string
	Google_Key     string
	SearchEngineId string

	config *configProps
)

type configProps struct {
	Token          string `json:"Token"`
	Google_Key     string `json:"Google_Key"`
	SearchEngineId string `json:"SearchEngineId"`
}

func ReadConfig() {

	Token = envs.Getenv("AUTH_TOKEN")
	Google_Key = envs.Getenv("GOOGLE_API_KEY")
	SearchEngineId = envs.Getenv("SEARCH_ENGINE_ID")
}
