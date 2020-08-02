package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)
//CompanyList to ...
func CompanyList() []models.CompanyDetail {
	var (
		result []models.CompanyDetail
	)
	companies := dao.CompanyList()
	for _, item := range companies {
		company := convertToCompanyDetail(item)
		result = append(result, company)
	}
	return result
}


//convertToCompanyDetail to ..
func convertToCompanyDetail(x models.CompanyBSON) models.CompanyDetail {
	result := models.CompanyDetail{
		ID:             x.ID,
		Name:           x.Name,
		Address:        x.Address,
		Balance:        x.Balance,
		LoyaltyProgram: x.LoyaltyProgram,
		Active:         x.Active,
		CreateAt:       x.CreateAt,
		UpdateAt:       x.UpdateAt,
	}
	return result
}
