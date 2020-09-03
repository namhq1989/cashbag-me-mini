package grpcuser

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	userpb "cashbag-me-mini/proto"
)

// GetUserBriefByID ...
func GetUserBriefByID(userID string) (UserBrief models.UserBrief, err error) {
	clientConn, client := CreateClient()
	defer clientConn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call
	result, err := client.GetUserBriefByID(ctx, &userpb.GetUserBriefByIDRequest{Id: userID})
	if err != nil {
		err = errors.New("Khong the get user brief by id")
		return
	}

	userBrief := convertToUserBrief(result.UserBrief)

	return userBrief, nil
}

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
