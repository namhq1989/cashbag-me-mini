package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"

	"github.com/labstack/echo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CompanyList ...
func CompanyList() []models.CompanyBSON {
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

//GetNameCompanyById func
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

//CompanyCreate func to ....
func CompanyCreate(c echo.Context) *models.CompanyDetail {
	var companyCollection = database.ConnectCol("companies")
	company := new(models.CompanyDetail)
	c.Bind(company)
	company.ID = primitive.NewObjectID()
	_, err := companyCollection.InsertOne(context.TODO(), company)
	if err != nil {
		log.Fatal(err)
	}
	return company
}

//CompanyUpdate func to ...
func CompanyUpdate(c echo.Context) *models.CompanyDetail {
	var companyCollection = database.ConnectCol("companies")
	company := new(models.CompanyDetail)
	if err := c.Bind(company); err != nil {
		log.Println(err)
	}
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":           company.Name,
		"address":        company.Address,
		"balance":        company.Balance,
		"loyaltyProgram": company.LoyaltyProgram,
		"active":         company.Active,
		"createAt":       company.CreateAt,
		"updateAt":       company.UpdateAt,
	}}
	_, err := companyCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return company

}

//CompanyActive  func to ...
func CompanyActive(c echo.Context) *models.CompanyDetail {
	var companyCollection = database.ConnectCol("companies")
	company := new(models.CompanyDetail)
	if err := c.Bind(company); err != nil {
		log.Println(err)
	}
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"active": company.Active,
	}}
	_, err := companyCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return company
}
