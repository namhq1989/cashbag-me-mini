package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

//CreateTransaction ...
func CreateTransaction(id models.TransactionBSON) *mongo.InsertOneResult {
	var (
		collection = database.ConnectCol("transaction")
		ctx        = context.Background()
	)
	result, err := collection.InsertOne(ctx, id)
	if err != nil {
		fmt.Println(err)
	}
	HandleTranAnalytic(id)
	return result
}
