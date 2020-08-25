package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// LoyaltyProgramCreate ...
func LoyaltyProgramCreate(doc models.LoyaltyProgramBSON) (models.LoyaltyProgramBSON, error) {
	var (
		loyaltyProgramCol = database.LoyaltyProgramCol()
		ctx               = context.Background()
	)

	//insert
	_, err := loyaltyProgramCol.InsertOne(ctx, doc)

	return doc, err
}

// LoyaltyProgramFindByCompanyID ...
func LoyaltyProgramFindByCompanyID(id primitive.ObjectID) (models.LoyaltyProgramBSON, error) {
	var (
		loyaltyProgramCol = database.LoyaltyProgramCol()
		ctx               = context.Background()
		result            models.LoyaltyProgramBSON
		filter            = bson.M{"companyID": id}
	)

	// Find
	err := loyaltyProgramCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}
