package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// CompanyAnalyticList ...
func CompanyAnalyticList() ([]models.CompanyAnalyticBSON, error) {
	var (
		companyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
		result             = make([]models.CompanyAnalyticBSON, 0)
	)

	// Find
	cursor, err := companyAnalyticCol.Find(ctx, bson.M{})

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}

// CompanyAnalyticCreate ...
func CompanyAnalyticCreate(companyAnalytic models.CompanyAnalyticBSON) error {
	var (
		companyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
	)

	// Create
	_, err := companyAnalyticCol.InsertOne(ctx, companyAnalytic)
	return err
}

// CompanyAnalyticUpdateByID ...
func CompanyAnalyticUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		CompanyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
	)

	_, err := CompanyAnalyticCol.UpdateOne(ctx, filter, updateData)

	return err
}

// CompanyAnalyticFindByCompanyID ...
func CompanyAnalyticFindByCompanyID(id primitive.ObjectID) (models.CompanyAnalyticBSON, error) {
	var (
		companyCol = database.CompanyAnalyticCol()
		ctx        = context.Background()
		result     models.CompanyAnalyticBSON
		filter     = bson.M{"companyID": id}
	)

	// Find
	err := companyCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}
