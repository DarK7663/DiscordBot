package repository

import (
	"discord/internal/database/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindCreate(discordID, username string) (*models.User, error) {
	user := &models.User{}

	result := r.DB.Where("discord_id = ?", discordID).FirstOrCreate(user, models.User{
		DiscordID: discordID,
		Username:  username,
	})
	return user, result.Error
}

func (r *UserRepository) IncrementMessages(discordID string) error {
	if err := r.DB.Model(&models.User{}).Where("discord_id = ?", discordID).UpdateColumn("messages", gorm.Expr("messages + 1")).Error; err != nil {
		return fmt.Errorf("error: %s", err)
	}
	return nil
}
