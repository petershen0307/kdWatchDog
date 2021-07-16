package models

import "time"

// User is the db structure for telegram user
type User struct {
	UserID      int       `bson:"user_id,omitempty"`
	BotPlatform int       `bson:"bot_platform,omitempty"`
	Stocks      []string  `bson:"stocks,omitempty"`
	LastUpdate  time.Time `bson:"last_update,omitempty"`
}
