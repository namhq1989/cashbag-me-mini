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

// TransactionAnalyticHandle ...
func TransactionAnalyticHandle(transaction models.TransactionBSON) {
	tranAnalytic, check := dao.TransactionAnalyticFilter(transaction)
	if check == false {
		dao.TransactionAnalyticCreate(transaction)
	} else {
		dao.TransactionAnalyticUpdate(tranAnalytic, transaction)
	}
}

// convertToTransactionAnalyticDetail ...
func convertToTransactionAnalyticDetail(doc models.TransactionAnalyticBSON) models.TransactionAnalyticDetail {
	var (
		company, _   = dao.CompanyFindByID(doc.CompanyID)
		nameCompany  = company.Name
		companyBrief models.CompanyBrief
		branch, _    = dao.BranchFindByID(doc.BranchID)
		nameBranch   = branch.Name
		brannchBrief models.BranchBrief
	)

	// Add information companyBrief
	companyBrief.ID = doc.CompanyID
	companyBrief.Name = nameCompany

	// Add information branchBrief
	brannchBrief.ID = doc.BranchID
	brannchBrief.Name = nameBranch

	// Transaction Analytic Detail
	result := models.TransactionAnalyticDetail{
		ID:               doc.ID,
		Company:          companyBrief,
		Branch:           brannchBrief,
		Date:             doc.Date,
		TotalTransaction: doc.TotalTransaction,
		TotalRevenue:     doc.TotalRevenue,
		TotalCommission:  doc.TotalCommission,
		UpdateAt:         doc.UpdateAt,
	}
	return result
}
