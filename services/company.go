package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
<<<<<<< HEAD

	"github.com/labstack/echo"
)

=======
)
>>>>>>> c4e7b26aca10bbf49db1183a063730b16411c8e6
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

<<<<<<< HEAD
//CompanyCreate func to ...
func CompanyCreate(c echo.Context) *models.CompanyDetail {
	company := dao.CompanyCreate(c)
	return company
}


//CompanyUpdate func to ...
func CompanyUpdate(c echo.Context) *models.CompanyDetail {
	company := dao.CompanyUpdate(c)
	return company
}
 
//CompanyActive func to ..
func CompanyActive(c echo.Context) *models.CompanyDetail {
	company := dao.CompanyActive(c)
	return company
}
=======
>>>>>>> c4e7b26aca10bbf49db1183a063730b16411c8e6

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
