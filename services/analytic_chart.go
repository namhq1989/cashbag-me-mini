package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// AnalyticChart ...
func AnalyticChart(companyID primitive.ObjectID) ([]models.AnalyticChartDetail, error) {
	var (
		result []models.AnalyticChartDetail
	)

	// Find
	analyticChartList, err := dao.AnalyticChart(companyID)

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
