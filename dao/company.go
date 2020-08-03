package dao

import (
	"cashbag-me-mini/modules/database"
	"context"
	"log"

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
