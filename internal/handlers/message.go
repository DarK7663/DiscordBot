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

		args := strings.Fields(m.Content)
		if len(args) == 0 {
			return
		}
		command := strings.ToLower(args[0])

		switch command {
		case "!ping":
			handlerPing(s, m)
		case "!help":
			handlerHelp(s, m, prefix)
		case "!info":
			handlerInfo(s, m)
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

func handlerInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, err := s.Guild(m.GuildID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	info := fmt.Sprintf("Guild Name:"+"`%s`\n"+"Guild members:"+"`%d`", guild.Name, guild.MemberCount)

	s.ChannelMessageSend(m.ChannelID, info)
}
