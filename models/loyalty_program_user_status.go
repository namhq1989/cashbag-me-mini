package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// LoyaltyProgramUserStatusBSON ...
	LoyaltyProgramUserStatusBSON struct {
		ID             primitive.ObjectID      `bson:"_id"`
		CompanyID      primitive.ObjectID      `bson:"companyID"`
		UserID         primitive.ObjectID      `bson:"userID"`
		Milestone      LoyaltyProgramMilestone `bson:"milestone"`
		CurrentExpense float64                 `bson:"currentExpense"`
		UpdatedAt      time.Time               `bson:"updatedAt"`
		CreatedAt      time.Time               `bson:"createdAt"`
	}
	
)
