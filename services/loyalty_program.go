package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// LoyaltyProgramCreate ...
func LoyaltyProgramCreate(body models.LoyaltyProgramCreatePayload, companyID primitive.ObjectID) (loyaltyProgram models.LoyaltyProgramBSON, err error) {

	// Validate Milestones
	err = validateLoyaltyProgramMilestones(body)
	if err != nil {
		return
	}

	// create userProgram
	loyaltyProgram = loyaltyProgramCreatePayloadToBSON(body,companyID)
	loyaltyProgram, err = dao.LoyaltyProgramCreate(loyaltyProgram)
	if err != nil {
		return
	}

	err = CompanyUpdateActiveTrue(loyaltyProgram.CompanyID)
	if err != nil {
		return
	}

	return
}
