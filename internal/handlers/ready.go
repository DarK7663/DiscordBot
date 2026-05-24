package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot name:", r.User.Username)

	s.UpdateCustomStatus("Worker bot")
}
