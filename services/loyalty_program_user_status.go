package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

func createLoyaltyProgramUserStatus(currentUserMilestone models.LoyaltyProgramMilestone, currentUserExpense float64, companyID, userID primitive.ObjectID) (err error) {
	loyaltyProgramUserNilStatus := models.LoyaltyProgramUserStatusBSON{
		ID:             primitive.NewObjectID(),
		CompanyID:      companyID,
		UserID:         userID,
		Milestone:      currentUserMilestone,
		CurrentExpense: currentUserExpense,
		CreatedAt:      time.Now(),
	}

	_, err = dao.LoyaltyProgramUserStatusCreate(loyaltyProgramUserNilStatus)
	return
}

func updateLoyaltyProgramUserStatus(beforeUserMilestone models.LoyaltyProgramMilestone, currentUserExpense float64, userID primitive.ObjectID) (err error) {
	var (
		filter = bson.M{
			"userID":    userID,
			"milestone": beforeUserMilestone,
		}
		update = bson.M{"$set": bson.M{
			"currentExpense": currentUserExpense,
			"updatedAt":      time.Now(),
		}}
	)

	err = dao.LoyaltyProgramUserStatusUpdate(filter, update)
	return
}
