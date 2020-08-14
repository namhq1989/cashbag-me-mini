package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// companyAnalyticBSON ...
	companyAnalyticBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		CompanyID        primitive.ObjectID `bson:"companyID"`
		ActiveBranch     int                `bson:"activeBranch"`
		InactiveBranch   int                `bson:"inactiveBranch"`
		TotalTransaction int                `bson:"totalTransaction"`
		TotalRevenue     float64            `bson:"totalRevenue"`
		TotalCommission  float64            `bson:"totalCommission"`
		TotalDebt        float64            `bson:"totalDebt"`
		CountPostpaid    int                `bson:"countPostpaid"`
		TotalUser        int                `bson:"totalUser"`
		UserSilver       int                `bson:"userSilver"`
		USerGolden       int                `bson:"userGolden"`
		UserDiamond      int                `bson:"userDiamond"`
		UpdatedAt        time.Time          `bson:"updatedAt"`
	}

	// companyAnalyticDetail ....
	companyAnalyticDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		CompanyID        primitive.ObjectID `json:"companyID"`
		ActiveBranch     int                `json:"activeBranch"`
		InactiveBranch   int                `json:"inactiveBranch"`
		TotalTransaction int                `json:"totalTransaction"`
		TotalRevenue     float64            `json:"totalRevenue"`
		TotalCommission  float64            `json:"totalCommission"`
		TotalDebt        float64            `json:"totalDebt"`
		CountPostpaid    int                `json:"countPostpaid"`
		TotalUser        int                `json:"totalUser"`
		UserSilver       int                `json:"userSilver"`
		USerGolden       int                `json:"userGolden"`
		UserDiamond      int                `json:"userDiamond"`
		UpdatedAt        time.Time          `json:"updatedAt"`
	}
)
