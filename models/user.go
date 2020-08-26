package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// UserBSON ....
	UserBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		Name      string             `bson:"name"`
		Address   string             `bson:"address"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	// UserCreatePayload ...
	UserCreatePayload struct {
		Name    string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address string `json:"address" valid:"stringlength(3|30),type(string)"`
	}

	// UserUpdatePayload ...
	UserUpdatePayload struct {
		Name    string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address string `json:"address" valid:"stringlength(3|30),type(string)"`
	}
)
