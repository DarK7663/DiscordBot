package handlers

import "github.com/bwmarrin/discordgo"

func handlerRoles(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "🧠 Роли",
				Description: "Выбирай роль по стэку",
				Color:       0x7f03fc,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "🧪 Scientist",
						Value:  "Исследователь и аналитик",
						Inline: true,
					},
					{
						Name:   "🔧 Contributor",
						Value:  "Контрибьютор проекта",
						Inline: true,
					},
					{
						Name:   "🤡 Vibecoder",
						Value:  "Вайб-кодер на максималках",
						Inline: false,
					},
					{
						Name:   "🕸️ Operator",
						Value:  "Оператор системы",
						Inline: false,
					},
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "🧪",
						Style:    discordgo.SuccessButton,
						CustomID: "role_scientist",
					},
					discordgo.Button{
						Label:    "🔧",
						Style:    discordgo.SecondaryButton,
						CustomID: "role_contributor",
					},
					discordgo.Button{
						Label:    "🤡",
						Style:    discordgo.DangerButton,
						CustomID: "role_vibecoder",
					},
					discordgo.Button{
						Label:    "🕸️",
						Style:    discordgo.SecondaryButton,
						CustomID: "role_operator",
					},
				},
			},
		},
	})
}
