package grpcuser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	userpb "cashbag-me-mini/proto/models/user"
)

func convertToUserBrief(data *userpb.UserBrief) models.UserBrief {
	var (
		userID, _ = primitive.ObjectIDFromHex(data.Id)
	)

	userBrief := models.UserBrief{
		ID:   userID,
		Name: data.Name,
	}
	return userBrief
}
