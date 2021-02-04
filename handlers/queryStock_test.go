package handlers

import (
	"io/ioutil"
	"testing"

	"github.com/petershen0307/kdWatchDog/models"
	"github.com/stretchr/testify/suite"
)

func TestRenderSuite(t *testing.T) {
	t.Skip()
	suite.Run(t, new(renderSuite))
}

type renderSuite struct {
	suite.Suite
	clientID string
	user     models.User
	stockMap map[string]models.StockInfo
}

func (s *renderSuite) SetupSuite() {
	s.clientID = ""
	s.user = models.User{
		Stocks: []string{"AAPL", "DDOG", "LMT"},
	}
	s.stockMap = map[string]models.StockInfo{
		"AAPL": {
			DailyPrice: models.Price{Close: "100.0000"},
			DailyKD:    models.STOCH{K: "1", D: "2"},
			WeeklyKD:   models.STOCH{K: "39", D: "94"},
			MonthlyKD:  models.STOCH{K: "95", D: "6"},
		},
		"DDOG": {
			DailyPrice: models.Price{Close: "2000.0000"},
			DailyKD:    models.STOCH{K: "12", D: "22"},
			WeeklyKD:   models.STOCH{K: "82", D: "42"},
			MonthlyKD:  models.STOCH{K: "52", D: "62"},
		},
		"LMT": {
			DailyPrice: models.Price{Close: "30000.0000"},
			DailyKD:    models.STOCH{K: "13", D: "23"},
			WeeklyKD:   models.STOCH{K: "33", D: "43"},
			MonthlyKD:  models.STOCH{K: "53", D: "63"},
		},
	}
}

func (s *renderSuite) TearDownSuite() {

}

func (s *renderSuite) Test_UpdloadImage_success() {
	image := RenderOneUserOutput(&s.user, s.stockMap)
	ioutil.WriteFile("test.png", image.Bytes(), 0755)
}
