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

func (h *HandlersProps) Img() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		responseAi := make(chan *AiResponse)
		errC := make(chan error)

		param := strings.Split(m.Content, " ")
		data := param[1]

		go PictureGenerator(data, responseAi, errC)

		select {
		case res := <-responseAi:
			var ogSize string

			for _, p := range res.Photos {
				ogSize = p.Src.Original
			}

			if m.Content == config.BotPrefix+"picture "+data {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Aqui esta o que você pediu! \n"+ogSize)
			}

		case err := <-errC:
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
			}
		}
	}
}

func (h *HandlersProps) HelpJava() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		str = fmt.Sprintf(
			"Opa %s, uma bomba JS 💣? Esses caras podem te ajudar 👇 \n %s \n %s \n %s",
			m.Author.Mention(),
			"<@209655533500628992>",
			"<@847935543018782772>",
			"<@599980091136671754>",
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

		if g.Member.JoinedAt == now {
			_, _ = s.ChannelMessageSend(s.State.User.ID, "Bem vindo"+s.State.User.Username)
		}
	}
}
