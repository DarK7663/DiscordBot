package handlers

import (
	"discord/internal/repository"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func MessageHandler(prefix string, db *gorm.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		repo := repository.NewUserRepository(db) // хуйня
		if m.Author.Bot {
			return
		}

		if err := repo.IncrementMessages(m.Author.ID); err != nil {
			log.Printf("Failed to increment message count: %v", err)
		}

		content := strings.TrimSpace(m.Content)
		if !strings.HasPrefix(content, prefix) {
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
		case "info":
			handlerInfo(s, m)
		case "profile":
			handlerProfile(s, m, db)
		}
	}
}

func handlerPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Яйца жаренные именно")
}

func handlerHelp(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	text := fmt.Sprintf("Список команд:\n"+"`%sping` - проверить бота\n"+"`%shelp` - список команд\n"+"`%sinfo`"+" - Информация сервера\n"+"`%sprofile`"+" - Профиль", prefix, prefix, prefix, prefix)
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

func handlerProfile(s *discordgo.Session, m *discordgo.MessageCreate, db *gorm.DB) {
	repo := repository.NewUserRepository(db) // хуйня
	user, err := repo.FindCreate(m.Author.ID, m.Author.Username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error get data")
		return
	}
	text := fmt.Sprintf("Author ID:"+"`%s`\n"+"Author Name:"+"`%s\n`"+"Messages user:"+"`%d`", user.DiscordID, user.Username, user.Messages)

	s.ChannelMessageSend(m.ChannelID, text)
}

func handlerRoles(s *discordgo.Session, m *discordgo.MessageCreate) {

}
