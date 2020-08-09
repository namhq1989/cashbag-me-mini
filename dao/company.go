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

// CompanyCreate ...
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
		doc        = make([]models.CompanyBSON, 0)
	)

	// Find
	cursor, err := companyCol.Find(ctx, bson.M{})

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &doc)

	return doc, err
}

// CompanyUpdate ...
func CompanyUpdate(id primitive.ObjectID, company models.CompanyBSON) (models.CompanyBSON, error) {
	var (
		filter = bson.M{"_id": id}
		update = bson.M{"$set": bson.M{
			"name":           company.Name,
			"address":        company.Address,
			"active":         company.Active,
			"balance":        company.Balance,
			"loyaltyProgram": company.LoyaltyProgram,
			"updatedAt":      time.Now(),
		}}
	)

	// Update
	err := CompanyUpdateByID(filter, update)

	// Get doc
	doc, _ := CompanyFindByID(id)

	return doc, err
}

// CompanyChangeActiveStatus ...
func CompanyChangeActiveStatus(id primitive.ObjectID) (models.CompanyBSON, error) {
	var (
		filter = bson.M{"_id": id}
		doc, _ = CompanyFindByID(id)
		update = bson.M{"$set": bson.M{"active": !(doc.Active)}}
	)

	// Update
	err := CompanyUpdateByID(filter, update)

	return doc, err
}

// CompanyFindByID ...
func CompanyFindByID(id primitive.ObjectID) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     models.CompanyBSON
		filter     = bson.M{"_id": id}
	)

	// Find
	err := companyCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// CompanyUpdateBalance ...
func CompanyUpdateBalance(id primitive.ObjectID, balance float64) {
	var (
		filter = bson.M{"_id": id}
		update = bson.M{"$set": bson.M{
			"balance": balance,
		}}
	)

	// Update
	err := CompanyUpdateByID(filter, update)

	if err != nil {
		log.Println(err)
	}
}

// CompanyUpdateByID ...
func CompanyUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		CompanyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	// Update
	_, err := CompanyCol.UpdateOne(ctx, filter, updateData)

	return err
}
