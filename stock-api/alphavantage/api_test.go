package alphavantage

import (
	"testing"

	stockapi "github.com/petershen0307/kdWatchDog/stock-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestAlphavantageAPISuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(alphavantageAPITestSuite))
}

type alphavantageAPITestSuite struct {
	suite.Suite
	api *API
}

func (s *alphavantageAPITestSuite) SetupSuite() {
	s.api = New("...")
}

func (s *alphavantageAPITestSuite) TearDownSuite() {
}

func (s *alphavantageAPITestSuite) TearDownTest() {
}

func (s *alphavantageAPITestSuite) Test_getStockSymbol_returnEmpty() {
	symbolList := s.api.GetStockSymbol("AAPL")
	assert.Empty(s.T(), symbolList, "no implement getStockSymbol in alphavantage, should be empty list")
}

func (s *alphavantageAPITestSuite) Test_getSTOCH_returnValue() {
	stockSTOCH := s.api.GetSTOCH("AAPL", stockapi.Daily, 9, 3, 3, 0, 0)
	assert.NotEmpty(s.T(), stockSTOCH.K)
	assert.NotEmpty(s.T(), stockSTOCH.D)
}

func (s *alphavantageAPITestSuite) Test_GetDailyPrice_returnValue() {
	price := s.api.GetDailyPrice("AAPL")
	assert.NotEmpty(s.T(), price.Close)
	assert.NotEmpty(s.T(), price.High)
	assert.NotEmpty(s.T(), price.Low)
	assert.NotEmpty(s.T(), price.Open)
}

func (s *alphavantageAPITestSuite) Test_GetWeeklyPrice_returnValue() {
	price := s.api.GetWeeklyPrice("AAPL")
	assert.NotEmpty(s.T(), price.Close)
	assert.NotEmpty(s.T(), price.High)
	assert.NotEmpty(s.T(), price.Low)
	assert.NotEmpty(s.T(), price.Open)
}

func (s *alphavantageAPITestSuite) Test_GetMonthlyPrice_returnValue() {
	price := s.api.GetMonthlyPrice("AAPL")
	assert.NotEmpty(s.T(), price.Close)
	assert.NotEmpty(s.T(), price.High)
	assert.NotEmpty(s.T(), price.Low)
	assert.NotEmpty(s.T(), price.Open)
}
