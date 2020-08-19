package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// CompanyAnalyticList ...
func CompanyAnalyticList() ([]models.CompanyAnalyticDetail, error) {
	var (
		result []models.CompanyAnalyticDetail
	)

	// Find
	doc, err := dao.CompanyAnalyticList()

	// Convert to CompanyDetail
	for _, item := range doc {
		company := convertToCompanyAnalyticDetail(item)
		result = append(result, company)
	}

	return result, err
}
