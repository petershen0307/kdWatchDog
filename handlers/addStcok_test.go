package handlers

import (
	"context"
	"fmt"
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

func TestAddStockIDSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(addHandleTestSuite))
}

type addHandleTestSuite struct {
	suite.Suite
	client *mongo.Client
	handle *Handler
}

func (s *addHandleTestSuite) SetupSuite() {
	var err error
	s.client, err = test.InitDB()
	assert.NoError(s.T(), err)
	s.handle = &Handler{
		userColl: test.GetCollection(s.client, db.CollectionNameUsers),
		mailbox:  make(chan Mail, 10),
	}
}

func (s *addHandleTestSuite) TearDownSuite() {
	assert.NoError(s.T(), test.Deinit(s.client, db.CollectionNameUsers))
}

func (s *addHandleTestSuite) TearDownTest() {
	assert.NoError(s.T(), test.RemoveDocs(s.handle.userColl))
}

func (s *addHandleTestSuite) Test_getAddStockHandler_oneData() {
	// arrange
	stockID := "1234"
	userID := 5566

	// act
	s.handle.AddStock(&Mail{platform: TelegramBot, fromMsg: fmt.Sprintf("/add %v", stockID), userID: userID})
	gotValue := <-s.handle.mailbox
	// assert
	assert.Equal(s.T(), fmt.Sprintf("add %v ok", stockID), gotValue.toMsg)
}

func (s *addHandleTestSuite) Test_getAddStockHandler_sameUserSecondData() {
	// arrange
	stockID := "12345"
	userID := 7788
	_, err := s.handle.userColl.UpdateOne(context.Background(), bson.M{"user_id": userID},
		bson.M{
			"$set": models.User{
				UserID: userID,
				Stocks: []string{"1111", "2222"},
			},
		}, options.Update().SetUpsert(true))
	assert.NoError(s.T(), err)

	// act
	s.handle.AddStock(&Mail{platform: TelegramBot, fromMsg: stockID, userID: userID})
	gotValue := <-s.handle.mailbox
	// assert
	assert.Equal(s.T(), fmt.Sprintf("add %v ok", stockID), gotValue.toMsg)
}

func (s *addHandleTestSuite) Test_getAddStockHandler_invalidStockID() {
	// arrange
	stockID := ""
	userID := 5566

	// act
	s.handle.AddStock(&Mail{fromMsg: stockID, userID: userID, platform: TelegramBot})
	gotValue := <-s.handle.mailbox
	// assert
	assert.Equal(s.T(), "invalid stock id", gotValue.toMsg)
}
