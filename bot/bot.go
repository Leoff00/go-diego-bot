package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/Leoff00/go-diego-bot/handlers"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string

	handler = handlers.HandlersProps{}
	sp      = handlers.SetupProps{AppId: BotID, Commands: handlers.Commands, GuildId: ""}
)

func addHandler(goBot *discordgo.Session) {
	goBot.AddHandlerOnce(sp.OnReady())

	goBot.AddHandler(handler.MessagePingPong())
	goBot.AddHandler(handler.HelpJava())
	goBot.AddHandler(handler.Greeting())
	goBot.AddHandler(handler.Img())
	goBot.AddHandler(handler.MCount())
	goBot.AddHandler(handler.HelpCmd())
	goBot.AddHandler(handler.UAvatar())
	goBot.AddHandler(handler.ClearMsg())
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

	addHandler(goBot)

	goBot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	sp.DeleteCommands()

	err = goBot.Open()
	defer goBot.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")

	fmt.Println("Press Ctrl + C to exit.")
	stsignal := make(chan os.Signal, 1)
	signal.Notify(stsignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stsignal
}
