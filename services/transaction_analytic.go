package services

import (
	"github.com/jinzhu/now"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// TransactionAnalyticList ...
func TransactionAnalyticList(date string) ([]models.TransactionAnalyticDetail, error) {
	var (
		result   []models.TransactionAnalyticDetail
		dateType = now.MustParse(date)
		beginDay = util.BeginningOfDay(dateType)
	)

	// Find
	transactionAnalyticList, err := dao.TransactionAnalyticList(beginDay)

	// Convert to TransactionAnalyticDetail
	for _, item := range transactionAnalyticList {
		transactionAnalytic := convertToTransactionAnalyticDetail(item)
		result = append(result, transactionAnalytic)
	}

	return result, err
}
