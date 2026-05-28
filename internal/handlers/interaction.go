package handlers

import "github.com/bwmarrin/discordgo"

func InteractionHandler(roles map[string]string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}
}
