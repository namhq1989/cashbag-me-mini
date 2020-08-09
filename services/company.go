package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// CompanyCreate ...
func CompanyCreate(body models.CompanyCreatePayload) (models.CompanyBSON, error) {
	var (
		company = CompanyCreatePayloadToBSON(body)
	)

	// Create Company
	doc, err := dao.CompanyCreate(company)

	return doc, err
}

// CompanyList ...
func CompanyList() ([]models.CompanyDetail, error) {
	var (
		result []models.CompanyDetail
	)
	// Find
	doc, err := dao.CompanyList()

	// Convert to CompanyDetail
	for _, item := range doc {
		company := convertToCompanyDetail(item)
		result = append(result, company)
	}

	return result, err
}

// CompanyChangeActiveStatus ...
func CompanyChangeActiveStatus(CompanyID primitive.ObjectID) (models.CompanyBSON, error) {
	// Change Active Status
	doc, err := dao.CompanyChangeActiveStatus(CompanyID)

	return doc, err
}

// CompanyUpdate ....
func CompanyUpdate(CompanyID primitive.ObjectID, body models.CompanyUpdatePayload) (models.CompanyBSON, error) {
	var (
		company = CompanyCreatePayloadToBSON(body)
	)

	// Update Company
	doc, err := dao.CompanyUpdate(CompanyID, body)

	return doc, err
}

// ConvertToCompanyDetail ...
func ConvertToCompanyDetail(x models.CompanyBSON) models.CompanyDetail {
	result := models.CompanyDetail{
		ID:             x.ID,
		Name:           x.Name,
		Address:        x.Address,
		Balance:        x.Balance,
		LoyaltyProgram: x.LoyaltyProgram,
		Active:         x.Active,
		CreatedAt:      x.CreatedAt,
		UpdatedAt:      x.UpdatedAt,
	}

	return result
}

// CompanyCreatePayloadToBSON ...
func CompanyCreatePayloadToBSON(payloadOfCreateCompany models.Co) models.CompanyBSON {

	result := models.CompanyBSON{
		Name:    payloadOfCreateCompany.Name,
		Address: payloadOfCreateCompany.Address,
		Active:  payloadOfCreateCompany.Active,
	}

	return result
}

// CompanyUpdatePayloadToBSON ...
func CompanyUpdatePayloadToBSON(payloadOfUpdateCompany models.PayloadOfUpdateCompany) models.CompanyBSON {
	result := models.CompanyBSON{
		Name:           payloadOfUpdateCompany.Name,
		Address:        payloadOfUpdateCompany.Address,
		Balance:        payloadOfUpdateCompany.Balance,
		LoyaltyProgram: payloadOfUpdateCompany.LoyaltyProgram,
		Active:         payloadOfUpdateCompany.Active,
	}

	return result
}
