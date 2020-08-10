package controllers

import (
	"cashbag-me-mini/config"
	"cashbag-me-mini/modules/redis"
	"cashbag-me-mini/modules/zookeeper"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/ultis"
)

type TransactionSuite struct {
	suite.Suite
}

func (s TransactionSuite) SetupSuite() {
	var cfg = config.GetEnv()
	database.Connect(cfg.DatabaseTestName)
	redis.Connect()
	zookeeper.Connect(cfg.ZookeeperTestURI)
	ultis.HelperCompanyCreateFake()
	ultis.HelperBranchCreateFake()
}

func (s TransactionSuite) TearDownSuite() {
	removeOldDataTransaction()
}
func removeOldDataTransaction() {
	database.BranchCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
	database.TransactionCol().DeleteMany(context.Background(), bson.M{})
	database.TransactionAnalyticCol().DeleteMany(context.Background(), bson.M{})
}

func (s *TransactionSuite) TestTransactionCreateSuccess() {
	var (
		companyID   = ultis.CompanyID
		branchID    = ultis.BranchID
		transaction = models.TransactionCreatePayload{
			CompanyID: companyID,
			BranchID:  branchID,
			User:      "hoang",
			Amount:    10000,
		}
		response ultis.Response
	)

	// Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", &transaction)

	// Call TransactionCreate
	TransactionCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	// Get totalTransaction
	transactionAnalytic := ultis.HelperTransactionAnalyticFindByID()
	totalTransaction := transactionAnalytic.TotalTransaction

	//Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
	assert.Equal(s.T(), 1, totalTransaction)
}

func (s *TransactionSuite) TestTransactionCreateFailBecauseCompanyID() {
	var (
		companyID   = "5f24d45125ea51bc57a8285"
		branchID    = ultis.BranchID
		transaction = models.TransactionCreatePayload{
			CompanyID: companyID,
			BranchID:  branchID,
			User:      "hoang",
			Amount:    10000,
		}
		response ultis.Response
	)

	// Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", &transaction)

	// create transaction
	TransactionCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	//Test
	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Khong tim thay Cong Ty ", response["message"])
}

func (s *TransactionSuite) TestTransactionCreateFailBecauseBranchID() {
	var (
		companyID   = ultis.CompanyID
		branchID    = "5f24d45125ea51bc57a8285"
		transaction = models.TransactionCreatePayload{
			CompanyID: companyID,
			BranchID:  branchID,
			User:      "hoang",
			Amount:    10000,
		}
		response ultis.Response
	)

	// Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", &transaction)

	// Call TransactionCreate
	TransactionCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	//Test
	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Khong tim thay Chi Nhanh", response["message"])
}

func (s *TransactionSuite) TestTransactionCreateFailBecauseUser() {
	var (
		companyID   = ultis.CompanyID
		branchID    = ultis.BranchID
		transaction = models.TransactionCreatePayload{
			CompanyID: companyID,
			BranchID:  branchID,
			User:      "",
			Amount:    10000,
		}
		response ultis.Response
	)

	// Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", &transaction)

	// Call TransactionCreate
	TransactionCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	//Test
	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(s.T(), nil, response["data"])
	assert.Equal(s.T(), "User khong nam trong danh sach hoan tien", response["message"])
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
