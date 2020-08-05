package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"

	"github.com/jinzhu/now"
)

//TranAnalytic ...
func TranAnalytic(date string) []models.TranAnalyticDetail {
	var (
		result   []models.TranAnalyticDetail
		dateType = now.MustParse(date)
	)
	dateType = dao.BeginningOfDay(dateType)
	tranAnalytics := dao.TranAnalytic(dateType)
	for _, item := range tranAnalytics {
		tranAnalytic := ConvertToTranAnalyticDetail(item)
		nameCompany := dao.GetNameCompanyById(item.CompanyId)
		nameBranch := dao.GetNameBranchById(item.BranchId)
		tranAnalytic.CompanyId.ID = item.CompanyId
		tranAnalytic.CompanyId.Name = nameCompany
		tranAnalytic.BranchId.ID = item.BranchId
		tranAnalytic.BranchId.Name = nameBranch
		result = append(result, tranAnalytic)
	}
	return result
}

//ConvertToTranAnalyticDetail ...
func ConvertToTranAnalyticDetail(x models.TranAnalyticBSON) models.TranAnalyticDetail {
	result := models.TranAnalyticDetail{
		ID:               x.ID,
		Date:             x.Date,
		TotalTransaction: x.TotalTransaction,
		TotalRevenue:     x.TotalRevenue,
		TotalCommission:  x.TotalCommission,
		UpdateAt:         x.UpdateAt,
	}
	return result
}
