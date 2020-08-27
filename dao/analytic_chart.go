package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// AnalyticChart ...
func AnalyticChart(companyID primitive.ObjectID, filter bson.M) ([]models.AnalyticChartBSON, error) {
	var (
		analyticChartCol = database.AnalyticChartCol()
		ctx              = context.Background()
		result           = make([]models.AnalyticChartBSON, 0)
	)

	//Find
	cursor, err := analyticChartCol.Find(ctx, filter)

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}
