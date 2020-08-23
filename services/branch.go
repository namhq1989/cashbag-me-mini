package services

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

	//if err 
	if err !=nil  {
		err = errors.New("Khong the tao branch ")
		return doc,err
	}
	// Update CompanyAnalytic
	errCompanyAnalyticHandle := companyAnalyticHandleForBranchCreate(doc)
	if errCompanyAnalyticHandle != nil {
		return doc, errCompanyAnalyticHandle
	}
	return doc, err
}

// BranchUpdate ...
func BranchUpdate(branchID primitive.ObjectID, body models.BranchUpdatePayload) (models.BranchBSON,error) {
	var (
		// Prepare update  data
		filter     = bson.M{"_id": branchID}
		updateData = bson.M{"$set": bson.M{
			"name":      body.Name,
			"address":   body.Address,
			"updatedAt": time.Now(),
		}}
		doc, _ = dao.BranchFindByID(branchID)
		
	)

	// Update Branch
	err := dao.BranchUpdateByID(filter, updateData)

	return doc, err
}

// BranchChangeActiveStatus ...
func BranchChangeActiveStatus(branchID primitive.ObjectID, active bool) (models.BranchBSON,error) {
	var (
		// Prepare update data
		filter = bson.M{"_id": branchID}
		update = bson.M{"$set": bson.M{"active": active}}
		doc, _ = dao.BranchFindByID(branchID)
	)

	// Update
	err := dao.BranchUpdateByID(filter, update)
	if err != nil {
		err = errors.New("Khong the cap nhat branch")
		return doc,err
	}
	
	// Update CompanyAnalytic
	errCompanyAnalyticHandle := companyAnalyticHandleForBranchChangeActive(doc)
	if errCompanyAnalyticHandle != nil {
			return doc, errCompanyAnalyticHandle
		}

	return doc, err
}
