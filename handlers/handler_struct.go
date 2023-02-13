package handlers

import "github.com/bwmarrin/discordgo"

type HandlersProps struct {
	msgPingPongHanlder func(s *discordgo.Session, m *discordgo.MessageCreate)
	ImgGenerator       func(s *discordgo.Session, m *discordgo.MessageCreate)
	helpJavaHandler    func(s *discordgo.Session, m *discordgo.MessageCreate)
	msgGreeting        func(s *discordgo.Session, m *discordgo.MessageCreate)
	notifyNewMember    func(s *discordgo.Session, g *discordgo.GuildMemberAdd)
}

type SrcProps struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type PhotoProps struct {
	ID              int       `json:"id"`
	Width           int       `json:"width"`
	Height          int       `json:"height"`
	URL             string    `json:"url"`
	Photographer    string    `json:"photographer"`
	PhotographerURL string    `json:"photographer_url"`
	PhotographerID  int       `json:"photographer_id"`
	AvgColor        string    `json:"avg_color"`
	Src             *SrcProps `json:"Src"`
	Liked           bool      `json:"liked"`
	Alt             string    `json:"alt"`
}

type AiResponse struct {
	Page          int           `json:"page"`
	Per_Page      int           `json:"per_page"`
	Photos        []*PhotoProps `json:"Photos"`
	Total_Results int           `json:"total_results"`
	Next_Page     string        `json:"next_page"`
	Prev_Page     string        `json:"prev_page"`
}
