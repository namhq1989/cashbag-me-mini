package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"time"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// BranchList ...
func BranchList() ([]models.BranchBSON, error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
		result    = make([]models.BranchBSON, 0)
	)

	// Find
	cursor, err := branchCol.Find(ctx, bson.M{})

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}

// BranchCreate ....
func BranchCreate(doc models.BranchBSON) (models.BranchBSON, error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	// Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.CreatedAt = time.Now()

	// Insert
	_, err := branchCol.InsertOne(ctx, doc)

	return doc, err
}

//BranchChangeActiveStatus func to ...
func BranchChangeActiveStatus(branchID primitive.ObjectID) (models.BranchBSON, error) {
	var (
		active bool
		filter = bson.M{"_id": branchID}
	)

	// Find Branch
	doc := BranchDocByID(branchID)

	// Change Active status
	active = !(doc.Active)
	update := bson.M{"$set": bson.M{"active": active}}
	err := BranchUpdateByID(filter, update)

	return doc, err
}

// BranchUpdate ...
func BranchUpdate(branchID primitive.ObjectID, body models.BranchBSON) (models.BranchBSON, error) {
	var (
		filter     = bson.M{"_id": branchID}
		updateData = bson.M{"$set": bson.M{
			"name":     body.Name,
			"address":  body.Address,
			"active":   body.Active,
			"updateAt": time.Now(),
		}}
	)

	// Update
	err := BranchUpdateByID(filter, updateData)

	// Get doc
	doc := BranchDocByID(branchID)

	return doc, err
}

// BranchUpdateByID ...
func BranchUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	_, err := branchCol.UpdateOne(ctx, filter, updateData)

	return err
}

// BranchDocByID ...
func BranchDocByID(branchID primitive.ObjectID) models.BranchBSON {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
		filter    = bson.M{"_id": branchID}
		result    models.BranchBSON
	)

	err := branchCol.FindOne(ctx, filter).Decode(&result)
	if err!=nil{
		return result
	}
	return result
}

// BranchValidateID ...
func BranchValidateID(branchID primitive.ObjectID) bool {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
		filter    = bson.M{"_id": branchID}
	)

	err := branchCol.FindOne(ctx, filter)
	if err != nil{
		return false
	}
	return true
}
