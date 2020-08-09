package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// TransactionAnalyticBSON ...
	TransactionAnalyticBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		CompanyID        primitive.ObjectID `bson:"companyId"`
		BranchID         primitive.ObjectID `bson:"branchId"`
		Date             time.Time          `bson:"date"`
		TotalTransaction int                `bson:"totalTransaction"`
		TotalRevenue     float64            `bson:"totalRevenue" `
		TotalCommission  float64            `bson:"totalCommission" `
		UpdateAt         time.Time          `bson:"updateAt"`
	}

	// TransactionAnalyticDetail ...
	TransactionAnalyticDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		Company          CompanyBrief       `json:"companyId"`
		Branch           BranchBrief        `json:"branchId"`
		Date             time.Time          `json:"date"`
		TotalTransaction int                `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue" `
		TotalCommission  float64            `json:"totalCommission" `
		UpdateAt         time.Time          `json:"updateAt"`
	}
)
