package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func MessageHandler(prefix string, db *gorm.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		if _, err := repo.FindCreate(m.Author.ID, m.Author.Username); err != nil {
			fmt.Printf("User not found in DB")
			return
		}
		if err := repo.IncrementMessages(m.Author.ID); err != nil {
			fmt.Printf("Failed to increment message count: %v", err)
			return
		}

		content := strings.TrimSpace(m.Content)
		if !strings.HasPrefix(content, prefix) {
			return
		}

		args := strings.Fields(strings.TrimPrefix(m.Content, prefix))
		if len(args) == 0 {
			return
		}
		command := strings.ToLower(args[0])

		switch command {
		case "ping":
			handlerPing(s, m)
		case "help":
			handlerHelp(s, m, prefix)
		case "info":
			handlerInfo(s, m)
		case "profile":
			handlerProfile(s, m, db)
		case "role":
			handlerRoles(s, m)
		}
	}
}
