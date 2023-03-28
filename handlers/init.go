package handlers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	str string
	arr []string
	//HandlerUtilFunctions struct
	huf *HandlerUtilFunctions
	//Channel to response API paralell
	responseAi = make(chan *AiResponse)
	//Channel to response ERROR API paralell
	errC = make(chan error)

	//Array that contains photo Infos
	p *PhotoProps
)

func (h *HandlersProps) MessagePingPong() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "ping":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   discordgo.MessageFlagsEphemeral,
						Content: "pong",
					},
				})
			case "pong":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   discordgo.MessageFlagsEphemeral,
						Content: "ping",
					},
				})
			}
		}
	}
}

func (h *HandlersProps) Img() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "img":
				param := i.ApplicationCommandData().Options[0].StringValue()

				if param != "" {
					go huf.PicGenerator(param, responseAi, errC)
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
						s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Flags:  discordgo.MessageFlagsEphemeral,
								Embeds: []*discordgo.MessageEmbed{msgEmbed},
							},
						})

					case err := <-errC:
						if err != nil {
							msgEmbed := &discordgo.MessageEmbed{
								Title: "A foto nao foi gerada corretamente... ",
								Type:  discordgo.EmbedTypeImage,
								Color: 10,
							}
							s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
								Type: discordgo.InteractionResponseChannelMessageWithSource,
								Data: &discordgo.InteractionResponseData{
									Flags:  discordgo.MessageFlagsEphemeral,
									Embeds: []*discordgo.MessageEmbed{msgEmbed},
								},
							})
						}
					}
				}
			}
		}
	}
}

func (h *HandlersProps) HelpJava() func(s *discordgo.Session, m *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "help_bombas":
				if strings.Contains(i.ApplicationCommandData().Options[0].StringValue(), "js") {

					str = fmt.Sprintf(
						"Opa %s, uma bomba em %s 💣? Esses caras de %s e %s podem te ajudar 👇",
						i.Member.User.Username,
						i.ApplicationCommandData().Options[0].StringValue(),
						"<@&1021801776116142110>",
						"<@&1021799560026259496>",
					)

					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Flags:   discordgo.MessageFlagsEphemeral,
							Content: str,
						},
					})
				}

				if strings.Contains(i.ApplicationCommandData().Options[0].StringValue(), "java") {

					str = fmt.Sprintf(
						"Opa %s, uma bomba em %s 💣? Esses caras de %s podem te ajudar 👇",
						i.Member.User.Username,
						i.ApplicationCommandData().Options[0].StringValue(),
						"<@&1021801356178239510>",
					)

					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Flags:   discordgo.MessageFlagsEphemeral,
							Content: str,
						},
					})
				}

			}
		}
	}
}

func (h *HandlersProps) Greeting() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "oi_diego":
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:   discordgo.MessageFlagsEphemeral,
						Content: huf.RandPh(i.Member.User.Username),
					},
				})
			}
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
						Flags:   discordgo.MessageFlagsEphemeral,
						Content: fmt.Sprintf("Temos exatamente %d pessoas no servidor ", gc.ApproximateMemberCount),
					},
				})
			}
		}
	}
}

func (h *HandlersProps) UAvatar() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "avatar":
				regex := regexp.MustCompile(`\d`)
				userId := strings.Join(regex.FindAllString(i.ApplicationCommandData().Options[0].StringValue(), -1), "")

				if userId != "" {
					mem, err := s.GuildMember(i.GuildID, userId)

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

					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Flags:  discordgo.MessageFlagsEphemeral,
							Embeds: []*discordgo.MessageEmbed{msgEmbed},
						},
					})

				}
			}
		}
	}
}

func (h *HandlersProps) HelpCmd() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "help":
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
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags:  discordgo.MessageFlagsEphemeral,
						Embeds: []*discordgo.MessageEmbed{msgEmbed},
					},
				})
			}
		}
	}
}

func (h *HandlersProps) ClearMsg() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			switch i.ApplicationCommandData().Name {
			case "clear":
				limit, err := strconv.Atoi(i.ApplicationCommandData().Options[0].StringValue())

				if err != nil {
					fmt.Println("Cannot convert string to integer", err)
				}

				mem, err := s.GuildMember(i.GuildID, i.Member.User.ID)

				if err != nil {
					fmt.Println("Error loading the member struct", err)
				}

				if limit <= 0 || limit > 100 {
					return
				}

				if mem.Roles[0] == "920531812760051722" || mem.Roles[0] == "610527002830569482" || mem.Roles[0] == "920532087881203713" {

					chMsg, _ := s.ChannelMessages(i.ChannelID, limit, "", "", "")

					for _, v := range chMsg {
						msgs := make([]string, 0)
						msgs = append(msgs, v.ID)
						time.Sleep(500)

						err := s.ChannelMessagesBulkDelete(i.ChannelID, msgs)

						if err != nil {
							fmt.Println("Error during deleting msgs", err)
						}

					}

					msgEmbed := &discordgo.MessageEmbed{
						Title:       "| Mensagens deletadas! 🔨 ",
						Description: fmt.Sprintf("| Total de mensagens deletadas: **%s** 📰", strconv.Itoa(limit)),
						Footer: &discordgo.MessageEmbedFooter{
							Text: "Autor do comando -> " + i.Member.User.Username,
						},
					}
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Flags:  discordgo.MessageFlagsEphemeral,
							Embeds: []*discordgo.MessageEmbed{msgEmbed},
						},
					})

				} else {
					msgEmbed := &discordgo.MessageEmbed{
						Title:       "| Você não tem permissão pra usar esse comando! ",
						Description: "| As mensagens não foram deletadas.",
						Footer: &discordgo.MessageEmbedFooter{
							Text: "Autor do comando -> " + i.Member.User.Username,
						},
					}
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Flags:  discordgo.MessageFlagsEphemeral,
							Embeds: []*discordgo.MessageEmbed{msgEmbed},
						},
					})
				}
			}
		}
	}
}
