package bot

import (
	"fmt"
	"go-diego-bot/config"
	"strings"

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

	goBot.AddHandler(message)
	goBot.AddHandler(helpJava)

	goBot.Identify.Intents = discordgo.IntentGuildMessages

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func message(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == config.BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if m.Content == config.BotPrefix+"pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
	}
}

func helpJava(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	str := fmt.Sprintf(
		"Opa %s, uma bomba java? 💣, esses caras podem te ajudar 👇 \n %s", m.Author.Mention(), "<@241680344791646209>",
	)

	if strings.Contains(m.Content, config.BotPrefix+"java") == true {
		_, _ = s.ChannelMessageSend(m.ChannelID, str)
	}
}
