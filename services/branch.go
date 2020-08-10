package services

import (
	"errors"

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
		branch := ConvertToBranchDetail(item)
		result = append(result, branch)
	}

	return result, err
}

// ConvertToBranchDetail ...
func ConvertToBranchDetail(branch models.BranchBSON) models.BranchDetail {
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

// BranchCreate ...
func BranchCreate(body models.BranchCreatePayload) (models.BranchBSON, error) {
	var (
		branch       models.BranchBSON
		companyID, _ = primitive.ObjectIDFromHex(body.CompanyID)
		company, _   = dao.CompanyFindByID(companyID)
	)

	// Validate CompanyID
	if company.ID.IsZero() {
		return branch, errors.New("Khong tim thay Cong Ty")
	}

	// Create Branch
	branch = BranchCreatePayloadToBSON(body)
	doc, err := dao.BranchCreate(branch)

	return doc, err
}

// BranchCreatePayloadToBSON ...
func BranchCreatePayloadToBSON(branchCreatePayload models.BranchCreatePayload) models.BranchBSON {
	var (
		companyID, _ = primitive.ObjectIDFromHex(branchCreatePayload.CompanyID)
	)
	result := models.BranchBSON{
		CompanyID: companyID,
		Name:      branchCreatePayload.Name,
		Address:   branchCreatePayload.Address,
		Active:    branchCreatePayload.Active,
	}

	return result
}

// BranchUpdate ...
func BranchUpdate(BranchID primitive.ObjectID, body models.BranchUpdateBPayload) (models.BranchBSON, error) {
	var (
		branch = BranchUpdatePayloadToBSON(body)
	)

	// Update Branch
	doc, err := dao.BranchUpdate(BranchID, branch)

	return doc, err
}

// BranchUpdatePayloadToBSON ...
func BranchUpdatePayloadToBSON(branchUpdatePayload models.BranchUpdateBPayload) models.BranchBSON {
	result := models.BranchBSON{
		Name:    branchUpdatePayload.Name,
		Address: branchUpdatePayload.Address,
		Active:  branchUpdatePayload.Active,
	}

	return result
}

// BranchChangeActiveStatus ...
func BranchChangeActiveStatus(branchID primitive.ObjectID) (models.BranchBSON, error) {
	// Change Active Status
	doc, err := dao.BranchChangeActiveStatus(branchID)

	return doc, err
}
