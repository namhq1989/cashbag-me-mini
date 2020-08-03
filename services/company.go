package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateCompany func to ...
func CreateCompany(body models.PostCompany) *mongo.InsertOneResult {
	var company models.CompanyBSON
	company = ConvertBodyToCompanyBSON(body)
	company.ID = primitive.NewObjectID()
	company.CreateAt = time.Now()
	result := dao.CreateCompany(company)
	return result
}

// ConvertBodyToCompanyBSON func ...
func ConvertBodyToCompanyBSON(body models.PostCompany) models.CompanyBSON {
	result := models.CompanyBSON{
		Name:    body.Name,
		Address: body.Address,
		Active:  body.Active,
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

//GetIdCompanyByName ...
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