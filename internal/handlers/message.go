package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(prefix string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		args := strings.Fields(strings.TrimPrefix(m.Content, prefix))
		if len(args) == 0 {
			return
		}

		command := strings.ToLower(args[0])

		switch command {
		case "ping":
			handlerPing(s, m)
		case "help":
			handlerHelp(s, m, prefix)
		}
	}
}

func handlerPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Яйца жаренные именно")
}

func handlerHelp(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	text := fmt.Sprintf("Список команд:\n"+"`%sping` - проверить бота\n"+"`%shelp` - список команд", prefix, prefix)
	s.ChannelMessageSend(m.ChannelID, text)
}
