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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func (s CreateModelSuite) SetupSuite() {
	database.Connectdb("CashBag-test")
	removeOldData()
	addRecord(idPatch)

}

func (s CreateModelSuite) TearDownSuite() {
	//removeOldData()
}
func removeOldData() {
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
			Name:    "HighLand",
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




func ToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

//addRecord ...
func addRecord(id primitive.ObjectID) {
	var (
		company = models.CompanyBSON{
			Name:     "HightLand",
			Address:  "Nguyen Huy Tuong",
			Active:   false,
			CreateAt: time.Now(),
		}
	)
	database.DB.Collection("companies").InsertOne(context.TODO(), company)
}

func TestCreateModelSuite(t *testing.T) {
	suite.Run(t, new(CreateModelSuite))
}
