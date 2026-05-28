package handlers

import (
	"discord/internal/database/models"
	"discord/internal/repository"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

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
	repo := repository.NewUserRepository(db)

	user, err := repo.FindCreate(m.Author.ID, m.Author.Username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error get data")
		return
	}
	text := fmt.Sprintf("Author ID:"+"`%s`\n"+"Author Name:"+"`%s\n`"+"Messages user:"+"`%d`", user.DiscordID, user.Username, user.Messages)

	s.ChannelMessageSend(m.ChannelID, text)
}

func handlerSetupRoles(s *discordgo.Session, m *discordgo.MessageCreate, db *gorm.DB) {
	if m.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		s.ChannelMessageSend(m.ChannelID, "❌ Нет прав администратора")
		return
	}

	roles := []models.SelfRole{
		{
			CustomID: "role_scientist",
			RoleID:   "1506652453464965180",
			Label:    "Scientist",
			GuildID:  m.GuildID,
		},
		{
			CustomID: "role_contributor",
			RoleID:   "1506652734986911814",
			Label:    "Scientist",
			GuildID:  m.GuildID,
		},
		{
			CustomID: "role_vibecoder",
			RoleID:   "1507801691637022771",
			Label:    "Scientist",
			GuildID:  m.GuildID,
		},
		{
			CustomID: "role_operator",
			RoleID:   "1506652696135073893",
			Label:    "Scientist",
			GuildID:  m.GuildID,
		},
	}

	if err := db.Create(&roles).Error; err != nil {
		s.ChannelMessageSend(m.ChannelID, "❌ Ошибка записи ролей в БД")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "✅ Роли настроены!")
}
