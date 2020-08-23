package services

import (
	"cashbag-me-mini/models"
)

// convertToCompanyDetail ...
func convertToCompanyDetail(doc models.CompanyBSON) models.CompanyDetail {
	result := models.CompanyDetail{
		ID:             doc.ID,
		Name:           doc.Name,
		Address:        doc.Address,
		Balance:        doc.Balance,
		LoyaltyProgram: doc.LoyaltyProgram,
		Active:         doc.Active,
		CreatedAt:      doc.CreatedAt,
		UpdatedAt:      doc.UpdatedAt,
	}
	return result
}

// companyCreatePayloadToBSON ...
func companyCreatePayloadToBSON(body models.CompanyCreatePayload) models.CompanyBSON {
	result := models.CompanyBSON{
		Name:    body.Name,
		Address: body.Address,
	}
	return result
}

