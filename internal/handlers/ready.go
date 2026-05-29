package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateCustomStatus("Worker bot")
}
