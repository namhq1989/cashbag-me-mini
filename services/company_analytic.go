package services

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// transactionAnalyticHandleForTransaction ...
func transactionAnalyticHandleForTransaction(transaction models.TransactionBSON) (err error) {
	var (
		companyID = transaction.CompanyID
	)

	// Find CompanyAnalytic
	companyAnalytic, err := dao.CompanyAnalyticFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong Tim Thay CompanyAnalytic")
		return
	}

	// Set data Update CompanyAnalytic
	companyAnalytic.TotalTransaction++
	companyAnalytic.TotalRevenue += transaction.Amount
	companyAnalytic.TotalCommission += transaction.Commission
	if transaction.Postpaid == true {
		companyAnalytic.TotalDebt += transaction.Commission
		companyAnalytic.CountPostpaid++
		countPostpaid := companyAnalytic.CountPostpaid
		err = postpaidLogHandle(countPostpaid, companyID)
		if err != nil {
			err = errors.New("postpaidLogHandle That Bai")
			return
		}
	}

	// Update CompanyAnalytic
	err = dao.CompanyAnalyticUpdateTransactionProperties(companyAnalytic)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}

// postpaidLogHandle
func postpaidLogHandle(countPostpaid int, companyID primitive.ObjectID) (err error) {
	if countPostpaid > 3 {
		log.Println("Số đơn hàng trả sau vượt mức cho phép")
	}
	if countPostpaid > 6 {
		log.Println("Số đơn hàng trả sau vượt mức cho phép và deactive company")
		_, err = dao.CompanyChangeActiveStatus(companyID)
		return
	}
	return
}

// companyAnalyticHandleForBranchCreate ...
func companyAnalyticHandleForBranchCreate(branch models.BranchBSON) (err error) {
	var (
		companyID = branch.CompanyID
	)

	// Find CompanyAnalytic
	companyAnalytic, err := dao.CompanyAnalyticFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong Tim Thay CompanyAnalytic")
		return
	}

	// Set data Update CompanyAnalytic
	if branch.Active == true {
		companyAnalytic.ActiveBranch++
	} else {
		companyAnalytic.InactiveBranch++
	}

	// Update CompanyAnalytic
	err = dao.CompanyAnalyticUpdateBranchProperties(companyAnalytic)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}

// companyAnalyticHandleForBranchChangeActive ...
func companyAnalyticHandleForBranchChangeActive(branch models.BranchBSON) (err error) {
	var (
		companyID = branch.CompanyID
	)

	// Find CompanyAnalytic
	companyAnalytic, err := dao.CompanyAnalyticFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong Tim Thay CompanyAnalytic")
		return
	}

	// Set data Update CompanyAnalytic
	if branch.Active == true {
		companyAnalytic.ActiveBranch++
		companyAnalytic.InactiveBranch--
	} else {
		companyAnalytic.InactiveBranch++
		companyAnalytic.ActiveBranch--
	}

	// Update CompanyAnalytic
	err = dao.CompanyAnalyticUpdateBranchProperties(companyAnalytic)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}
