package main

import (
	"github.com/Leoff00/go-diego-bot/bot"
	"github.com/Leoff00/go-diego-bot/config"
)

func main() {

	config.ReadConfig()

	bot.Start()

}
