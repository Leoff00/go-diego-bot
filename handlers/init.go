package handlers

import (
	"fmt"
	"strings"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	str        string
	arr        []string
	huf        *HandlerUtilFunctions
	responseAi = make(chan *AiResponse)
	errC       = make(chan error)
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

		if m.Content != "" && strings.HasPrefix(m.Content, config.BotPrefix+"picture") {
			data := huf.ParamSeparator(m.Content)

			go huf.PicGenerator(data, responseAi, errC)

			select {
			case res := <-responseAi:
				var ogSize string

				for _, p := range res.Photos {
					ogSize = p.Src.Original
				}
				_, _ = s.ChannelMessageSend(m.ChannelID, "Aqui esta o que você pediu! \n"+ogSize)

			case err := <-errC:
				if err != nil {
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
				}
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
		if m.Content == "oi diego" {
			_, _ = s.ChannelMessageSend(m.ChannelID, huf.RandPh(m.Author.Username))
		}
	}
}

func (h *HandlersProps) MsgHelpCmd() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		helpStr := fmt.Sprintf(`
Iaee %s meu nome é Die**go**, bot em go feito pra te ajudar com algumas 
utilidades no server esses são os comandos pelo qual eu respondo:
**oi diego -> responderei você de volta!**
**!picture [parametros] -> gerarei pra você uma imagem com o dado que você me forneceu!**
**!java [mensagem] -> marcarei 3 pessoas que manjam de java no server para te ajudar!**
**!ping ou !pong -> jogarei um ping pong com você :)!** `, m.Author.Username)

		if m.Content == config.BotPrefix+"help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, helpStr)
		}
	}
}

func (h *HandlersProps) Intest() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {

			switch i.ApplicationCommandData().Name {
			case "testt":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "testt!",
					},
				})
			}
		}
	}
}
