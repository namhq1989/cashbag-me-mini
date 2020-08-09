package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//collection name
const (
	company = "companies"
	transaction ="transactions"
)
// ConnectCol ...
func ConnectCol(nameCol string) *mongo.Collection {
	return DB.Collection(nameCol)
}
// CompanyCol ...
func CompanyCol() *mongo.Collection{
	return DB.Collection(company)
}
//TransactionCol ...
func TransactionCol() *mongo.Collection{
	return DB.Collection(transaction)
}