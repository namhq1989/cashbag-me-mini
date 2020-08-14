package services

import(
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// UserCreate ...
func UserCreate(body models.UserCreatePayload) (models.UserBSON,error){
	var (
		user =userCreatePayloadToBSON(body)
	)

	//Create user
	doc,err := dao.UserCreate(user)

	return doc,err
}

//UserUpdate ...
func UserUpdate(id primitive.ObjectID,body models.UserUpdatePayload) (models.UserBSON,error){
	var (
		user  = userUpdatePayloadToBSON(body)
	)

	// Update user
	doc,err :=dao.UserUpdate(id,user)

	return doc,err
}



