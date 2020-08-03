package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//ListBranch ...
func ListBranch() []models.BranchBSON {
	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
		result           []models.BranchBSON
	)
	cursor, err := branchCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &result)
	return result
}

//CreateBranch ...
func CreateBranch(branch interface{}) *mongo.InsertOneResult {
	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
	)
	result, err := branchCollection.InsertOne(ctx, branch)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

//PatchBranch ...
func PatchBranch(idBranch interface{}) *mongo.UpdateResult {
	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
	)
	filter := bson.M{"_id": idBranch}
	update := bson.M{"$set": bson.M{"active": true}}
	result, err := branchCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
