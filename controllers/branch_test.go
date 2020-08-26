package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/util"
)

type BranchSuite struct {
	suite.Suite
}

func (s BranchSuite) SetupSuite() {
	util.HelperConnect()
	util.HelperBranchCreateFake()
	util.HelperCompanyCreateFake()
	util.HelperCompanyAnalyticCreateFake()
}

func (s BranchSuite) TearDownSuite() {
	removeOldData()
}

func removeOldData() {
	database.BranchCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyAnalyticCol().DeleteMany(context.Background(), bson.M{})
}

// TestBranchList ...
func (s *BranchSuite) TestBranchList() {
	var (
		response util.Response
	)

	// Create Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/branches", nil)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)

	// Call BranchList
	BranchList(c)

	// Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	// Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

// TestBranchCreateSuccess ...
func (s *BranchSuite) TestBranchCreateSuccess() {
	var (
		branch = models.BranchCreatePayload{
			CompanyID: util.CompanyString,
			Name:      "Hight SonLa",
			Address:   "120 SonLa",
			Active:    true,
		}
		response util.Response
	)

	// Create Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/branches", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", branch)
	c.Set("companyID", util.CompanyID)

	// Call BranchCreate
	BranchCreate(c)

	// Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	// Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

// TestBranchChangeActiveStatus ...
func (s *BranchSuite) TestBranchChangeActiveStatus() {
	var (
		response util.Response
	)

	// Create Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/branches/:id/active", nil)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("branch", util.Branch)

	// Call BranchChangeActiveStatus
	BranchChangeActiveStatus(c)

	// Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	// Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

// TestBranchUpdate ...
func (s *BranchSuite) TestBranchUpdate() {
	var (
		response             util.Response
		branchUpdateBPayload = models.BranchUpdatePayload{
			Name:    "Hight BinhDinh",
			Address: "111 BinhDinh",
		}
	)

	// Create Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/branches", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", branchUpdateBPayload)
	c.Set("branch", util.Branch)

	// Call BranchUpdate
	BranchUpdate(c)

	// Parse
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	// Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])
	assert.Equal(s.T(), "Thanh Cong!", response["message"])
}

func TestBranchSuite(t *testing.T) {
	suite.Run(t, new(BranchSuite))
}
