package test

import (
	"bytes"
	"cashbag-me-mini/controllers"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/services"
	"context"
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type BranchSuite struct {
	suite.Suite
	Branches []models.BranchBSON
}

var idActive = primitive.NewObjectID()
var idUpdate = primitive.NewObjectID()

func (s BranchSuite) SetupSuite() {
	database.Connectdb("CashBag-test")
	addRecord(idActive) // for test Patch
	addRecord(idUpdate) // for test Put
	addCompany()        //for test Put
}

func (s BranchSuite) TearDownSuite() {
	removeOldData()
}

func removeOldData() {
	database.DB.Collection("branches").DeleteMany(context.Background(), bson.M{})
}

//TestListBranch ...
func (s *BranchSuite) TestListBranch() {
	var (
		branches []models.BranchDetail
		res      []models.BranchDetail
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/branches", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controllers.ListBranch(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	branches = services.ListBranch()
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(s.T(), branches, res)
}

//TestCreateBranch ...
func (s *BranchSuite) TestCreateBranch() {
	var (
		branch = models.PostBranch{
			NameCompany: "Hightland",
			Name:        "Hight SonLa",
			Address:     "120 SonLa",
			Active:      false,
		}
		res = struct {
			InsertedID string `json:"InsertedID"`
		}{}
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/branches", ToIOReader(branch))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("body", &branch)
	controllers.CreateBranch(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.NotEqual(s.T(), res, nil)
}

//ToIOReader ...
func ToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

//TestPatchBranch ...
func (s *BranchSuite) TestPatchBranch() {
	var (
		res = struct {
			MatchedCount  int         `json:"MatchedCount"`
			ModifiedCount int         `json:"ModifiedCount"`
			UpsertedCount int         `json:"UpsertedCount"`
			UpsertedID    interface{} `json:"UpsertedID"`
		}{}
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/branches/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(idActive.Hex())
	controllers.PatchBranch(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(s.T(), 1, res.MatchedCount)
	assert.Equal(s.T(), 1, res.ModifiedCount)
	assert.Equal(s.T(), 0, res.UpsertedCount)
	assert.Equal(s.T(), nil, res.UpsertedID)
}

//TestPutBranch ...
func (s *BranchSuite) TestPutBranch() {
	var (
		body = models.PutBranch{
			Name:    "Hight BinhDinh",
			Address: "111 BinhDinh",
			Active:  false,
		}
		res = struct {
			MatchedCount  int         `json:"MatchedCount"`
			ModifiedCount int         `json:"ModifiedCount"`
			UpsertedCount int         `json:"UpsertedCount"`
			UpsertedID    interface{} `json:"UpsertedID"`
		}{}
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/branches/:id", ToIOReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(idUpdate.Hex())
	c.Set("body", &body)
	controllers.PutBranch(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
	json.Unmarshal([]byte(rec.Body.String()), &res)
	assert.Equal(s.T(), 1, res.MatchedCount)
	assert.Equal(s.T(), 1, res.ModifiedCount)
	assert.Equal(s.T(), 0, res.UpsertedCount)
	assert.Equal(s.T(), nil, res.UpsertedID)
}

//addRecord ...
func addRecord(id primitive.ObjectID) {
	var (
		companyID, _ = primitive.ObjectIDFromHex("5f24d45125ea51bc57a8285b")
		branch       = models.BranchBSON{
			ID:        id,
			CompanyId: companyID,
			Name:      "Hight QuangTri",
			Address:   "120 QuangTri",
			Active:    false,
			CreateAt:  time.Now(),
		}
	)
	database.DB.Collection("branches").InsertOne(context.TODO(), branch)
}

//addCompany
func addCompany() {
	var (
		companyID, _ = primitive.ObjectIDFromHex("5f24d45125ea51bc57a8285b")
		company      = models.CompanyBSON{
			ID:             companyID,
			Name:           "Hightland",
			Address:        "HaiPhong",
			Balance:        10000000,
			LoyaltyProgram: 10,
			Active:         false,
			CreateAt:       time.Now(),
		}
	)
	database.DB.Collection("companies").InsertOne(context.TODO(), company)
}

func TestBranchSuite(t *testing.T) {
	suite.Run(t, new(BranchSuite))
}
