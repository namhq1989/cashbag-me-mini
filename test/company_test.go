package test

import (
	//"bytes"
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/services"
	"context"
	"encoding/json"
	//"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateModelSuite struct {
	suite.Suite
	company []models.CompanyBSON
}

var idPatch = primitive.NewObjectID()
var idPut = primitive.NewObjectID()

func (s CreateModelSuite) SetupSuite() {
	database.Connectdb("CashBag-test")
	//removeOldDataCompany()
	addRecordCompany(idPatch)
	addRecordCompany(idPut)

}

func (s CreateModelSuite) TearDownSuite() {
	removeOldDataCompany()
}
func removeOldDataCompany() {
	database.DB.Collection("companies").DeleteMany(context.Background(), bson.M{})
}

func (s *CreateModelSuite) TestListCompany() {
	e := echo.New()
	var (
		res  []models.CompanyDetail
		list []models.CompanyDetail
	)
	req, _ := http.NewRequest(http.MethodGet, "/companies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controllers.ListCompany(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	list = services.ListCompany()
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(s.T(), list, res)
}

func (s *CreateModelSuite) TestCreateCompany() {
	e := echo.New()
	var (
		company = models.PostCompany{
			Name:    "Highland",
			Address: " 48 Nguyen Chanh",
			Active:  true,
		}
		res = struct {
			InsertedID string `json:"InsertedID"`
		}{}
	)
	req, _ := http.NewRequest(http.MethodPost, "/companies", ToIOReader(company))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("body", &company)
	controllers.CreateCompany(c)
	assert.Equal(s.T(), http.StatusCreated, rec.Code)
	json.Unmarshal([]byte(rec.Body.String()), &res)
	assert.NotEqual(s.T(), res, nil)
}
func (s *CreateModelSuite) TestPatchCompany() {
	var (
		x = struct {
			MatchedCount  int    `json:"MatchedCount"`
			ModifiedCount int    `json:"ModifiedCount"`
			UpsertedCount int    `json:"UpsertedCount"`
			UpsertedID    string `json:"UpsertedID"`
		}{}
	)

	e := echo.New()
	req, _ := http.NewRequest(http.MethodPatch, "/companies/:id", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(idPatch.Hex())

	controllers.PatchCompany(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal(rec.Body.Bytes(), &x)
	assert.Equal(s.T(), 1, x.MatchedCount)
	assert.Equal(s.T(), 1, x.ModifiedCount)
	assert.Equal(s.T(), 0, x.UpsertedCount)
	assert.Equal(s.T(), "", x.UpsertedID)
}

func (s *CreateModelSuite) TestPutCompany() {
	var (
		company = models.PutCompany{
			Name:           "the coffee house",
			Address:        "67 Nguyen Huy Tuong",
			Balance:        100000,
			LoyaltyProgram: 100,
			Active:         false,
		}
		x = struct {
			MatchedCount  int    `json:"MatchedCount"`
			ModifiedCount int    `json:"ModifiedCount"`
			UpsertedCount int    `json:"UpsertedCount"`
			UpsertedID    string `json:"UpsertedID"`
		}{}
	)
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPut, "/companies/:id", ToIOReader(company))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(idPut.Hex())
	c.Set("body", &company)
	controllers.PutCompany(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal([]byte(rec.Body.String()), &x)
	assert.Equal(s.T(), 1, x.MatchedCount)
	assert.Equal(s.T(), 1, x.ModifiedCount)
	assert.Equal(s.T(), 0, x.UpsertedCount)
	assert.Equal(s.T(), "", x.UpsertedID)
}

//addRecord ...
func addRecordCompany(id primitive.ObjectID) {
	var (
		company = models.CompanyBSON{
			ID:      id,
			Name:    "tch",
			Address: "Nguyen Huy Tuong",
			Active:  false,
		}
	)
	database.DB.Collection("companies").InsertOne(context.TODO(), company)
}

// //ToIOReader ...
// func ToIOReader(i interface{}) io.Reader {
// 	b, _ := json.Marshal(i)
// 	return bytes.NewReader(b)
// }
func TestCreateModelSuite(t *testing.T) {
	suite.Run(t, new(CreateModelSuite))
}
