package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// UserCreate ...
func UserCreate(doc models.UserBSON) (models.UserBSON, error) {
	var (
		collection = database.UserCol()
		ctx        = context.Background()
	)

	//add information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}

	doc.CreatedAt = time.Now()

	//insert one
	_, err := collection.InsertOne(ctx, doc)
	return doc, err

}

// UserUpdate ...
func UserUpdate(id primitive.ObjectID, user models.UserBSON) (models.UserBSON, error) {
	var (
		filter = bson.M{"_id": id}
		update = bson.M{"$set": bson.M{
			"name":      user.Name,
			"address":   user.Address,
			"updatedAt": time.Now(),
		}}
	)

	// Update
	err := UserUpdateByID(filter, update)

	// Get doc
	doc, _ := UserFindByID(id)

	return doc, err

}