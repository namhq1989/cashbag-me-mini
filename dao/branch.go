package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)


//ListBranch ...
func ListBranch() []models.BranchBSON {
	var (
		branchCollection  = database.ConnectCol("branchs")
		ctx    = context.Background()
		result []models.BranchBSON
	)
	cursor, err := branchCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &result)
	return result
}
