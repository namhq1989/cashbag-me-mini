package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection name
const (
	company              = "companies"
	branch               = "branches"
	transaction          = "transactions"
	transactionAnalytics = "transactionAnalytics"
	user                 = "users"
	userProgram          = "userProgram"
	companyAnalytics     = "companyAnalytics"
	testCharts           = "testCharts"
)

// CompanyCol ...
func CompanyCol() *mongo.Collection {
	return db.Collection(company)
}

// BranchCol ...
func BranchCol() *mongo.Collection {
	return db.Collection(branch)
}

// TransactionCol ...
func TransactionCol() *mongo.Collection {
	return db.Collection(transaction)
}

// TransactionAnalyticCol ...
func TransactionAnalyticCol() *mongo.Collection {
	return db.Collection(transactionAnalytics)
}

// UserCol ...
func UserCol() *mongo.Collection {
	return db.Collection(user)
}

// UserProgramCol ...
func UserProgramCol() *mongo.Collection {
	return db.Collection(userProgram)
}

// CompanyAnalyticCol ...
func CompanyAnalyticCol() *mongo.Collection {
	return db.Collection(companyAnalytics)
}

// AnalyticChartCol ...
func AnalyticChartCol() *mongo.Collection {
	return db.Collection(testCharts)
}
