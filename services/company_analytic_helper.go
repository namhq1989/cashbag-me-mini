package services

import (
	"strconv"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

func convertToCompanyAnalyticDetail(doc models.CompanyAnalyticBSON) models.CompanyAnalyticDetail {
	var (
		company, _     = dao.CompanyFindByID(doc.CompanyID)
		nameCompany    = company.Name
		branch         string
		activeBranch   = doc.ActiveBranch
		totalBranch    = doc.TotalBranch
		inactiveBranch = totalBranch - activeBranch
	)
	branch = strconv.Itoa(activeBranch) + "(" + strconv.Itoa(inactiveBranch) + ")"

	// Convert
	result := models.CompanyAnalyticDetail{
		ID:              doc.ID,
		Company:         nameCompany,
		Branch:          branch,
		TotalRevenue:    doc.TotalRevenue,
		TotalCommission: doc.TotalCommission,
		TotalDebt:       doc.TotalDebt,
		CountPostpaid:   doc.CountPostpaid,
		Members:         doc.Members,
		UpdatedAt:       doc.UpdatedAt,
	}
	return result
}
