package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

// BranchCreate ...
func BranchCreate(doc models.BranchBSON) (models.BranchBSON, error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	// Insert
	_, err := branchCol.InsertOne(ctx, doc)

	return doc, err
}

// BranchUpdateByID ...
func BranchUpdateByID(filter bson.M, updateData bson.M) (err error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	_, err = branchCol.UpdateOne(ctx, filter, updateData)

	return err
}

// BranchFindByID ...
func BranchFindByID(branchID primitive.ObjectID) (models.BranchBSON, error) {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
		filter    = bson.M{"_id": branchID}
		result    models.BranchBSON
	)

	// Find
	err := branchCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}
