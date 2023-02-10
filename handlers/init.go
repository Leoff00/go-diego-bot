package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	str string
	arr []string
)

func (h *HandlersProps) MessagePingPong() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
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

func (h *HandlersProps) HelpJava() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		str = fmt.Sprintf(
			"Opa %s, uma bomba java 💣? Esses caras podem te ajudar 👇 \n %s \n %s \n %s",
			m.Author.Mention(),
			"<@241680344791646209>",
			"<@430150392068702229>",
			"<@958819349580349490>",
		)

		if strings.Contains(m.Content, config.BotPrefix+"java") == true {
			_, _ = s.ChannelMessageSend(m.ChannelID, str)
		}
	}
}

func (h *HandlersProps) Greeting() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "Oi diego" {
			_, _ = s.ChannelMessageSend(m.ChannelID, RandPhrase(m.Author.Username))
		}
	}
}

func (h *HandlersProps) NotifyNewMember() func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	return func(s *discordgo.Session, g *discordgo.GuildMemberAdd) {

		now := time.Now().UTC().Local()

		if g.Member.JoinedAt.UTC().Local() == now {
			_, _ = s.ChannelMessageSend(s.State.SessionID, "Bem vindo")
		}
	}
}
