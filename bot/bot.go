package bot

import (
	"fmt"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(HandlerInitializer().msgPingPongHanlder)
	goBot.AddHandler(HandlerInitializer().helpJavaHandler)
	goBot.AddHandler(HandlerInitializer().msgGreeting)
	goBot.AddHandler(HandlerInitializer().notifyNewMember)

	goBot.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentMessageContent |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuilds

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}
