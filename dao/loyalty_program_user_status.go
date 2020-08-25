package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// LoyaltyProgramUserStatusCreate ...
func LoyaltyProgramUserStatusCreate(doc models.LoyaltyProgramUserStatusBSON) (models.LoyaltyProgramUserStatusBSON, error) {
	var (
		loyaltyProgramUserStatusCol = database.LoyaltyProgramUserStatusCol()
		ctx                         = context.Background()
	)

	//insert
	_, err := loyaltyProgramUserStatusCol.InsertOne(ctx, doc)

	return doc, err
}

// LoyaltyProgramUserStatusFindByCompanyIDAndUserID ...
func LoyaltyProgramUserStatusFindByCompanyIDAndUserID(companyID primitive.ObjectID, userID primitive.ObjectID) (models.LoyaltyProgramUserStatusBSON, error) {
	var (
		loyaltyProgramUserStatusCol = database.LoyaltyProgramUserStatusCol()
		ctx                         = context.Background()
		result                      models.LoyaltyProgramUserStatusBSON
		filter                      = bson.M{
			"companyID": companyID,
			"userID":    userID,
		}
		findOptions = options.FindOne()
	)
	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})

	// Find
	err := loyaltyProgramUserStatusCol.FindOne(ctx, filter, findOptions).Decode(&result)

	return result, err
}

// LoyaltyProgramUserStatusUpdate ...
func LoyaltyProgramUserStatusUpdate(filter bson.M, updateData bson.M) (err error) {
	var (
		loyaltyProgramUserStatusCol = database.LoyaltyProgramUserStatusCol()
		ctx       = context.Background()
	)

	_, err = loyaltyProgramUserStatusCol.UpdateOne(ctx, filter, updateData)

	return err
}
