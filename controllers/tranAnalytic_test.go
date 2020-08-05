package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/services"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TranAnalyticSuite ...
type TranAnalyticSuite struct {
	suite.Suite
	tranAnalytics []models.TranAnalyticBSON
}

// SetupSuite ...
func (s TranAnalyticSuite) SetupSuite() {
	database.Connectdb("CashBag-test")
}

//TearDownSuite ...
func (s TranAnalyticSuite) TearDownSuite() {
}

//TestTranAnalytic ...
func (s *TranAnalyticSuite) TestTranAnalytic() {
	var (
		TranAnalytics []models.TranAnalyticDetail
		res           []models.TranAnalyticDetail
		date = "2020-08-04"
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tranAnalytic/?date="+date, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	TranAnalytic(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	TranAnalytics = services.TranAnalytic(date)
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(s.T(), TranAnalytics, res)
}

//TestTranAnalyticSuite ...

func TestTranAnalyticSuite(t *testing.T) {
	suite.Run(t, new(TranAnalyticSuite))
}
