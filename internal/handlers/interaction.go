package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type RoleMap map[string]string

func InteractionHandler(roles RoleMap) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		customID := i.MessageComponentData().CustomID

		roleID, exists := roles[customID]
		if exists {
			err := handleRoleButton(s, i, roleID)
			if err != nil {
				fmt.Printf("Ошибка при обработке кнопки роли: %v\n", err)
			}
			return
		}

		fmt.Printf("Неизвестный CustomID кнопки: %s\n", customID)
	}
}

func handleRoleButton(s *discordgo.Session, i *discordgo.InteractionCreate, roleID string) error {
	memberID := i.Member.User.ID

	err := s.GuildMemberRoleAdd(i.GuildID, memberID, roleID)
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Роль успешно выдана!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
