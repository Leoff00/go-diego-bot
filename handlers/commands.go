package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "help_bombas",
			Description: "irei sugerir pessoas nos cargos para te ajudar (opções suportadas: java/js)",
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
	**/oi_diego -> responderei você de volta!**
	**/img [parametros] -> gerarei pra você uma imagem com o dado que você me forneceu!**
	**/help_bombas [linguagem] use este comando quando tiver uma duvida em determinada linguagem!**
	**/ping ou /pong -> jogarei um ping pong com você :)!**
	**/clear [limite] (limite de 100) para excluir as mensagens do chat (usado apenas por administradores).** 
	**/people para mostrar quantas pessoas temos no servidor!** 
	**/avatar para mostrar a foto de alguem no servidor.** 
	`
)
