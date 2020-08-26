package services

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

	if err != nil {
		err = errors.New("Khong the tao company")
		return doc, err
	}

	// Create company analytic
	err = companyCreateCompanyAnalytic(doc.ID)
	if err != nil {
		return doc, err
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
func CompanyChangeActiveStatus(companyID primitive.ObjectID, active bool) (doc models.CompanyBSON, err error) {
	var (
		// Prepare update data
		filter = bson.M{"_id": companyID}
		update = bson.M{"$set": bson.M{"active": active}}
	)

	// Update
	err = dao.CompanyUpdateByID(filter, update)
	if err != nil {
		err = errors.New("Khong the cap nhat company")
		return doc, err
	}
	doc, _ = dao.CompanyFindByID(companyID)

	return doc, err
}

// CompanyUpdate ...
func CompanyUpdate(companyID primitive.ObjectID, body models.CompanyUpdatePayload) (doc models.CompanyBSON, err error) {
	var (
		filter     = bson.M{"_id": companyID}
		updateData = bson.M{"$set": bson.M{
			"name":            body.Name,
			"address":         body.Address,
			"balance":         body.Balance,
			"cashbackPercent": body.CashbackPercent,
			"active":          body.Active,
			"paidType":        body.PaidType,
			"updatedAt":       time.Now(),
		}}
	)

	// Update company
	err = dao.CompanyUpdateByID(filter, updateData)
	if err != nil {
		err = errors.New("Khong the cap nhat company")
		return
	}
	doc, _ = dao.CompanyFindByID(companyID)

	return
}

// CompanyUpdateActiveTrue ...
func CompanyUpdateActiveTrue(companyID primitive.ObjectID) error {
	var (
		// Prepare update  data
		filter     = bson.M{"_id": companyID}
		updateData = bson.M{"$set": bson.M{
			"active": true,
		}}
	)

	// Update company
	err := dao.CompanyUpdateByID(filter, updateData)

	return err
}
