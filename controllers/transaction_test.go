package controllers

import (
	"cashbag-me-mini/modules/redis"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/modules/database"
	"cashbag-me-mini/util"
)

type TransactionSuite struct {
	suite.Suite
}

func (s TransactionSuite) SetupSuite() {
	util.HelperConnect()
	redis.Connect()
	util.HelperCompanyCreateFake()
	util.HelperBranchCreateFake()
	util.HelperUserCreateFake()
}

func (s TransactionSuite) TearDownSuite() {
	removeOldDataTransaction()
}
func removeOldDataTransaction() {
	database.BranchCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
	database.TransactionCol().DeleteMany(context.Background(), bson.M{})
	database.TransactionAnalyticCol().DeleteMany(context.Background(), bson.M{})
}

func (s *TransactionSuite) TestTransactionCreateSuccess() {
	s.T().Skip()
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
