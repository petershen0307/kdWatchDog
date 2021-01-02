package models

// User is the db structure for telegram user
type User struct {
	UserID int      `bson:"user_id"`
	Stocks []string `bson:"stocks"`
}
