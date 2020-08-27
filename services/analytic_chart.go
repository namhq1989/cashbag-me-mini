package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// AnalyticChart ...
func AnalyticChart(companyID primitive.ObjectID) ([]models.AnalyticChartDetail, error) {
	var (
		result   []models.AnalyticChartDetail
		toDate   = time.Now()
		fromDate = util.BeginningOfDay(toDate).AddDate(0, 0, -7)
		filter   = bson.M{
			"companyID": companyID,
			"date": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}
	)

	// Find
	analyticChartList, err := dao.AnalyticChart(companyID, filter)

	// Convert to TransactionAnalyticDetail
	for _, item := range analyticChartList {
		transactionAnalytic := convertToAnalyticChartDetail(item)
		result = append(result, transactionAnalytic)
	}

	return result, err
}

// convertToAnalyticChartDetail ...
func convertToAnalyticChartDetail(doc models.AnalyticChartBSON) models.AnalyticChartDetail {
	result := models.AnalyticChartDetail{
		ID:               doc.ID,
		CompanyID:        doc.CompanyID,
		Date:             doc.Date,
		TotalTransaction: doc.TotalTransaction,
		TotalRevenue:     doc.TotalRevenue,
		TotalCommission:  doc.TotalCommission,
		UpdateAt:         doc.UpdateAt,
	}
	return result
}
