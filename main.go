package main

import (
	"fmt"

	"github.com/Leoff00/go-diego-bot/bot"
	"github.com/Leoff00/go-diego-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()
}
