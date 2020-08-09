package services

import (
	"github.com/jinzhu/now"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

//TranAnalytic ...
func TranAnalytic(date string) []models.TransactionAnalyticDetail {
	var (
		result   []models.TransactionAnalyticDetail
		dateType = now.MustParse(date)
		beginDay = dao.BeginningOfDay(dateType)
	)

	// Find
	docs, err := dao.TransactionAnalytic(beginDay)

	// Convert to TransactionAnalyticDetail
	for _, item := range tranAnalytics {
		tranAnalytic := ConvertToTranAnalyticDetail(item)
		result = append(result, tranAnalytic)
	}

	return result, err
}

//ConvertToTranAnalyticDetail ...
func ConvertToTranAnalyticDetail(x models.TransactionAnalyticBSON) models.TransactionAnalyticDetail {
	var(
		nameCompany = dao.GetNameCompanyById(x.CompanyId).Name
		companyBrief  models.CompanyBrief
		nameBranch = dao.BranchDocById(x.BranchId).Name
		brannchBrief  models.BranchBrief
	)

	// Add information companyBrief
	companyBrief.ID = x.CompanyID
	companyBrief.Name = nameCompany

	// Add information branchBrief
	brannchBrief.ID = x.BranchID
	tbrannchBrief.Name = nameBranch

	// Transaction Analytic Detail
	result := models.TranAnalyticDetail{
		ID:               x.ID,
		Company:		companyBrief,
		Branch:			brannchBrief,
		Date:             x.Date,
		TotalTransaction: x.TotalTransaction,
		TotalRevenue:     x.TotalRevenue,
		TotalCommission:  x.TotalCommission,
		UpdateAt:         x.UpdateAt,
	}
	return result
}
