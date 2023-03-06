package handlers

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "testt",
			Description: "command help test",
		},
		{
			Name:        "help",
			Description: "Comando para visualizar os comandos que eu (Diego) suporto!",
		},
	}
)
