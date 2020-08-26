package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// CompanyanalyticMember ...
	CompanyAnalyticMember struct {
		ID    string
		Total int
	}

	// CompanyAnalyticBSON ...
	CompanyAnalyticBSON struct {
		ID               primitive.ObjectID      `bson:"_id"`
		CompanyID        primitive.ObjectID      `bson:"companyID"`
		ActiveBranch     int                     `bson:"activeBranch"`
		TotalBranch      int                     `bson:"totalBranch"`
		TotalTransaction int                     `bson:"totalTransaction"`
		TotalRevenue     float64                 `bson:"totalRevenue"`
		TotalCommission  float64                 `bson:"totalCommission"`
		TotalDebt        float64                 `bson:"totalDebt"`
		CountPostpaid    int                     `bson:"countPostpaid"`
		Members          []CompanyAnalyticMember `bson:"members"`
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
		Members          []CompanyAnalyticMember `json:"members"`
		UpdatedAt        time.Time               `json:"updatedAt"`
	}
)
