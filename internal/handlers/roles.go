package handlers

import "github.com/bwmarrin/discordgo"

func handlerRoles(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: "Выбери роль:",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "🧪 Scientist",
						Style:    discordgo.SuccessButton,
						CustomID: "role_scientist",
					},
					discordgo.Button{
						Label:    "🔧Contributor",
						Style:    discordgo.SecondaryButton,
						CustomID: "role_contributor",
					},
					discordgo.Button{
						Label:    "🤡Vibecoder",
						Style:    discordgo.DangerButton,
						CustomID: "role_vibecoder",
					},
					discordgo.Button{
						Label:    "🕸️Operator",
						Style:    discordgo.SecondaryButton,
						CustomID: "role_operator",
					}, discordgo.Button{
						Label:    "🤪Council",
						Style:    discordgo.PrimaryButton,
						CustomID: "role_council",
					},
				},
			},
		},
	})
}
