package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type TransactionSuite struct {
	suite.Suite
	transactions []models.TransactionBSON
}

func (s TransactionSuite) SetupSuite() {
	database.Connectdb("CashBag-test")
}
func (s TransactionSuite) TearDownSuite() {
	removeOldDataTransaction()
}

func removeOldDataTransaction() {
	database.DB.Collection("transaction").DeleteMany(context.Background(), bson.M{})
}

func (s *TransactionSuite) TestCreateTransaction() {
	var (
		transaction = models.PostTransaction{
			NameCompany: "Highland",
			NameBranch:  "Hight QuangTri",
			User:        "Mars",
			Amount:      100,
		}
		res = struct {
			InsertedID string `json:"InsertedID"`
		}{}
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions", ToIOReader(transaction))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("body", &transaction)
	CreateTransaction(c)
	assert.Equal(s.T(), http.StatusCreated, rec.Code)
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.NotEqual(s.T(), res, nil)
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}