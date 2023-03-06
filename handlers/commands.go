package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
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
		},
	}

	HelpStr1 = `Iaee meu nome é Die**go**, bot em go feito pra te ajudar com algumas
utilidades no server esses são os comandos pelo qual eu respondo: `
	HelpStr2 = `
	**oi diego -> responderei você de volta!**
	**!picture [parametros] -> gerarei pra você uma imagem com o dado que você me forneceu!**
	**!java [mensagem] -> marcarei 3 pessoas que manjam de java no server para te ajudar!**
	**!ping ou !pong -> jogarei um ping pong com você :)!**
	**!clear [quantidade] (limite de 100) para excluir as mensagens do chat.** 
	**/people para mostrar quantas pessoas temos no servidor!** 
	**/avatar para mostrar sua foto, contendo tambem a url dela.** 
	`
)
