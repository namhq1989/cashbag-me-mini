package dao

import (
	"cashbag-me-mini/util"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// AnalyticChart ...
func AnalyticChart(companyID primitive.ObjectID) ([]models.AnalyticChartBSON, error) {
	var (
		analyticChartCol = database.AnalyticChartCol()
		ctx              = context.Background()
		result           = make([]models.AnalyticChartBSON, 0)
		toDate           = time.Now()
		fromDate         = util.BeginningOfDay(toDate).AddDate(0, 0, -7)
		filter           = bson.M{
			"companyID": companyID,
			"date": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}
	)

	//Find
	cursor, err := analyticChartCol.Find(ctx, filter)

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}
