package handlers

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/petershen0307/kdWatchDog/db"
	"github.com/petershen0307/kdWatchDog/models"
	"github.com/petershen0307/kdWatchDog/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func TestDelStockIDSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(delHandleTestSuite))
}

type delHandleTestSuite struct {
	suite.Suite
	client *mongo.Client
	handle *Handler
}

func (s *delHandleTestSuite) SetupSuite() {
	var err error
	s.client, err = test.InitDB()
	assert.NoError(s.T(), err)
	s.handle = &Handler{
		userColl: test.GetCollection(s.client, db.CollectionNameUsers),
		mailbox:  make(chan Mail, 10),
	}
}

func (s *delHandleTestSuite) TearDownSuite() {
	assert.NoError(s.T(), test.Deinit(s.client, db.CollectionNameUsers))
}

func (s *delHandleTestSuite) TearDownTest() {
	assert.NoError(s.T(), test.RemoveDocs(s.handle.userColl))
}

func (s *delHandleTestSuite) Test_getDelStockHandler_delNone() {
	// arrange
	stockID := "1234"
	userID := 5566

	// act
	s.handle.DelStock(&Mail{
		fromMsg:  fmt.Sprintf("/del %v", stockID),
		userID:   userID,
		platform: TelegramBot,
	})
	gotValue := <-s.handle.mailbox
	// assert
	assert.Equal(s.T(), "no record", gotValue.toMsg)
}

func (s *delHandleTestSuite) Test_getDelStockHandler_delOne() {
	// arrange
	mockUser := models.User{
		BotPlatform: int(TelegramBot),
		UserID:      7788,
		Stocks:      []string{"1101", "2330", "aapl"},
	}
	_, err := s.handle.userColl.UpdateOne(context.Background(), bson.M{"user_id": mockUser.UserID},
		bson.M{
			"$set": mockUser,
		}, options.Update().SetUpsert(true))
	assert.NoError(s.T(), err)

	// act
	s.handle.DelStock(&Mail{
		fromMsg:  fmt.Sprintf("/del %v", 1101),
		userID:   mockUser.UserID,
		platform: TelegramBot,
	})
	gotValue := <-s.handle.mailbox

	// assert
	actualResult := models.User{}
	err = s.handle.userColl.FindOne(context.Background(), bson.M{"user_id": mockUser.UserID}).Decode(&actualResult)
	sort.Strings(actualResult.Stocks)
	expectResult := models.User{
		UserID:     7788,
		Stocks:     []string{"2330", "aapl"},
		LastUpdate: actualResult.LastUpdate,
	}
	sort.Strings(expectResult.Stocks)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectResult, actualResult)
	assert.Equal(s.T(), "success", gotValue.toMsg)
}
