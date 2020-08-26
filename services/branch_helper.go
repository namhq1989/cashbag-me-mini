package services

import (
	"errors"
	"time"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func convertToBranchDetail(branch models.BranchBSON) models.BranchDetail {
	var (
		companyBrief models.CompanyBrief
		company, _   = dao.CompanyFindByID(branch.CompanyID)
		nameCompany  = company.Name
	)

	// Add information for companyBrief
	companyBrief.ID = branch.CompanyID
	companyBrief.Name = nameCompany

	// branchDetail
	result := models.BranchDetail{
		ID:        branch.ID,
		Company:   companyBrief,
		Name:      branch.Name,
		Address:   branch.Address,
		Active:    branch.Active,
		CreatedAt: branch.CreatedAt,
		UpdatedAt: branch.UpdatedAt,
	}

	return result
}

func branchCreatePayloadToBSON(branchCreatePayload models.BranchCreatePayload, companyID primitive.ObjectID) models.BranchBSON {
	result := models.BranchBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: companyID,
		Name:      branchCreatePayload.Name,
		Address:   branchCreatePayload.Address,
		Active:    branchCreatePayload.Active,
		CreatedAt: time.Now(),
	}

	return result
}

func branchCreateUpdateCompanyAnalytic(branch models.BranchBSON) (err error) {
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
	}
	companyAnalytic.TotalBranch++

	// Update CompanyAnalytic
	filter := bson.M{"_id": companyAnalytic.ID}
	update := bson.M{"$set": bson.M{
		"activeBranch": companyAnalytic.ActiveBranch,
		"totalBranch":  companyAnalytic.TotalBranch,
		"updatedAt":    time.Now(),
	}}
	err = dao.CompanyAnalyticUpdateByID(filter, update)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}

func branchChangeActiveStatusUpdateCompanyAnalytic(branch models.BranchBSON) (err error) {
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
		companyAnalytic.ActiveBranch--
	}

	// Update CompanyAnalytic
	filter := bson.M{"_id": companyAnalytic.ID}
	update := bson.M{"$set": bson.M{
		"activeBranch": companyAnalytic.ActiveBranch,
		"updatedAt":    time.Now(),
	}}
	err = dao.CompanyAnalyticUpdateByID(filter, update)
	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}
