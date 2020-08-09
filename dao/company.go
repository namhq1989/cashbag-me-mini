package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// CompanyCreate ....
func CompanyCreate(doc models.CompanyBSON) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	// Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.CreatedAt = time.Now()

	// Insert
	_, err := companyCol.InsertOne(ctx, doc)

	return doc, err
}

// CompanyList  ...
func CompanyList() ([]models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = make([]models.CompanyBSON, 0)
	)

	// Find
	cursor, err := companyCol.Find(ctx, bson.M{})

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}

//CompanyChangeActiveStatus func to ...
func CompanyChangeActiveStatus(id primitive.ObjectID) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		filter     = bson.M{"_id": id}
		update     = bson.M{"$set": bson.M{"active": true}}
	)

	doc := CompanyFindByID(id)

	_, err := companyCol.UpdateOne(ctx, filter, update)
	return doc, err
}

//CompanyUpdate ...
func CompanyUpdate(CompanyID primitive.ObjectID, body models.CompanyUpdatePayload) (models.CompanyBSON, error) {
	// Set filter
	filter := bson.M{"_id": CompanyID}

	// Add information for update
	updateData := bson.M{"$set": bson.M{
		"name":           body.Name,
		"address":        body.Address,
		"active":         body.Active,
		"balance":        body.Balance,
		"loyaltyProgram": body.LoyaltyProgram,
		"updateAt":       time.Now(),
	}}

	err := CompanyUpdateByID(filter, updateData)
	doc := CompanyFindByID(CompanyID)

	return doc, err
}

// CompanyUpdateByID ...
func CompanyUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	_, err := companyCol.UpdateOne(ctx, filter, updateData)

	return err
}

//CompanyFindByID func ...
func CompanyFindByID(id primitive.ObjectID) models.CompanyBSON {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = models.CompanyBSON{}
		filter     = bson.M{"_id": id}
	)

	//Find
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err!=nil{
		log.Println(err)
	}

	return result
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
