package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// LoyaltyProgramCreate ...
func LoyaltyProgramCreate(body models.LoyaltyProgramCreatePayload) (loyaltyProgram models.LoyaltyProgramBSON, err error) {

	// Validate Milestones
	err = validateLoyaltyProgramMilestones(body)
	if err != nil {
		return
	}

	// create userProgram
	loyaltyProgram = loyaltyProgramCreatePayloadToBSON(body)
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
