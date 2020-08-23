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
		UpdatedAt  time.Time                 `bson:"updatedAt"`
	}

	// UserProgramCreatePayload ...
	UserProgramCreatePayload struct {
		CompanyID         string  `json:"companyID"`
		SilverSpending    float64 `json:"silverSpending"`
		SilverCommission  float64 `json:"silverCommission"`
		GoldenSpending    float64 `json:"goldenSpending"`
		GoldenCommission  float64 `json:"goldenCommission"`
		DiamondSpending   float64 `json:"diamondSpending"`
		DiamondCommission float64 `json:"diamondCommission"`
	}
)
