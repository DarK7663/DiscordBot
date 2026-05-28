package bot

import (
	"discord/config"
	"discord/internal/handlers"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type Bot struct {
	Session *discordgo.Session
	Config  *config.Config
	DB      *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) (*Bot, error) {
	session, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания сессии: %w", err)
	}

	session.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentMessageContent |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuildPresences

	var bot = &Bot{
		Session: session,
		Config:  cfg,
		DB:      db,
	}
	bot.Session.AddHandler(handlers.ReadyHandler)
	bot.Session.AddHandler(handlers.MessageHandler(cfg.CommandPrefix, db))
	bot.Session.AddHandler(handlers.InteractionHandler(db))

	return bot, nil
}

func (b *Bot) Start() error {
	if err := b.Session.Open(); err != nil {
		return fmt.Errorf("Start func error: %w", err)
	}
	fmt.Println("Bot started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	fmt.Println("Bot stopped")

	b.Session.Close()

	return nil
}
