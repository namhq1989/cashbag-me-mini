package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (

	// SilverButton ...
	SilverButton struct {
		spending        float64
		commissionLevel float64
	}

	// GoldenButton ....
	GoldenButton struct {
		spending        float64
		commissionLevel float64
	}

	// DiamondButton .....
	DiamondButton struct {
		spending        float64
		commissionLevel float64
	}

	// UserProgram ...
	UserProgram struct {
		ID        primitive.ObjectID `bson:"_id"`
		CompanyID primitive.ObjectID `bson:"companyID"`
		Silver    SilverButton       `bson:"silver"`
		Golden    GoldenButton       `bson:"golden"`
		Diamond   DiamondButton      `bson:"diamond"`
	}
)
