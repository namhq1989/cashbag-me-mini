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
<<<<<<< HEAD
		CreatedAt      time.Time          `bson:"createdAt"`
=======
		CreatedAt       time.Time          `bson:"createdAt"`
>>>>>>> 35bc62c3407eecf25d9e630ee81dc51e89f7e4bb
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
<<<<<<< HEAD
		CreatedAt      time.Time          `json:"createdAt"`
	}

	// TransactionCreatePayload is a  struct of body request
	TransactionCreatePayload struct {
		CompanyID string  `json:"companyID"`
		BranchID  string  `json:"branchID"`
		User      string  `json:"user" valid:"stringlength(3|100),type(string)"`
		Amount    float64 `json:"amount"`
=======
		CreatedAt       time.Time          `json:"createdAt"`
	}
	// TransactionCreatePayload is a  struct of body request
	TransactionCreatePayload struct {
		CompanyID      primitive.ObjectID `json:"companyId"`
		BranchID       primitive.ObjectID `json:"branchId"`
		User           string  `json:"user" valid:"stringlength(3|30),type(string)"`
		Amount         float64 `json:"amount"`
>>>>>>> 35bc62c3407eecf25d9e630ee81dc51e89f7e4bb
	}
)
