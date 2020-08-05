package controllers

import (
	"cashbag-me-mini/dao"
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
	//removeOldDataTransaction()
}

func removeOldDataTransaction() {
	database.DB.Collection("transaction").DeleteMany(context.Background(), bson.M{})
}

func (s *TransactionSuite) TestCreateTransaction() {
	var (
		transaction = models.PostTransaction{
			NameCompany: "Hightland",
			NameBranch:  "Hight QuangTri",
			User: "Win",
			Amount:      100,
		}
		res = struct {
			InsertedID string `json:"InsertedID"`
		}{}
		commission  float64
		balance    float64
		loyaltyProgram  float64
		amount         float64
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions", ToIOReader(transaction))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("body", &transaction)
	loyaltyProgram = dao.GetLoyaltyProgramByCompany(transaction.NameCompany)
	balance = dao.GetBalanceByCompanyName(transaction.NameCompany)
	amount = transaction.Amount
	commission = (loyaltyProgram / 100) * amount
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	if balance > commission {
		CreateTransaction(c)
	}
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.NotEqual(s.T(), res, nil)
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
