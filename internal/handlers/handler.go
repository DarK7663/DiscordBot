package handlers

import (
	db "discord/internal/database"
	"discord/internal/repository"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

var repo = repository.NewUserRepository(db.DB)

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
	user, err := repo.FindCreate(m.Author.ID, m.Author.Username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error get data")
		return
	}
	text := fmt.Sprintf("Author ID:"+"`%s`\n"+"Author Name:"+"`%s\n`"+"Messages user:"+"`%d`", user.DiscordID, user.Username, user.Messages)

	s.ChannelMessageSend(m.ChannelID, text)
}
