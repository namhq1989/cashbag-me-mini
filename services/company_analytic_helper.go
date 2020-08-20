package services

import (
	"errors"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// convertToCompanyAnalyticDetail ...
func convertToCompanyAnalyticDetail(doc models.CompanyAnalyticBSON) models.CompanyAnalyticDetail {
	var (
		company, _     = dao.CompanyFindByID(doc.CompanyID)
		nameCompany    = company.Name
		branch         string
		activeBranch   = doc.ActiveBranch
		inactiveBranch = doc.InactiveBranch
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
		TotalUser:       doc.TotalUser,
		UserSilver:      doc.UserSilver,
		UserGolden:      doc.UserGolden,
		UserDiamond:     doc.UserDiamond,
		UpdatedAt:       doc.UpdatedAt,
	}
	return result
}

// companyAnalyticHandleForTransaction ...
func companyAnalyticHandleForTransaction(transaction models.TransactionBSON, beforeUserLevel string, currentUserLevel string) (err error) {
	var (
		companyID    = transaction.CompanyID
		silverLevel  = "Silver"
		goldenLevel  = "Golden"
		diamondLevel = "Diamond"
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
	if beforeUserLevel == "" {
		if currentUserLevel == silverLevel {
			companyAnalytic.UserSilver++
		} else if currentUserLevel == goldenLevel {
			companyAnalytic.UserGolden++
		} else if currentUserLevel == diamondLevel {
			companyAnalytic.UserDiamond++
		}
	}
	if beforeUserLevel == silverLevel && currentUserLevel == goldenLevel {
		companyAnalytic.UserSilver--
		companyAnalytic.UserGolden++
	}
	if beforeUserLevel == silverLevel && currentUserLevel == diamondLevel {
		companyAnalytic.UserSilver--
		companyAnalytic.UserDiamond++
	}
	if beforeUserLevel == goldenLevel && currentUserLevel == diamondLevel {
		companyAnalytic.UserGolden--
		companyAnalytic.UserDiamond++
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

// companyAnalyticHandleForUserCreate ...
func companyAnalyticHandleForUserCreate(user models.UserBSON) (err error) {
	var (
		companyID = user.CompanyID
	)

	// Find CompanyAnalytic
	companyAnalytic, err := dao.CompanyAnalyticFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong Tim Thay CompanyAnalytic")
		return
	}

	// Set data Update CompanyAnalytic
	companyAnalytic.TotalUser++

	// Update CompanyAnalytic
	err = dao.CompanyAnalyticUpdateUserProperties(companyAnalytic)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}
