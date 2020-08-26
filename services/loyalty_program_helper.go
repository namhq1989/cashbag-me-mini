package services

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// loyaltyProgramCreatePayloadToBSON ...
func loyaltyProgramCreatePayloadToBSON(body models.LoyaltyProgramCreatePayload) models.LoyaltyProgramBSON {
	var (
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		silver       = "silver"
		gold         = "gold"
		diamond      = "diamond"
		milestones   []models.LoyaltyProgramMilestone
	)

	// Define milestones
	silverMilestone := models.LoyaltyProgramMilestone{
		ID:              silver,
		Expense:         body.SilverExpense,
		CashbackPercent: body.SilverCashbackPercent,
	}
	goldMilestone := models.LoyaltyProgramMilestone{
		ID:              gold,
		Expense:         body.GoldExpense,
		CashbackPercent: body.GoldCashbackPercent,
	}
	diamondMilestone := models.LoyaltyProgramMilestone{
		ID:              diamond,
		Expense:         body.DiamondExpense,
		CashbackPercent: body.DiamondCashbackPercent,
	}
	milestones = append(milestones, silverMilestone, goldMilestone, diamondMilestone)

	// LoyaltyProgramBSON ...
	result := models.LoyaltyProgramBSON{
		ID:         primitive.NewObjectID(),
		CompanyID:  companyID,
		Milestones: milestones,
		CreatedAt:  time.Now(),
	}
	return result
}

func validateLoyaltyProgramMilestones(body models.LoyaltyProgramCreatePayload) (err error) {
	var (
		silverExpense          = body.SilverExpense
		silverCashbackPercent  = body.SilverCashbackPercent
		goldExpense            = body.GoldExpense
		goldCashbackPercent    = body.GoldCashbackPercent
		diamondExpense         = body.DiamondExpense
		diamondCashbackPercent = body.DiamondCashbackPercent
	)
	// validate
	if silverExpense <= 0 || silverCashbackPercent <= 0 {
		err = errors.New("silver milestone khong hop li")
		return
	}
	if goldExpense <= silverExpense || goldCashbackPercent <= silverCashbackPercent {
		err = errors.New("golden milestone khong hop li")
		return
	}
	if diamondExpense <= goldExpense || diamondCashbackPercent <= goldCashbackPercent {
		err = errors.New("diamond milestone khong hop li")
		return
	}
	return
}
