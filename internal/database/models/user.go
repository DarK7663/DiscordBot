package models

import "time"

type User struct {
	DiscordID string    `gorm:"primaryKey;not null" json:"discord_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdat"`
	Messages  int32     `json:"messages"`
}
