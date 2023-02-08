package config

import (
	"encoding/json"
	"fmt"
	"go-diego-bot/envs"
	"os"
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
		fmt.Println((err.Error()))
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = envs.Getenv("AUTH_TOKEN")
	BotPrefix = config.BotPrefix

	return nil
}
