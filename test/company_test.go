package test

import (
	"bytes"
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type CreateModelSuite struct {
	suite.Suite
	companies []models.CompanyDetail
}

func (s CreateModelSuite) SetupSuite() {
	database.Connectdb("CashBag-test")
	removeOldData()

}

func (s CreateModelSuite) TearDownSuite() {
	//removeOldData()
}
func removeOldData() {
	database.ConnectCol("Companies").DeleteMany(context.Background(), bson.M{})
}

func (s *CreateModelSuite) TestCreateCompany() {
	e := echo.New()
	company := models.CompanyDetail{
		Name:           "Tra Sua",
		Address:        "Le Duan",
		Balance:        10000,
		LoyaltyProgram: 100,
	}
	req, _ := http.NewRequest(http.MethodPost, "/companies", ToIOReader(company))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controllers.CreateCompany(c)
	assert.Equal(s.T(), http.StatusCreated, rec.Code)

}

func ToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}
func TestCreateModelSuite(t *testing.T) {
	suite.Run(t, new(CreateModelSuite))
}
