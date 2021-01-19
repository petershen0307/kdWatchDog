package handlers

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/petershen0307/kdWatchDog/models"
	"github.com/petershen0307/kdWatchDog/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	tg "gopkg.in/tucnak/telebot.v2"
)

func TestListStockIDSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(listHandleTestSuite))
}

type listHandleTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *listHandleTestSuite) SetupSuite() {
	var err error
	s.client, err = test.InitDB()
	assert.NoError(s.T(), err)
	s.collection = test.GetCollection(s.client, "users")
}

func (s *listHandleTestSuite) TearDownSuite() {
	assert.NoError(s.T(), test.Deinit(s.client, "users"))
}

func (s *listHandleTestSuite) TearDownTest() {
	assert.NoError(s.T(), test.RemoveDocs(s.collection))
}

func (s *listHandleTestSuite) Test_getListStockHandler_noData() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	userID := 5566

	// act
	command, f := getListStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: userID,
		},
	})

	// assert
	assert.Equal(s.T(), listCommand, command)
	assert.Equal(s.T(), gotValue, "no data")
}

func (s *listHandleTestSuite) Test_getListStockHandler_expectStockList() {
	// arrange
	gotValue := ""
	responseCallback := func(p *post) {
		gotValue = p.what.(string)
	}
	mockUser := models.User{
		UserID: 7788,
		Stocks: []string{"2330", "1101", "aapl"},
	}
	_, err := s.collection.UpdateOne(context.Background(), bson.M{"user_id": mockUser.UserID},
		bson.M{
			"$set": mockUser,
		}, options.Update().SetUpsert(true))
	assert.NoError(s.T(), err)
	sort.Strings(mockUser.Stocks)

	// act
	command, f := getListStockHandler(responseCallback, s.collection)
	f(&tg.Message{
		Sender: &tg.User{
			ID: mockUser.UserID,
		},
	})

	// assert
	assert.Equal(s.T(), listCommand, command)
	assert.Equal(s.T(), strings.Join(mockUser.Stocks, "\n"), gotValue)
}
