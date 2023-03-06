package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/bwmarrin/discordgo"
)

var wg sync.WaitGroup

var (
	str string
	arr []string
	//HandlerUtilFunctions struct
	huf *HandlerUtilFunctions
	//Channel to response API paralell
	responseAi = make(chan *AiResponse)
	//Channel to response ERROR API paralell
	errC = make(chan error)

	//Picture props

	//Array with contain photo Infos
	p *PhotoProps
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

				for _, v := range res.Photos {
					p = v
				}

				msgEmbed := &discordgo.MessageEmbed{
					Title:       "📸 | Aqui esta sua foto!",
					Description: p.Alt,
					Type:        discordgo.EmbedTypeImage,
					Color:       10,
					Image:       &discordgo.MessageEmbedImage{URL: p.Src.Original},
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Autor " + p.Photographer,
					},
				}

				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, msgEmbed)

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

		if strings.Contains(m.Content, config.BotPrefix+"java") {
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

func (h *HandlersProps) MCount() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		gc, _ := s.GuildWithCounts(i.GuildID)
		if i.Type == discordgo.InteractionApplicationCommand {

			switch i.ApplicationCommandData().Name {
			case "people":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("Temos exatamente %d pessoas no servidor ", gc.ApproximateMemberCount),
					},
				})
			}
		}
	}
}

func (h *HandlersProps) UAvatar() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		mem, err := s.GuildMember(i.GuildID, i.Member.User.ID)

		if err != nil {
			fmt.Println("Error loading the member struct", err)
		}

		msgEmbed := &discordgo.MessageEmbed{
			Title:       "📸 | " + mem.User.Username,
			Description: "Mas que magnifica foto!!",
			Type:        discordgo.EmbedTypeImage,
			Color:       10,
			Image:       &discordgo.MessageEmbedImage{URL: mem.AvatarURL("1024")},
		}

		if i.Type == discordgo.InteractionApplicationCommand {

			switch i.ApplicationCommandData().Name {
			case "avatar":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{msgEmbed},
					},
				})
			}
		}
	}
}

func (h *HandlersProps) HelpCmd() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		mem, err := s.GuildMember(i.GuildID, i.Member.User.ID)

		if err != nil {
			fmt.Println("Error loading the member struct", err)
		}

		msgEmbed := &discordgo.MessageEmbed{
			Title:       HelpStr1,
			Description: HelpStr2,
			Type:        discordgo.EmbedTypeArticle,
			Color:       10,
			Footer:      &discordgo.MessageEmbedFooter{Text: "Autor: " + mem.User.Username},
		}

		if i.Type == discordgo.InteractionApplicationCommand {

			switch i.ApplicationCommandData().Name {
			case "help":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{msgEmbed},
					},
				})
			}
		}
	}
}

func (h *HandlersProps) ClearMsg() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content != "" && strings.HasPrefix(m.Content, config.BotPrefix+"clear") {
			limit, err := strconv.Atoi(huf.ParamSeparator(m.Content))

			if limit == 0 || limit > 100 {
				return
			}

			if err != nil {
				fmt.Println("Failed to convert to int", err)
			}

			chMsg, _ := s.ChannelMessages(m.ChannelID, limit, "", "", "")
			msgs := make([]string, len(chMsg))

			for _, v := range chMsg {
				msgs = append(msgs, v.ID)
			}

			for i, _ := range msgs {
				time.Sleep(200)
				s.ChannelMessageDelete(m.ChannelID, msgs[i])
				time.Sleep(200)
			}

			msgEmbed := &discordgo.MessageEmbed{
				Title:       "| Mensagens deletadas! 🔨 ",
				Description: fmt.Sprintf("| Total de mensagens deletadas: **%s** 📰", strconv.Itoa(limit)),
				Footer: &discordgo.MessageEmbedFooter{
					Text: "Autor do comando -> " + m.Author.Username,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, msgEmbed)

		}
	}
}
