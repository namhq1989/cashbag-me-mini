package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//PutBranch ...
func PutBranch(idBranch interface{}, body models.PutBranch) *mongo.UpdateResult {
	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
	)
	filter := bson.M{"_id": idBranch}
	update := bson.M{"$set": bson.M{
		"name":     body.Name,
		"address":  body.Address,
		"active":   body.Active,
		"updateAt": time.Now(),
	}}
	result, err := branchCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

//GetNameBranchById ....
func GetNameBranchById(id interface{}) string {
	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
		result           = struct {
			Name string `bson:"name"`
		}{}
		filter = bson.M{"_id": id}
	)
	err := branchCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Name
}

//GetIdBranchByName .....
func GetIdBranchByName(NameBranch interface{}) primitive.ObjectID {

	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
		result           = struct {
			ID primitive.ObjectID `bson:"_id"`
		}{}
		filter = bson.M{"name": NameBranch}
	)
	err := branchCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.ID
}
