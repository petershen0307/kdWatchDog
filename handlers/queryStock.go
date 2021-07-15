package handlers

import (
	"bytes"
	"context"
	"log"
	"strconv"

	"github.com/petershen0307/kdWatchDog/imgur"
	"github.com/petershen0307/kdWatchDog/models"
	tableimage "github.com/petershen0307/kdWatchDog/table-image"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const queryCommand = "/query"

// GetStockMap create stock map cache
func GetStockMap(stockColl *mongo.Collection) map[string]models.StockInfo {
	// query all stock info
	cursor, err := stockColl.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("connect to db err=%v", err)
		return nil
	}
	allStock := []models.StockInfo{}
	cursor.All(context.Background(), &allStock)

	// create stock map
	stockMap := map[string]models.StockInfo{}
	for _, stock := range allStock {
		stockMap[stock.ID] = stock
	}
	return stockMap
}

// RenderOneUserOutput render one user stock info
func RenderOneUserOutput(user *models.User, stockMap map[string]models.StockInfo) *bytes.Buffer {
	ti := tableimage.Init("#fff", tableimage.PNG, "")
	ti.AddTH(
		tableimage.TR{
			BorderColor: "#000",
			Tds: []tableimage.TD{
				{
					Color: "#000",
					Text:  "stock ID",
				},
				{
					Color: "#000",
					Text:  "close",
				},
				{
					Color: "#000",
					Text:  "day K",
				},
				{
					Color: "#000",
					Text:  "day D",
				},
				{
					Color: "#000",
					Text:  "week K",
				},
				{
					Color: "#000",
					Text:  "week D",
				},
				{
					Color: "#000",
					Text:  "month K",
				},
				{
					Color: "#000",
					Text:  "month D",
				},
			},
		},
	)
	trList := []tableimage.TR{}
	for _, stockID := range user.Stocks {
		trList = append(trList, tableimage.TR{
			BorderColor: "#000",
			Tds: []tableimage.TD{
				{
					Color: "#000",
					Text:  stockID,
				},
				{
					Color: "#000",
					Text:  stockMap[stockID].DailyPrice.Close,
				},
				getTDByKDValue(stockMap[stockID].DailyKD.K),
				getTDByKDValue(stockMap[stockID].DailyKD.D),
				getTDByKDValue(stockMap[stockID].WeeklyKD.K),
				getTDByKDValue(stockMap[stockID].WeeklyKD.D),
				getTDByKDValue(stockMap[stockID].MonthlyKD.K),
				getTDByKDValue(stockMap[stockID].MonthlyKD.D),
			},
		})
	}
	ti.AddTRs(trList)
	return ti.Get()
}

func getTDByKDValue(kdValueStr string) tableimage.TD {
	kdNum, _ := strconv.ParseFloat(kdValueStr, 32)
	color := "#000"
	if kdNum <= 30.0 {
		color = "#008000"
	}
	if kdNum >= 80.0 {
		color = "#ff0000"
	}
	return tableimage.TD{Color: color, Text: kdValueStr}
}

// QueryStock handle query command
func (handle *Handler) QueryStock(mail *Mail) {
	defer func() {
		handle.mailbox <- *mail
	}()
	// query all user
	var user models.User
	if err := handle.userColl.FindOne(context.Background(), bson.M{"user_id": mail.userID}).Decode(&user); err != nil {
		log.Printf("connect to db err=%v", err)
		mail.toMsg = "connect user collection failed"
		return
	}

	stockMap := GetStockMap(handle.stockColl)
	if stockMap == nil {
		mail.toMsg = "connect user collection failed"
		return
	}
	msg := RenderOneUserOutput(&user, stockMap)
	// upload to imgur
	link, _ := imgur.UploadImage(handle.imgurClientID, msg.Bytes())
	mail.toMsg = link
}
