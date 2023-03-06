package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func addCmds(commands []*discordgo.ApplicationCommand, s *discordgo.Session) {

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.Application.ID, "", cmd)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (sp *SetupProps) OnReady() func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		addCmds(sp.Commands, s)
	}
}

func (sp *SetupProps) DeleteCommands() func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {

		commands, err := s.ApplicationCommands(sp.AppId, "")
		if err != nil {
			fmt.Println("Hasn't command to read ", err)
			return
		}
		for _, command := range commands {
			err = s.ApplicationCommandDelete(sp.AppId, "", command.ID)
			if err != nil {
				fmt.Println("Cannot remove the commands ", err)
			}
		}
	}
}
