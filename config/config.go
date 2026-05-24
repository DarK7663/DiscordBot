package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken  string
	GuildID       string
	CommandPrefix string
}

func Load() (*Config, error) {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system env vars")
	}

	cfg := &Config{
		DiscordToken:  os.Getenv("DISCORD_TOKEN"),
		GuildID:       os.Getenv("GUILD_ID"),
		CommandPrefix: os.Getenv("COMMAND_PREFIX"),
	}

	if cfg.DiscordToken == "" {
		return nil, fmt.Errorf("DISCORD_TOKEN не задан в .env файле")
	}

	if cfg.CommandPrefix == "" {
		cfg.CommandPrefix = "!"
	}

	return cfg, nil
}
