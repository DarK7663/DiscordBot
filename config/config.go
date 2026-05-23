package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DiscordToken  string
	GuildID       string
	CommandPrefix string
}

var db *gorm.DB

func Load() (*Config, error) {
	var err error
	if err := godotenv.Load(); err != nil {
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

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, sslMode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB.AutoMigrate(&cfg)
	return cfg, nil
}
