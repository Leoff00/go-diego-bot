package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "help_bombas",
			Description: "selecione 3 pessoas para te ajudar",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "language",
					Description: "Qual a linguagem?",
				},
			},
		},
		{
			Name:        "oi_diego",
			Description: "Saudações!!",
		},
		{
			Name:        "clear",
			Description: "exclua até 100 mensagens",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "limit",
					Description: "limite de mensagens fornecido",
					Required:    true,
					MaxLength:   100,
				},
			},
		},
		{
			Name:        "ping",
			Description: "ping",
		},
		{

			Name:        "pong",
			Description: "pong",
		},
		{
			Name:        "img",
			Description: "| Comando para buscar uma imagem",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Nome da imagem que voce deseja buscar",
					Required:    true,
				},
			},
		},
		{
			Name:        "people",
			Description: "| Comando para listar quantas pessoas temos no servidor",
		},
		{
			Name:        "help",
			Description: "| Comando para visualizar os comandos que eu (Diego) suporto!",
		},
		{
			Name:        "avatar",
			Description: "| Veja sua linda foto 📸.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "user",
					Description: "usuario de busca",
					Required:    true,
				},
			},
		},
	}

	HelpStr1 = `Iaee meu nome é Die**go**, bot em go feito pra te ajudar com algumas
utilidades no server esses são os comandos pelo qual eu respondo: `
	HelpStr2 = `
	**oi diego -> responderei você de volta!**
	**/img [parametros] -> gerarei pra você uma imagem com o dado que você me forneceu!**
	**!java [mensagem] -> marcarei 3 pessoas que manjam de java no server para te ajudar!**
	**/ping ou /pong -> jogarei um ping pong com você :)!**
	**!clear [quantidade] (limite de 100) para excluir as mensagens do chat.** 
	**/people para mostrar quantas pessoas temos no servidor!** 
	**/avatar para mostrar sua foto, contendo tambem a url dela.** 
	`
)
