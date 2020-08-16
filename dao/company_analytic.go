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

// CompanyAnalyticCreate ...
func CompanyAnalyticCreate(companyID primitive.ObjectID) error {
	var (
		companyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
	)

	// Set CompanyAnalytic
	CompanyAnalytic := models.CompanyAnalyticBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: companyID,
		UpdatedAt: time.Now(),
	}

	// Create
	_, err := companyAnalyticCol.InsertOne(ctx, CompanyAnalytic)
	if err != nil {
		log.Println(err)
	}
	return err
}

// CompanyAnalyticUpdate ...
func CompanyAnalyticUpdate(CompanyAnalytic models.CompanyAnalyticBSON, transaction models.TransactionBSON) {

	// Set for update Transaction Analytic
	CompanyAnalytic.TotalTransaction++
	CompanyAnalytic.TotalRevenue += transaction.Amount
	CompanyAnalytic.TotalCommission += transaction.Commission

	// Update Transaction Analytic
	filter := bson.M{"_id": CompanyAnalytic.ID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": CompanyAnalytic.TotalTransaction,
		"totalRevenue":     CompanyAnalytic.TotalRevenue,
		"totalCommission":  CompanyAnalytic.TotalCommission,
		"updateAt":         time.Now(),
	}}

	// Update
	err := CompanyAnalyticUpdateByID(filter, update)
	if err != nil {
		log.Println(err)
	}
}

// CompanyAnalyticUpdateBranchProperties ...
func CompanyAnalyticUpdateBranchProperties(companyAnalytic models.CompanyAnalyticBSON) error {
	var (
		filter = bson.M{"_id": companyAnalytic.ID}
		update = bson.M{"$set": bson.M{
			"activeBranch":   companyAnalytic.ActiveBranch,
			"inactiveBranch": companyAnalytic.InactiveBranch,
			"updateAt":       time.Now(),
		}}
	)

	// Update
	err := CompanyAnalyticUpdateByID(filter, update)

	return err
}

// CompanyAnalyticUpdateByID ...
func CompanyAnalyticUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		CompanyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
	)

	_, err := CompanyAnalyticCol.UpdateOne(ctx, filter, updateData)

	return err
}

// CompanyAnalyticFindByCompanyID ...
func CompanyAnalyticFindByCompanyID(id primitive.ObjectID) (models.CompanyAnalyticBSON, error) {
	var (
		companyCol = database.CompanyAnalyticCol()
		ctx        = context.Background()
		result     models.CompanyAnalyticBSON
		filter     = bson.M{"companyID": id}
	)

	// Find
	err := companyCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}
