package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// LoyaltyProgramUserStatusBSON ...
	LoyaltyProgramUserStatusBSON struct {
		ID             primitive.ObjectID      `bson:"_id"`
		CompanyID      float64                 `bson:"companyID"`
		UserID         float64                 `bson:"userID"`
		Milestone      LoyaltyProgramMilestone `bson:"milestone"`
		CurrentExpense float64                 `bson:"currentExpense"`
		CreatedAt      time.Time               `bson:"createdAt"`
	}
)
