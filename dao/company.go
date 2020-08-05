package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateCompany func to ....
func CreateCompany(company interface{}) *mongo.InsertOneResult {
	var companyCollection = database.ConnectCol("companies")
	result, err := companyCollection.InsertOne(context.TODO(), company)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//ListCompany ...
func ListCompany() []models.CompanyBSON {
	var (
		companyCollection = database.ConnectCol("companies")
		ctx               = context.Background()
		result            []models.CompanyBSON
	)
	cursor, err := companyCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &result)
	return result
}

//PatchCompany  func to ...
func PatchCompany(idCompany interface{}) *mongo.UpdateResult {
	var (
		companyCollection = database.ConnectCol("companies")
		ctx               = context.Background()
	)
	filter := bson.M{"_id": idCompany}
	update := bson.M{"$set": bson.M{"active": true}}
	result, err := companyCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//PutCompany  func to ...
func PutCompany(idCompany interface{},body models.PutCompany) *mongo.UpdateResult {
	var companyCollection = database.ConnectCol("companies")
	filter := bson.M{"_id": idCompany}
	update := bson.M{"$set": bson.M{
		"name":           body.Name,
		"address":        body.Address,
		"active":         body.Active,
		"balance":        body.Balance,
		"loyaltyProgram": body.LoyaltyProgram,
		"updateAt":       time.Now(),
	}}
	result, err := companyCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//GetNameCompanyById ....
func GetNameCompanyById(id interface{}) string {
	var (
		companyCollection = database.ConnectCol("companies")
		ctx               = context.Background()
		result            = struct {
			Name string `bson:"name"`
		}{}
		filter = bson.M{"_id": id}
	)
	err := companyCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Name
}

// GetIdCompanyByName ...
func GetIdCompanyByName(nameCompany interface{}) primitive.ObjectID {
	var (
		//db                = database.Connectdb("CashBag")
		companyCollection = database.ConnectCol("companies")
		ctx               = context.Background()
		result            = struct {
			ID primitive.ObjectID `bson:"_id"`
		}{}
		filter = bson.M{"name": nameCompany}
	)
	err := companyCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.ID
}
