package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (

	// Button .....
	Button struct {
		spending        float64 `bson:"spending"`
		commissionLevel float64 `bson:"commissionLevel"`
	}

	// UserProgramBSON ...
	UserProgramBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		CompanyID primitive.ObjectID `bson:"companyID"`
		Silver    Button             `bson:"silver"`
		Golden    Button             `bson:"golden"`
		Diamond   Button             `bson:"diamond"`
	}
	// UserProgramCreatePayload ...
	UserProgramCreatePayload struct {
		CompanyID         primitive.ObjectID `json:"companyID"`
		SilverSpending    float64            `json:"silverSpending"`
		SilverCommission  float64            `json:"silverCommission"`
		GoldenSpending    float64            `json:"goldenSpending"`
		GoldenCommission  float64            `json:"goldenCommission"`
		DiamondSpending   float64            `json:"diamondSpending"`
		DiamondCommission float64            `json:"diamondCommission"`
	}
)
