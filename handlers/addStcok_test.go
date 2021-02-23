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
	tg "gopkg.in/tucnak/telebot.v2"
)

func TestAddStockIDSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(addHandleTestSuite))
}

type addHandleTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *addHandleTestSuite) SetupSuite() {
	var err error
	s.client, err = test.InitDB()
	assert.NoError(s.T(), err)
	s.collection = test.GetCollection(s.client, db.CollectionNameUsers)
}

func (s *addHandleTestSuite) TearDownSuite() {
	assert.NoError(s.T(), test.Deinit(s.client, db.CollectionNameUsers))
}

func (s *addHandleTestSuite) TearDownTest() {
	assert.NoError(s.T(), test.RemoveDocs(s.collection))
}

func (s *addHandleTestSuite) Test_getAddStockHandler_oneData() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	stockID := "1234"
	userID := 5566

	// act
	command, f := getAddStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: userID,
		},
		Text: fmt.Sprintf("/add %v", stockID),
	})

	// assert
	assert.Equal(s.T(), addCommand, command)
	assert.Equal(s.T(), fmt.Sprintf("add %v ok", stockID), gotValue)
}

func (s *addHandleTestSuite) Test_getAddStockHandler_sameUserSecondData() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	stockID := "12345"
	userID := 7788
	_, err := s.collection.UpdateOne(context.Background(), bson.M{"user_id": userID},
		bson.M{
			"$set": models.User{
				UserID: userID,
				Stocks: []string{"1111", "2222"},
			},
		}, options.Update().SetUpsert(true))
	assert.NoError(s.T(), err)

	// act
	command, f := getAddStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: userID,
		},
		Text: fmt.Sprintf("/add %v", stockID),
	})

	// assert
	assert.Equal(s.T(), addCommand, command)
	assert.Equal(s.T(), fmt.Sprintf("add %v ok", stockID), gotValue)
}

func (s *addHandleTestSuite) Test_getAddStockHandler_invalidStockID() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	stockID := ""
	userID := 5566

	// act
	command, f := getAddStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: userID,
		},
		Text: fmt.Sprintf("/add %v", stockID),
	})

	// assert
	assert.Equal(s.T(), addCommand, command)
	assert.Equal(s.T(), "invalid stock id", gotValue)
}
