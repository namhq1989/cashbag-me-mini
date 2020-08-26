package services

import (
	"cashbag-me-mini/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserCreatePayloadToBSON ...
func userCreatePayloadToBSON(body models.UserCreatePayload) models.UserBSON {

	result := models.UserBSON{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		Address:   body.Address,
		CreatedAt: time.Now(),
	}
	return result
}
