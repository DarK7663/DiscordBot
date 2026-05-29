package models

// type RoleAssignment struct {
// 	UserID     string    `json:"user_id"    gorm:"primaryKey;index"`
// 	RoleID     string    `json:"role_id"    gorm:"primaryKey;index"`
// 	CustomID   string    `json:"custom_id"  gorm:"index"`
// 	GuildID    string    `json:"guild_id"   gorm:"index"`
// 	AssignedAt time.Time `json:"assigned_at"`
// }

type SelfRole struct {
	CustomID string `json:"custom_id" gorm:"primaryKey"`
	RoleID   string `json:"role_id"`
	Label    string `json:"label"`
	Emoji    string `json:"emoji,omitempty"`
	Style    int    `json:"style"`
	GuildID  string `json:"guild_id"`
}
