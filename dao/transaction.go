package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"

)

// TransactionCreate ...
func TransactionCreate(doc models.TransactionBSON) (models.TransactionBSON, error) {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)

	// Insert
	_, err := collection.InsertOne(ctx, doc)

	return doc, err
}

// TransactionFindByUserID ...
func TransactionFindByUserID(userID primitive.ObjectID) ([]models.TransactionBSON, error) {
	var (
		transactionCol = database.TransactionCol()
		ctx       = context.Background()
		filter    = bson.M{"userID": userID}
		result     = make([]models.TransactionBSON, 0)
		findOptions = options.Find()
	)

	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
	
	// Find
	cursor,err := transactionCol.Find(ctx, filter,findOptions)

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}






