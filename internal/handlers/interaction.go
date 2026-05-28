package handlers

import (
	"discord/internal/database/models"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func InteractionHandler(db *gorm.DB) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		customID := i.MessageComponentData().CustomID
		guildID := i.GuildID

		var selfRole models.SelfRole
		if err := db.Where("custom_id = ? AND guild_id = ?", customID, guildID).First(&selfRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Роль не найдена.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				return
			}
			fmt.Println("Error request SelfRole:", err)
			return
		}

		err := handleRoleButton(s, i, db, &selfRole)
		if err != nil {
			fmt.Println("Ошибка toggle роли:", err)
		}
	}
}

func handleRoleButton(s *discordgo.Session, i *discordgo.InteractionCreate, db *gorm.DB, selfRole *models.SelfRole) error {
	userID := i.Member.User.ID
	guildID := i.GuildID
	roleID := selfRole.RoleID

	member, err := s.GuildMember(guildID, userID)
	if err != nil {
		return err
	}

	hasRole := false
	for _, rID := range member.Roles {
		if rID == roleID {
			hasRole = true
			break
		}
	}

	var responseText string

	if hasRole {
		err = s.GuildMemberRoleRemove(guildID, userID, roleID)
		if err != nil {
			return err
		}
		responseText = "✅ Роль успешно **снята**!"
	} else {
		err = s.GuildMemberRoleAdd(guildID, userID, roleID)
		if err != nil {
			return err
		}
		responseText = "✅ Роль успешно **выдана**!"
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: responseText,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
