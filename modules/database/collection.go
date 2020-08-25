package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection name
const (
	companies                = "companies"
	branches                 = "branches"
	transactions             = "transactions"
	transactionAnalytics     = "transactionAnalytics"
	users                    = "users"
	loyaltyPrograms          = "loyaltyPrograms"
	companyAnalytics         = "companyAnalytics"
	testCharts               = "testCharts"
	loyaltyProgramUserStatus = "loyaltyProgramUserStatus"
)

// CompanyCol ...
func CompanyCol() *mongo.Collection {
	return db.Collection(companies)
}

// BranchCol ...
func BranchCol() *mongo.Collection {
	return db.Collection(branches)
}

// TransactionCol ...
func TransactionCol() *mongo.Collection {
	return db.Collection(transactions)
}

// TransactionAnalyticCol ...
func TransactionAnalyticCol() *mongo.Collection {
	return db.Collection(transactionAnalytics)
}

// UserCol ...
func UserCol() *mongo.Collection {
	return db.Collection(users)
}

// LoyaltyProgramCol ...
func LoyaltyProgramCol() *mongo.Collection {
	return db.Collection(loyaltyPrograms)
}

// CompanyAnalyticCol ...
func CompanyAnalyticCol() *mongo.Collection {
	return db.Collection(companyAnalytics)
}

// AnalyticChartCol ...
func AnalyticChartCol() *mongo.Collection {
	return db.Collection(testCharts)
}

// LoyaltyProgramUserStatusCol ...
func LoyaltyProgramUserStatusCol() *mongo.Collection {
	return db.Collection(loyaltyProgramUserStatus)
}
