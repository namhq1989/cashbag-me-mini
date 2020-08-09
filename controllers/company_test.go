package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"

	//"cashbag-me-mini/services"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	//"time"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"cashbag-me-mini/services"
)

type CompanySuite struct {
	suite.Suite
	company []models.CompanyBSON
}

var activeID = primitive.NewObjectID()
var updateID = primitive.NewObjectID()

func (s CompanySuite) SetupSuite() {
	database.Connect("CashBag-test")
	//removeOldDataCompany()
	addRecordCompany(activeID)
	addRecordCompany(updateID)

}

func (s CompanySuite) TearDownSuite() {
	//removeOldDataCompany()
}
func removeOldDataCompany() {
	database.DB.Collection("companies").DeleteMany(context.Background(), bson.M{})
}

func (s *CompanySuite) TestCompanyList() {
	var (
		res  []models.CompanyDetail
		list []models.CompanyDetail
		ctx  = context.Background()
	)

	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, "/companies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CompanyList(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)

	cursor, err := database.DB.Collection("todos").Find(ctx, bson.M{})
	if err != nil {
		assert.Equal(s.T(), http.StatusBadRequest, rec.Code)
	}
	cursor.All(ctx, &list)
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(s.T(), list, res)

}
func (s *CompanySuite) TestCompanyCreate() {
	e := echo.New()
	var (
		company = models.CompanyCreatePayload{
			Name:    "Highland",
			Address: "48 Nguyen Chanh",
			Active:  true,
		}
	)
	req, _ := http.NewRequest(http.MethodPost, "/companies", ToIOReader(company))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("body", &company)
	CompanyCreate(c)
	var res models.CompanyBSON
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal([]byte(rec.Body.String()), &res)
	assert.NotEqual(s.T(), nil, res)
	assert.Equal(s.T(), "Highland", company.Name)
	assert.Equal(s.T(), "48 Nguyen Chanh", company.Address)
}

func (s *CompanySuite) TestCompanyChangeActiveStatus() {
	e := echo.New()
	var x models.CompanyBSON
	req, _ := http.NewRequest(http.MethodPatch, "/companies/:id", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(activeID.Hex())
	CompanyChangeActiveStatus(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal(rec.Body.Bytes(), &x)
	assert.NotEqual(s.T(),nil,x.Active)
}
func (s *CompanySuite) TestCompanyUpdate() {
	var (
		company = models.CompanyUpdatePayload{
			Name:           "the coffee house",
			Address:        "67 Nguyen Huy Tuong",
			Balance:        100000,
			LoyaltyProgram: 100,
			Active:         false,
		}
		x models.CompanyBSON
	)
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPut, "/companies/:id", ToIOReader(company))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(updateID.Hex())
	c.Set("body", &company)
	CompanyUpdate(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal([]byte(rec.Body.String()), &x)
	assert.NotEqual(s.T(),nil,x.Name)

}

//addRecordCompany  ...
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


func TestCompanySuite(t *testing.T) {
	suite.Run(t, new(CompanySuite))
}
