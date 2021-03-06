package grpcuser

import (
	"context"
	"errors"
	"time"

	"cashbag-me-mini/models"
	userpb "cashbag-me-mini/proto/models/user"
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

	// Convert to user brief
	userBrief := convertToUserBrief(result.UserBrief)

	return userBrief, nil
}
