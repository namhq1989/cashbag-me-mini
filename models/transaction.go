package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	//TransactionBSON ...
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
	//TransactionDetail ...
	TransactionDetail struct {
		ID             primitive.ObjectID `json:"_id"`
		CompanyID      primitive.ObjectID `json:"companyId"`
		BranchID       primitive.ObjectID `json:"branchId"`
		User           string             `json:"user"`
		Amount         float64            `json:"amount"`
		Commission     float64            `json:"commission"`
		LoyaltyProgram float64            `json:"loyaltyprogram"`
		CreateAt       time.Time          `json:"createAt"`
	}
	//PostTransaction is a  struct of body request
	PostTransaction struct {
		NameCompany    string  `json:"nameCompany" valid:"stringlength(3|30),type(string)"`
		NameBranch     string  `json:"nameBranch" valid:"stringlength(3|30),type(string)"`
		User           string  `json:"user" valid:"stringlength(3|30),type(string)"`
		Amount         float64 `json:"amount" valid:"required"`
		Commission float64     `json:"commission"`
	}
)
