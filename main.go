package main

import (
	"fmt"
	"go-diego-bot/bot"
	"go-diego-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
