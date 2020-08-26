package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// LoyaltyProgramMilestone ...
	LoyaltyProgramMilestone struct {
		ID              string
		Expense         float64
		CashbackPercent float64
	}

	// LoyaltyProgramBSON ...
	LoyaltyProgramBSON struct {
		ID         primitive.ObjectID        `bson:"_id"`
		CompanyID  primitive.ObjectID        `bson:"companyID"`
		Milestones []LoyaltyProgramMilestone `bson:"milestones"`
		CreatedAt  time.Time                 `bson:"createdAt"`
	}

	// LoyaltyProgramCreatePayload ...
	LoyaltyProgramCreatePayload struct {
		CompanyID              string  `json:"companyID"`
		SilverExpense          float64 `json:"silverExpense"`
		SilverCashbackPercent  float64 `json:"silverCashbackPercent"`
		GoldExpense            float64 `json:"goldExpense"`
		GoldCashbackPercent    float64 `json:"goldCashbackPercent"`
		DiamondExpense         float64 `json:"diamondExpense"`
		DiamondCashbackPercent float64 `json:"diamondCashbackPercent"`
	}
)
