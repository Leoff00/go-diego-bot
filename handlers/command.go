package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "testt",
			Description: "command help test",
		},
	}
)

func addCmds(commands []*discordgo.ApplicationCommand, s *discordgo.Session) {

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.Application.ID, "", cmd)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (orp *OnReadyProps) OnReady() func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		addCmds(commands, s)
	}
}
