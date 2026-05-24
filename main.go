package main

import (
	"discord/config"
	"discord/internal/bot"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config error:", err)
	}
	b, err := bot.New(cfg)
	if err != nil {
		log.Fatal("Bot error:", err)
	}
	if err := b.Start(); err != nil {
		log.Fatal("Start error:", err)
	}
}
