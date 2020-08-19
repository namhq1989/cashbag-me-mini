package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// UserProgramCreate ...
func UserProgramCreate(doc models.UserProgramBSON) (models.UserProgramBSON, error) {
	var (
		userProgramCol = database.UserProgramCol()
		ctx            = context.Background()
	)

	//add information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}

	//insert
	_, err := userProgramCol.InsertOne(ctx, doc)

	return doc, err
}

// UserProgramFindByCompanyID ...
func UserProgramFindByCompanyID(id primitive.ObjectID) (models.UserProgramBSON, error) {
	var (
		userProgramCol = database.UserProgramCol()
		ctx            = context.Background()
		result         models.UserProgramBSON
		filter         = bson.M{"companyID": id}
	)

	// Find
	err := userProgramCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}
