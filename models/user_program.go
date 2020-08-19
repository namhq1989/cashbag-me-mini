package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (

	// SilverButton ...
	SilverButton struct {
		Spending   float64 `bson:"spending"`
		Commission float64 `bson:"commission"`
	}

	// GoldenButton ..
	GoldenButton struct {
		Spending   float64 `bson:"spending"`
		Commission float64 `bson:"commission"`
	}

	// DiamondButton ...
	DiamondButton struct {
		Spending   float64 `bson:"spending"`
		Commission float64 `bson:"commission"`
	}
	
	// UserProgramBSON ...
	UserProgramBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		CompanyID primitive.ObjectID `bson:"companyID"`
		Silver    SilverButton       `bson:"silver"`
		Golden    GoldenButton       `bson:"golden"`
		Diamond   DiamondButton      `bson:"diamond"`
	}
	
	// UserProgramCreatePayload ...
	UserProgramCreatePayload struct {
		CompanyID         string  `json:"companyID"`
		SilverSpending    float64 `json:"silverSpending"`
		SilverCommission  float64 `json:"silverCommission"`
		GoldenSpending    float64 `json:"goldenSpending"`
		GoldenCommission  float64 `json:"goldenCommission"`
		DiamondSpending   float64 `json:"diamondSpending"`
		DiamondCommission float64 `json:"diamondCommission"`
	}
)
