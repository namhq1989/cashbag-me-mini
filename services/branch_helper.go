package services

import (
	"cashbag-me-mini/util"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// convertToBranchDetail ...
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

// branchCreatePayloadToBSON ...
func branchCreatePayloadToBSON(branchCreatePayload models.BranchCreatePayload) models.BranchBSON {
	var (
		companyID, _ = util.ValidationObjectID(branchCreatePayload.CompanyID)
	)
	result := models.BranchBSON{
		CompanyID: companyID,
		Name:      branchCreatePayload.Name,
		Address:   branchCreatePayload.Address,
		Active:    branchCreatePayload.Active,
	}

	return result
}

// branchUpdatePayloadToBSON ...
func branchUpdatePayloadToBSON(branchUpdatePayload models.BranchUpdateBPayload) models.BranchBSON {
	result := models.BranchBSON{
		Name:    branchUpdatePayload.Name,
		Address: branchUpdatePayload.Address,
	}

	return result
}
