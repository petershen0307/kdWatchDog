package models

import "time"

// StockInfo stored stock price
type StockInfo struct {
	ID         string    `bson:"stock_id"`
	DailyKD    STOCH     `bson:"daily_kd,omitempty"`
	WeeklyKD   STOCH     `bson:"weekly_kd,omitempty"`
	MonthlyKD  STOCH     `bson:"monthly_kd,omitempty"`
	DailyPrice Price     `bson:"daily_price,omitempty"`
	LastUpdate time.Time `bson:"last_update,omitempty"`
}

// STOCH is kd indicator
type STOCH struct {
	K string `bson:"k,omitempty"`
	D string `bson:"d,omitempty"`
}

// Price is stock price
type Price struct {
	Open  string `bson:"open,omitempty"`
	High  string `bson:"high,omitempty"`
	Low   string `bson:"low,omitempty"`
	Close string `bson:"close,omitempty"`
}
