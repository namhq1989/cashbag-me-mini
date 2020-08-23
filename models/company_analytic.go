package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// CompanyanalyticMember ...
	CompanyanalyticMember struct {
		ID    string
		Total int
	}

	// CompanyAnalyticBSON ...
	CompanyAnalyticBSON struct {
		ID               primitive.ObjectID      `bson:"_id"`
		CompanyID        primitive.ObjectID      `bson:"companyID"`
		ActiveBranch     int                     `bson:"activeBranch"`
		InactiveBranch   int                     `bson:"inactiveBranch"`
		TotalTransaction int                     `bson:"totalTransaction"`
		TotalRevenue     float64                 `bson:"totalRevenue"`
		TotalCommission  float64                 `bson:"totalCommission"`
		TotalDebt        float64                 `bson:"totalDebt"`
		CountPostpaid    int                     `bson:"countPostpaid"`
		TotalUser        int                     `bson:"totalUser"`
		Members          []CompanyanalyticMember `bson:"members"`
		UpdatedAt        time.Time               `bson:"updatedAt"`
	}

	// CompanyAnalyticDetail ....
	CompanyAnalyticDetail struct {
		ID               primitive.ObjectID      `json:"_id"`
		Company          string                  `json:"company"`
		Branch           string                  `json:"branch"`
		TotalTransaction int                     `json:"totalTransaction"`
		TotalRevenue     float64                 `json:"totalRevenue"`
		TotalCommission  float64                 `json:"totalCommission"`
		TotalDebt        float64                 `json:"totalDebt"`
		CountPostpaid    int                     `json:"countPostpaid"`
		TotalUser        int                     `json:"totalUser"`
		Members          []CompanyanalyticMember `json:"members"`
		UpdatedAt        time.Time               `json:"updatedAt"`
	}
)
