package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// AnalyticChartBSON ....
	AnalyticChartBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		CompanyID        primitive.ObjectID `bson:"companyID"`
		Date             time.Time          `bson:"date"`
		TotalTransaction int                `bson:"totalTransaction"`
		TotalRevenue     float64            `bson:"totalRevenue" `
		TotalCommission  float64            `bson:"totalCommission" `
		UpdateAt         time.Time          `bson:"updateAt"`
	}

	// AnalyticChartDetail ...
	AnalyticChartDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		CompanyID        primitive.ObjectID `json:"companyID"`
		Date             time.Time          `json:"date"`
		TotalTransaction int                `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue" `
		TotalCommission  float64            `json:"totalCommission" `
		UpdateAt         time.Time          `json:"updateAt"`
	}
)
