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
	botPrefix string
	BotId     string
	config    *ConfigProps
)

func readConfig() error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Default().Fatalln("Failed to read config file", err)
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Default().Fatalln("Failed to unmarshal JSON struct", err)
		return err
	}

	botPrefix = config.BotPrefix

	return nil
}

func messageHandler(s *discordgo.Session, m *discordgo.Message) {
	if m.Author.ID != BotId {
		return
	}

	if m.Content == botPrefix+"ping" {
		fmt.Println(m.ChannelID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}

func Start() {
	token := envs.Getenv("AUTH_TOKEN")

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
		log.Default().Fatalln(err)
	}

	fmt.Println("Bot is running...")
}

func main() {
	err := readConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	Start()

	<-make(chan struct{})
}
