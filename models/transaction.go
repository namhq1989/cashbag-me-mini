package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	TransactionBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		CompanyId      primitive.ObjectID `bson:"companyId"`
		BranchId       primitive.ObjectID `bson:"branchId"`
		UserId         string             `bson:"userId"`
		Amount         float64            `bson:"amount" `
		Commission     float64            `bson:"commission" `
		LoyaltyProgram float64            `bson:"loyaltyProgram" `
		CreateAT       time.Time          `bson:"createAT"`
	}
)
