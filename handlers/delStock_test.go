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
	tg "gopkg.in/tucnak/telebot.v2"
)

func TestDelStockIDSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(delHandleTestSuite))
}

type delHandleTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *delHandleTestSuite) SetupSuite() {
	var err error
	s.client, err = test.InitDB()
	assert.NoError(s.T(), err)
	s.collection = test.GetCollection(s.client, db.CollectionNameUsers)
}

func (s *delHandleTestSuite) TearDownSuite() {
	assert.NoError(s.T(), test.Deinit(s.client, db.CollectionNameUsers))
}

func (s *delHandleTestSuite) TearDownTest() {
	assert.NoError(s.T(), test.RemoveDocs(s.collection))
}

func (s *delHandleTestSuite) Test_getDelStockHandler_delNone() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	stockID := "1234"
	userID := 5566

	// act
	command, f := getDelStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: userID,
		},
		Text: fmt.Sprintf("/del %v", stockID),
	})

	// assert
	assert.Equal(s.T(), delCommand, command)
	assert.Equal(s.T(), "no record", gotValue)
}

func (s *delHandleTestSuite) Test_getDelStockHandler_delOne() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	mockUser := models.User{
		UserID: 7788,
		Stocks: []string{"1101", "2330", "aapl"},
	}
	_, err := s.collection.UpdateOne(context.Background(), bson.M{"user_id": mockUser.UserID},
		bson.M{
			"$set": mockUser,
		}, options.Update().SetUpsert(true))
	assert.NoError(s.T(), err)

	// act
	command, f := getDelStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: mockUser.UserID,
		},
		Text: "/del 1101",
	})

	// assert
	actualResult := models.User{}
	err = s.collection.FindOne(context.Background(), bson.M{"user_id": mockUser.UserID}).Decode(&actualResult)
	sort.Strings(actualResult.Stocks)
	expectResult := models.User{
		UserID:     7788,
		Stocks:     []string{"2330", "aapl"},
		LastUpdate: actualResult.LastUpdate,
	}
	sort.Strings(expectResult.Stocks)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectResult, actualResult)
	assert.Equal(s.T(), delCommand, command)
	assert.Equal(s.T(), "success", gotValue)
}
