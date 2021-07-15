package handlers

import (
	"context"
	"sort"
	"strings"
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

func TestListStockIDSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(listHandleTestSuite))
}

type listHandleTestSuite struct {
	suite.Suite
	client *mongo.Client
	handle *Handler
}

func (s *listHandleTestSuite) SetupSuite() {
	var err error
	s.client, err = test.InitDB()
	assert.NoError(s.T(), err)
	s.handle = &Handler{
		userColl: test.GetCollection(s.client, db.CollectionNameUsers),
		mailbox:  make(chan Mail, 10),
	}
}

func (s *listHandleTestSuite) TearDownSuite() {
	assert.NoError(s.T(), test.Deinit(s.client, db.CollectionNameUsers))
}

func (s *listHandleTestSuite) TearDownTest() {
	assert.NoError(s.T(), test.RemoveDocs(s.handle.userColl))
}

func (s *listHandleTestSuite) Test_getListStockHandler_noData() {
	// arrange
	userID := 5566

	// act
	s.handle.ListStock(&Mail{userID: userID, platform: TelegramBot})
	gotValue := <-s.handle.mailbox
	// assert
	assert.Equal(s.T(), "no data", gotValue.toMsg)
}

func (s *listHandleTestSuite) Test_getListStockHandler_expectStockList() {
	// arrange
	mockUser := models.User{
		UserID: 7788,
		Stocks: []string{"2330", "1101", "aapl"},
	}
	_, err := s.handle.userColl.UpdateOne(context.Background(), bson.M{"user_id": mockUser.UserID},
		bson.M{
			"$set": mockUser,
		}, options.Update().SetUpsert(true))
	assert.NoError(s.T(), err)
	sort.Strings(mockUser.Stocks)

	// act

	s.handle.ListStock(&Mail{userID: mockUser.UserID, platform: TelegramBot})
	gotValue := <-s.handle.mailbox

	// assert
	assert.Equal(s.T(), strings.Join(mockUser.Stocks, "\n"), gotValue.toMsg)
}
