package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// BranchList ...
func BranchList() ([]models.BranchDetail, error) {
	var (
		result []models.BranchDetail
	)

	// Find
	docs, err := dao.BranchList()

	// Convert to BranchDetail
	for _, item := range docs {
		branch := convertToBranchDetail(item)
		result = append(result, branch)
	}

	return result, err
}

// BranchCreate ...
func BranchCreate(body models.BranchCreatePayload) (models.BranchBSON, error) {
	var (
		branch models.BranchBSON
	)

	// Create Branch
	branch = branchCreatePayloadToBSON(body)
	doc, err := dao.BranchCreate(branch)

	// Update CompanyAnalytic
	if err == nil {
		errCompanyAnalyticHandle := companyAnalyticHandleForBranchCreate(doc)
		if errCompanyAnalyticHandle != nil {
			return doc, errCompanyAnalyticHandle
		}
	}
	return doc, err
}

// BranchUpdate ...
func BranchUpdate(BranchID primitive.ObjectID, body models.BranchUpdateBPayload) (models.BranchBSON, error) {
	var (
		branch = branchUpdatePayloadToBSON(body)
	)

	// Update Branch
	doc, err := dao.BranchUpdate(BranchID, branch)

	return doc, err
}

// BranchChangeActiveStatus ...
func BranchChangeActiveStatus(branchID primitive.ObjectID) (models.BranchBSON, error) {
	// Change Active Status
	doc, err := dao.BranchChangeActiveStatus(branchID)

	// Update CompanyAnalytic
	if err == nil {
		errCompanyAnalyticHandle := companyAnalyticHandleForBranchChangeActive(doc)
		if errCompanyAnalyticHandle != nil {
			return doc, errCompanyAnalyticHandle
		}
	}
	return doc, err
}
