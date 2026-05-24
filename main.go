package main

import (
	"discord/config"
	"discord/internal/bot"
	db "discord/internal/database"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config error:", err)
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal("DB error:", err)
	}

	b, err := bot.New(cfg, database)
	if err != nil {
		log.Fatal("Bot error:", err)
	}
	if err := b.Start(); err != nil {
		log.Fatal("Start error:", err)
	}
}
