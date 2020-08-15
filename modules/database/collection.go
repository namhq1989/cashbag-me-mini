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
	user                 = "user"
	userProgram          = "userProgram"
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
func UserProgramCol() *mongo.Collection{
	return db.Collection(userProgram)
}