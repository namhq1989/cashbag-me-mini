package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// UserBSON ....
	UserBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		CompanyID primitive.ObjectID `bson:"companyID"`
		Name      string             `bson:"name"`
		Address   string             `bson:"address"`
		Level     string             `bson:"level"`
		Spending  float64            `bson:"spending"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	// UserCreatePayload ...
	UserCreatePayload struct {
		CompanyID string `json:"companyID"`
		Name      string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address   string `json:"address" valid:"stringlength(3|30),type(string)"`
	}

	// UserUpdatePayload ...
	UserUpdatePayload struct {
		CompanyID string  `json:"companyID"`
		Name      string  `json:"name" valid:"stringlength(3|30),type(string)"`
		Address   string  `json:"address" valid:"stringlength(3|30),type(string)"`
		Level     string  `json:"level"`
		Spending  float64 `json:"spending"`
	}
)
