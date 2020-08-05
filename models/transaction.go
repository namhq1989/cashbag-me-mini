package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	TransactionBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		CompanyID      primitive.ObjectID `bson:"companyId"`
		BranchID       primitive.ObjectID `bson:"branchId"`
		User           string             `bson:"user"`
		Amount         float64            `bson:"amount"`
		Commission     float64            `bson:"commission"`
		LoyaltyProgram float64            `bson:"loyaltyprogram"`
		CreateAt       time.Time          `bson:"createAt"`
	}
)
