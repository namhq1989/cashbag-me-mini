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

	// create company
	doc, err := dao.CompanyCreate(company)

	return doc, err
}

// CompanyList ...
func CompanyList() ([]models.CompanyDetail, error) {
	var (
		result []models.CompanyDetail
	)

	doc, err := dao.CompanyList()

	for _, item := range doc {
		company := ConvertToCompanyDetail(item)
		result = append(result, company)
	}

	return result, err
}

// CompanyChangeActiveStatus . ...
func CompanyChangeActiveStatus(id primitive.ObjectID) (models.CompanyBSON, error) {
	doc, err := dao.CompanyChangeActiveStatus(id)

	return doc, err
}

// CompanyUpdate ....
func CompanyUpdate(id primitive.ObjectID, body models.CompanyUpdatePayload) (models.CompanyBSON, error) {
	var (
		company = CompanyUpdatePayloadToBSON(body)
	)
	//update company
	doc, err := dao.CompanyUpdate(id, company)

	return doc, err
}

// ConvertToCompanyDetail ....
func ConvertToCompanyDetail(doc models.CompanyBSON) models.CompanyDetail {
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

// CompanyCreatePayloadToBSON ....
func CompanyCreatePayloadToBSON(body models.CompanyCreatePayload) models.CompanyBSON {
	result := models.CompanyBSON{
		Name:    body.Name,
		Address: body.Address,
		Active:  body.Active,
	}
	return result
}

// CompanyUpdatePayloadToBSON ....
func CompanyUpdatePayloadToBSON(body models.CompanyUpdatePayload) models.CompanyBSON {
	result := models.CompanyBSON{
		Name:           body.Name,
		Address:        body.Address,
		Balance:        body.Balance,
		LoyaltyProgram: body.LoyaltyProgram,
		Active:         body.Active,
	}
	return result
}
