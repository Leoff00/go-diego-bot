package config

import (
	"github.com/Leoff00/go-diego-bot/envs"
)

var (
	Token     string
	BotPrefix string

	config *configProps
)

type configProps struct {
	Token string `json:"Token"`
}

func ReadConfig() {

	Token = envs.Getenv("AUTH_TOKEN")

}
