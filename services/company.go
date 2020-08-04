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

//ListCompany to ...
func ListCompany() []models.CompanyDetail {
	var (
		result []models.CompanyDetail
	)
	companies := dao.ListCompany()
	for _, item := range companies {
		company := convertToCompanyDetail(item)
		result = append(result, company)
	}
	return result
}

//PatchCompany func to. ...
func PatchCompany(idCompany interface{}) *mongo.UpdateResult {
	result := dao.PatchCompany(idCompany)
	return result
}

//PutCompany func ....
func PutCompany(idCompany interface{},body models.PutCompany) *mongo.UpdateResult {
	result := dao.PutCompany(idCompany,body)
	return result
}

//convertToCompanyDetail to ..
func convertToCompanyDetail(x models.CompanyBSON) models.CompanyDetail {
	result := models.CompanyDetail{
		ID:             x.ID,
		Name:           x.Name,
		Address:        x.Address,
		Balance:        x.Balance,
		LoyaltyProgram: x.LoyaltyProgram,
		Active:         x.Active,
		CreateAt:       x.CreateAt,
		UpdateAt:       x.UpdateAt,
	}
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

//GetNameCompanyById func ....
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

//GetIdCompanyByName func ....
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
