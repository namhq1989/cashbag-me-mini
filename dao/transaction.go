package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

//CreateTransaction ...
func CreateTransaction(transaction models.TransactionBSON, balance float64) *mongo.InsertOneResult {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)
	balanceCurrent := balance - transaction.Commission
	result, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		log.Fatal(err)
	}
	UpdateBalance(transaction.CompanyID, balanceCurrent)
	TransactionAnalyticHandle(transaction)

	return result
}
