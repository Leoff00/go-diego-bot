package main

import (
	"encoding/json"
	"fmt"
	"go-diego-bot/envs"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

type ConfigProps struct {
	BotPrefix string `json:"BotPrefix"`
}

const BOT_AUTH_PREFIX = "Bot "

var (
	BotId  string
	config *ConfigProps
)

func readConfig() (string, error) {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Default().Fatalln("Failed to read config file", err)
		return "", err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Default().Fatalln("Failed to unmarshal JSON struct", err)
		return "", err
	}

	return config.BotPrefix, nil
}

func main() {
	token := envs.Getenv("AUTH_TOKEN")
	_, err := readConfig()

	if err != nil {
		fmt.Println(err)
	}

	goBot, err := discordgo.New(BOT_AUTH_PREFIX + token)

	if err != nil {
		log.Fatalln("Failed to load DiscordGo instance", err)
	}

	usr, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err)
	}

	BotId = usr.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot running...")

}

// TODO: Implement message handler to send I/O and consume Session API
func messageHandler(s *discordgo.Session, m *discordgo.Disconnect) {}
