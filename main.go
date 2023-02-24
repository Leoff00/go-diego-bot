package main

import (
	"fmt"

	"github.com/Leoff00/go-diego-bot/bot"
	"github.com/Leoff00/go-diego-bot/config"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered", r)
		}
	}()

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

}
