package imgur

import (
	"testing"

	"github.com/petershen0307/kdWatchDog/handlers"
	"github.com/petershen0307/kdWatchDog/models"
	"github.com/stretchr/testify/suite"
)

func TestAPISuiteSuite(t *testing.T) {
	suite.Run(t, new(apiSuite))
}

type apiSuite struct {
	suite.Suite
	clientID string
	user     models.User
	stockMap map[string]models.StockInfo
}

func (s *apiSuite) SetupSuite() {
	s.clientID = ""
	s.user = models.User{
		Stocks: []string{"AAPL", "DDOG", "LMT"},
	}
	s.stockMap = map[string]models.StockInfo{
		"AAPL": {
			DailyPrice: models.Price{Close: "100"},
			DailyKD:    models.STOCH{K: "1", D: "2"},
			WeeklyKD:   models.STOCH{K: "3", D: "4"},
			MonthlyKD:  models.STOCH{K: "5", D: "6"},
		},
		"DDOG": {
			DailyPrice: models.Price{Close: "200"},
			DailyKD:    models.STOCH{K: "12", D: "22"},
			WeeklyKD:   models.STOCH{K: "32", D: "42"},
			MonthlyKD:  models.STOCH{K: "52", D: "62"},
		},
		"LMT": {
			DailyPrice: models.Price{Close: "300"},
			DailyKD:    models.STOCH{K: "13", D: "23"},
			WeeklyKD:   models.STOCH{K: "33", D: "43"},
			MonthlyKD:  models.STOCH{K: "53", D: "63"},
		},
	}
}

func (s *apiSuite) TearDownSuite() {

}

func (s *apiSuite) Test_UpdloadImage_success() {
	image := handlers.RenderOneUserOutput(&s.user, s.stockMap)
	imageLink, err := UploadImage(s.clientID, image.Bytes())
	s.NoError(err)
	s.NotEmpty(imageLink)
}
