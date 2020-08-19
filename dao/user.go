package dao

import (
	"context"
	"log"
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

// UserFindByID ...
func UserFindByID(id primitive.ObjectID) (models.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		result  models.UserBSON
		filter  = bson.M{"_id": id}
	)

	// Find
	err := userCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// UserUpdateByID ...
func UserUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// Update
	_, err := userCol.UpdateOne(ctx, filter, updateData)

	return err
}

// UserUpdateSpendingAndLevel ...
func UserUpdateSpendingAndLevel(id primitive.ObjectID,level string, spending float64) error {
	var (
		filter = bson.M{"_id": id}
		update = bson.M{"$set": bson.M{
			"level":level,
			"spending": spending,
		}}
	)

	// Update
	err := UserUpdateByID(filter, update)

	if err != nil {
		log.Println(err)
	}
	return err
}

