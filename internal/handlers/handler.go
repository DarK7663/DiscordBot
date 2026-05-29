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

	embed := &discordgo.MessageEmbed{
		Title:       "👨🏻‍🦯‍➡️ Навигация по командам",
		Description: "Таблица префикс-команд",
		Color:       0x00FF00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "🤖 Команды",
				Value: fmt.Sprintf(
					"`%sping` — проверить бота\n"+
						"`%shelp` — список команд\n"+
						"`%sinfo` — информация о сервере\n"+
						"`%sprofile` — твой профиль\n"+
						"`%sroles` — выбор роли",
					prefix, prefix, prefix, prefix, prefix,
				),
				Inline: false,
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func handlerInfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, err := s.GuildWithCounts(m.GuildID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "👨🏼‍🦽‍➡️ " + guild.Name,
		Description: "Информация о Сервере",
		Color:       0x00CED1,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🆔 Discord ID",
				Value:  guild.ID,
				Inline: true,
			},
			{
				Name:   "🧑‍🧑‍🧒‍🧒 Members",
				Value:  fmt.Sprintf("%d", guild.ApproximateMemberCount),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Chugyn Bot",
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func handlerProfile(s *discordgo.Session, m *discordgo.MessageCreate, db *gorm.DB) {
	repo := repository.NewUserRepository(db)

	user, err := repo.FindCreate(m.Author.ID, m.Author.Username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error get data")
		return
	}
	embed := &discordgo.MessageEmbed{
		Title:       "👤 " + user.Username,
		Description: "Информация о пользователе",
		Color:       0x5865F2,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🆔 Discord ID",
				Value:  user.DiscordID,
				Inline: true,
			},
			{
				Name:   "💬 Сообщений",
				Value:  fmt.Sprintf("%d", user.Messages),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Chugyn Bot",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
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
