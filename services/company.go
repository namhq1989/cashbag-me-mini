package services

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// CompanyCreate ...
func CompanyCreate(body models.CompanyCreatePayload) (models.CompanyBSON, error) {
	var (
		company = companyCreatePayloadToBSON(body)
	)

	// Create company
	doc, err := dao.CompanyCreate(company)

	// Create company analytic
	if err == nil {
		errCompanyAnalyticCreate := dao.CompanyAnalyticCreate(doc.ID)
		if errCompanyAnalyticCreate != nil {
			errCompanyAnalyticCreate = errors.New("Khong The Tao Company Analytic")
			return doc, errCompanyAnalyticCreate
		}
	}

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
func CompanyChangeActiveStatus(id primitive.ObjectID) (models.CompanyBSON, error) {
	// Change active
	doc, err := dao.CompanyChangeActiveStatus(id)

	return doc, err
}

// CompanyUpdate ...
func CompanyUpdate(id primitive.ObjectID, body models.CompanyUpdatePayload) (models.CompanyBSON, error) {
	var (
		company = companyUpdatePayloadToBSON(body)
	)

	// Update company
	doc, err := dao.CompanyUpdate(id, company)

	return doc, err
}
