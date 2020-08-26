package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/util"
)

type UserSuite struct {
	suite.Suite
}

func (s UserSuite) SetupSuite() {
	util.HelperConnect()
	util.HelperCompanyCreateFake()
	util.HelperCompanyAnalyticCreateFake()
}

func (s UserSuite) TearDownSuite() {
	removeOldDataUser()
}
func removeOldDataUser() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyAnalyticCol().DeleteMany(context.Background(), bson.M{})
}

func (s *UserSuite) TestUserCreateSuccess() {
	var (
		user = models.UserCreatePayload{
			Name:      util.UserName,
			Address:   util.UserAddress,
		}
		response util.Response
	)

	//Create Context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", user)

	// Call UserCreate
	UserCreate(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.String()), &response)

	//Test
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(s.T(), nil, response["data"])

}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
