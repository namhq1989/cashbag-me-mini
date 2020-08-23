package dao

import (
	"context"
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

// CompanyChangeActiveStatus ...
func CompanyChangeActiveStatus(companyID primitive.ObjectID, active bool) (err error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	_, err = companyCol.UpdateOne(ctx, companyID, active)

	return err
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

// CompanyUpdateActiveTrue ...
// func CompanyUpdateActiveTrue(id primitive.ObjectID) error{
// 	var (
// 		filter = bson.M{"_id": id}
// 		update = bson.M{"$set": bson.M{
// 			"active":true,
// 		}}
// 	)

// 	// Update
// 	err := CompanyUpdateByID(filter, update)

// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return err
// }

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
