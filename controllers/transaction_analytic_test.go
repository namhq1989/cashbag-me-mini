package controllers

import (
	"cashbag-me-mini/modules/zookeeper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"net/http"
	"net/http/httptest"
	"testing"

	"cashbag-me-mini/util"
	"encoding/json"
)

// TransactionAnalyticSuite ...
type TransactionAnalyticSuite struct {
	suite.Suite
}

// SetupSuite ...
func (s TransactionAnalyticSuite) SetupSuite() {
	zookeeper.Connect()
	util.HelperConnect()
}

// TearDownSuite ...
func (s TransactionAnalyticSuite) TearDownSuite() {
}

// TestTransactionAnalytic ...
func (s *TransactionAnalyticSuite) TestTransactionAnalyticList() {
	var (
		response util.Response
		date     = "2020-08-04"
	)

	// Create Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tranAnalytic/?date="+date, nil)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)

	// Call TransactionAnalyticList
	TransactionAnalyticList(c)

	// Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	// Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

func TestTransactionAnalyticSuite(t *testing.T) {
	suite.Run(t, new(TransactionAnalyticSuite))
}
