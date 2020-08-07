package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// CompanyCreate ....
func CompanyCreate(doc models.CompanyBSON) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	//Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.UpdatedAt = time.Now()

	_, err := companyCol.InsertOne(ctx, doc)
	return doc, err
}

// CompanyList  ...
func CompanyList() ([]models.CompanyBSON ,error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = make([]models.CompanyBSON,0)
	)

	cursor, err := companyCol.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	cursor.All(ctx, &result)
	return result, err
}

//CompanyChangeActiveStatus func to ...
func CompanyChangeActiveStatus(id primitive.ObjectID,status models.CompanyUpdate) (models.CompanyUpdate,error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		filter = bson.M{"_id": id}
		update = bson.M{"$set": bson.M{"active": !(status.Active)}}
	)
	
	_, err := companyCol.UpdateOne(ctx, filter, update)
	log.Println("status 1",status)
	return status,err
}

//PutCompany  func to ...
func PutCompany(idCompany interface{}, body models.CompanyUpdate) *mongo.UpdateResult {
	var companyCol = database.CompanyCol()
	filter := bson.M{"_id": idCompany}
	update := bson.M{"$set": bson.M{
		"name":           body.Name,
		"address":        body.Address,
		"active":         body.Active,
		"balance":        body.Balance,
		"loyaltyProgram": body.LoyaltyProgram,
		"updateAt":       time.Now(),
	}}
	result, err := companyCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//GetNameCompanyById ....
func GetNameCompanyById(id interface{}) string {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = struct {
			Name string `bson:"name"`
		}{}
		filter = bson.M{"_id": id}
	)
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Name
}

// GetIdCompanyByName ...
func GetIdCompanyByName(nameCompany interface{}) primitive.ObjectID {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = struct {
			ID primitive.ObjectID `bson:"_id"`
		}{}
		filter = bson.M{"name": nameCompany}
	)
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.ID
}

//GetLoyaltyProgramByCompany func ...
func GetIFCompanyByName(NameCompany interface{}) models.IFCompany {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = models.IFCompany{}
		filter     = bson.M{"name": NameCompany}
	)
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//UpdateBalance ...
func UpdateBalance(idCompany interface{}, balance float64) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)
	filter := bson.M{"_id": idCompany}
	update := bson.M{"$set": bson.M{
		"balance": balance,
	}}
	log.Println(balance)
	_, err := companyCol.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
