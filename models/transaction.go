package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// TransactionBSON ....
	TransactionBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		CompanyID      primitive.ObjectID `bson:"companyID"`
		BranchID       primitive.ObjectID `bson:"branchID"`
		User           string             `bson:"user"`
		Amount         float64            `bson:"amount"`
		Commission     float64            `bson:"commission"`
		LoyaltyProgram float64            `bson:"loyaltyprogram"`
		UserProgram    float64            `bson:"userProgram"`
		Postpaid       bool               `bson:"postpaid"`
		CreatedAt      time.Time          `bson:"createdAt"`
	}

	//TransactionDetail ...
	TransactionDetail struct {
		ID             primitive.ObjectID `json:"_id"`
		CompanyID      primitive.ObjectID `json:"companyID"`
		BranchID       primitive.ObjectID `json:"branchID"`
		User           string             `json:"user"`
		Amount         float64            `json:"amount"`
		Commission     float64            `json:"commission"`
		LoyaltyProgram float64            `json:"loyaltyprogram"`
		UserProgram    float64            `json:"userProgram"`
		Postpaid       bool               `json:"postpaid"`
		CreatedAt      time.Time          `json:"createdAt"`
	}

	// TransactionCreatePayload is a  struct of body request
	TransactionCreatePayload struct {
		CompanyID string  `json:"companyID"`
		BranchID  string  `json:"branchID"`
		User      string  `json:"user" valid:"stringlength(3|100),type(string)"`
		Amount    float64 `json:"amount"`
	}
)
