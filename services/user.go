package services

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// UserCreate ...
func UserCreate(body models.UserCreatePayload) (models.UserBSON, error) {
	var (
		user = userCreatePayloadToBSON(body)
	)

	//Create user
	doc, err := dao.UserCreate(user)
	if err != nil {
		err = errors.New("Khong the tao user")
		return doc, err
	}

	return doc, err
}

//UserUpdate ...
func UserUpdate(userID primitive.ObjectID, body models.UserUpdatePayload) (models.UserBSON, error) {
	var (
		// Prepare update  data
		filter     = bson.M{"_id": userID}
		updateData = bson.M{"$set": bson.M{
			"name":      body.Name,
			"address":   body.Address,
			"updatedAt": time.Now(),
		}}
	)

	// Update User
	err := dao.UserUpdateByID(filter, updateData)
	doc, _ := dao.UserFindByID(userID)

	return doc, err

}
