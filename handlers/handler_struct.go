package handlers

import "github.com/bwmarrin/discordgo"

type HandlersProps struct {
	msgPingPongHanlder func(s *discordgo.Session, m *discordgo.MessageCreate)
	helpJavaHandler    func(s *discordgo.Session, m *discordgo.MessageCreate)
	msgGreeting        func(s *discordgo.Session, m *discordgo.MessageCreate)
	notifyNewMember    func(s *discordgo.Session, g *discordgo.GuildMemberAdd)
}
