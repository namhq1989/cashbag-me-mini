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

	// UserBrief ...
	UserBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}
)
