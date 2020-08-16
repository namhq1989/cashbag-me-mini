package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"errors"
)

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

	if branch.Active == true {
		companyAnalytic.ActiveBranch++
	} else {
		companyAnalytic.InactiveBranch++
	}

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

	if branch.Active == true {
		companyAnalytic.ActiveBranch++
		companyAnalytic.InactiveBranch--
	} else {
		companyAnalytic.InactiveBranch++
		companyAnalytic.ActiveBranch--
	}

	err = dao.CompanyAnalyticUpdateBranchProperties(companyAnalytic)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}
