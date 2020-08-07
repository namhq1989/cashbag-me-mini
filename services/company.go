package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CompanyCreate func to ...
func CompanyCreate(body models.CompanyCreate) (models.CompanyBSON ,error){
	company := ConvertBodyToCompanyBSON(body)
	doc,err	:= dao.CompanyCreate(company)
	return doc,err
}

// CompanyList ...
func CompanyList() ([]models.CompanyDetail,error) {
	var (
		result []models.CompanyDetail
	)
	doc ,err := dao.CompanyList()
	for _, item := range doc {
		company := convertToCompanyDetail(item)
		result = append(result, company)
	}
	return result,err
}

// CompanyChangeActiveStatus func to. ...
func CompanyChangeActiveStatus(id primitive.ObjectID, status models.CompanyUpdate) (models.CompanyUpdate,error) {
	doc,err := dao.CompanyChangeActiveStatus(id,status)
	return doc,err
	
}

//PutCompany func ....
func PutCompany(idCompany interface{}, body models.CompanyUpdate) *mongo.UpdateResult {
	result := dao.PutCompany(idCompany, body)
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
		CreatedAt:       x.CreatedAt,
		UpdatedAt:       x.UpdatedAt,
	}
	return result
}

// ConvertBodyToCompanyBSON func ...
func ConvertBodyToCompanyBSON(body models.CompanyCreate) models.CompanyBSON {
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
 