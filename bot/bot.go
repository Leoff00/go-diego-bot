package bot

import (
	"fmt"
	"go-diego-bot/config"
	"go-diego-bot/handlers"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

type HandlersProps struct {
	msgPingPongHanlder func(s *discordgo.Session, m *discordgo.MessageCreate)
	helpJavaHandler    func(s *discordgo.Session, m *discordgo.MessageCreate)
	msgGreeting        func(s *discordgo.Session, m *discordgo.MessageCreate)
}

func HandlerInitializer() *HandlersProps {
	handler := HandlersProps{
		msgPingPongHanlder: handlers.MessagePingPong(BotID),
		helpJavaHandler:    handlers.HelpJava(BotID),
		msgGreeting:        handlers.GreetingMessage(BotID),
	}
	return &handler
}

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

	goBot.Identify.Intents = discordgo.IntentGuildMessages

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}
