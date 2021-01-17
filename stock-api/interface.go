package stockapi

import "github.com/petershen0307/kdWatchDog/models"

type stockAPI interface {
	Wait()
	GetStockSymbol(exchange string) []string
	GetSTOCH(symbol string, interval ResolutionInterval, fastkperiod, slowkperiod, slowdperiod, slowkmatype, slowdmatype uint8) models.STOCH
	GetDailyPrice(symbol string) models.Price
	GetWeeklyPrice(symbol string) models.Price
	GetMonthlyPrice(symbol string) models.Price
}

// ResolutionInterval is the inteval definition
type ResolutionInterval int

const (
	// OneMin 1min
	OneMin ResolutionInterval = iota
	// FiveMin 5min
	FiveMin
	// FifthteenMin 15min
	FifthteenMin
	// ThirtyMin 30min
	ThirtyMin
	// SixtyMin 60min
	SixtyMin
	// Daily 1day
	Daily
	// Weekly 1week
	Weekly
	// Monthly 1month
	Monthly
)
