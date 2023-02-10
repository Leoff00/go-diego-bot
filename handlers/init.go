package handlers

import (
	"fmt"
	"go-diego-bot/config"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	str string
	arr []string
)

func randPhrase(user string) string {

	rand.Seed(time.Now().Unix())

	g1 := fmt.Sprintf("Ola %s!", user)
	g2 := fmt.Sprintf("Iaee %s!", user)
	g3 := fmt.Sprintf("Oiee %s!", user)

	arr = append(arr, g1, g2, g3)
	return arr[rand.Intn(len(arr))]

}

func MessagePingPong(BotID string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
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
}

func HelpJava(BotID string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == BotID {
			return
		}
		str = fmt.Sprintf(
			"Opa %s, uma bomba java 💣? E sses caras podem te ajudar 👇 \n %s \n %s \n %s", m.Author.Mention(),
			"<@241680344791646209>",
			"<@430150392068702229>",
			"<@958819349580349490>",
		)

		if strings.Contains(m.Content, config.BotPrefix+"java") == true {
			_, _ = s.ChannelMessageSend(m.ChannelID, str)
		}
	}
}

func GreetingMessage(BotID string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "Oi diego" {
			_, _ = s.ChannelMessageSend(m.ChannelID, randPhrase(m.Author.Username))
		}
	}
}
