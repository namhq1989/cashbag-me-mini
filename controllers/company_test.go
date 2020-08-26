package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/util"
)

type CompanySuite struct {
	suite.Suite
}

func (s CompanySuite) SetupSuite() {
	util.HelperConnect()
	util.HelperCompanyCreateFake()
}

func (s CompanySuite) TearDownSuite() {
	removeOldDataCompany()
}
func removeOldDataCompany() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
}

func (s *CompanySuite) TestCompanyList() {
	var (
		response     util.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/phuc/cashbag-me-mini/src/detail_companies.json")
	)

	//Create Context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, "/companies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call CompanyList
	CompanyList(c)
	documentLoader := gojsonschema.NewGoLoader(&response)
	// Parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	
	
	//Test
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println("err1")	
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func (s *CompanySuite) TestCompanyCreateSuccess() {
	var (
		company = models.CompanyCreatePayload{
			Name:    "Highland",
			Address: "48 Nguyen Chanh",
		}
		response util.Response
	)

	//Create Context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/companies", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", company)

	// Call CompanyCreate
	CompanyCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	//Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

func (s *CompanySuite) TestCompanyChangeActiveStatus() {
	var (
		response  util.Response
		companyID = util.CompanyID
	)

	//Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPatch, "/companies/:id/active", nil)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.SetParamNames("id")
	c.SetParamValues(companyID)

	//Call CompanyChangeActiveStatus
	CompanyChangeActiveStatus(c)

	//Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	//Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}
func (s *CompanySuite) TestCompanyUpdate() {
	var (
		response             util.Response
		companyID            = util.CompanyID
		companyUpdatePayload = models.CompanyUpdatePayload{
			Name:            "the coffee house",
			Address:         "67 Nguyen Huy Tuong",
			Balance:         100000,
			CashbackPercent: 100,
			Active:          false,
		}
	)

	//Create context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPut, "/companies/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.SetParamNames("id")
	c.SetParamValues(companyID)
	c.Set("body", companyUpdatePayload)

	// Call CompanyUpdate
	CompanyUpdate(c)

	//Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	//Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

func TestCompanySuite(t *testing.T) {
	suite.Run(t, new(CompanySuite))
}
