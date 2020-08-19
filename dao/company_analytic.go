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

// CompanyAnalyticList ...
func CompanyAnalyticList() ([]models.CompanyAnalyticBSON, error) {
	var (
		companyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
		result             = make([]models.CompanyAnalyticBSON, 0)
	)

	// Find
	cursor, err := companyAnalyticCol.Find(ctx, bson.M{})

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}

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

// CompanyAnalyticUpdateTransactionProperties ...
func CompanyAnalyticUpdateTransactionProperties(companyAnalytic models.CompanyAnalyticBSON) error {
	var (
		filter = bson.M{"_id": companyAnalytic.ID}
		update = bson.M{"$set": bson.M{
			"totalTransaction": companyAnalytic.TotalTransaction,
			"totalRevenue":     companyAnalytic.TotalRevenue,
			"totalCommission":  companyAnalytic.TotalCommission,
			"totalDebt":        companyAnalytic.TotalDebt,
			"countPostpaid":    companyAnalytic.CountPostpaid,
			"userSilver":       companyAnalytic.UserSilver,
			"userGolden":       companyAnalytic.UserGolden,
			"userDiamond":      companyAnalytic.UserDiamond,
			"updatedAt":        time.Now(),
		}}
	)

	// Update
	err := CompanyAnalyticUpdateByID(filter, update)

	return err
}

// CompanyAnalyticUpdateBranchProperties ...
func CompanyAnalyticUpdateBranchProperties(companyAnalytic models.CompanyAnalyticBSON) error {
	var (
		filter = bson.M{"_id": companyAnalytic.ID}
		update = bson.M{"$set": bson.M{
			"activeBranch":   companyAnalytic.ActiveBranch,
			"inactiveBranch": companyAnalytic.InactiveBranch,
			"updatedAt":      time.Now(),
		}}
	)

	// Update
	err := CompanyAnalyticUpdateByID(filter, update)

	return err
}

// CompanyAnalyticUpdateUserProperties ...
func CompanyAnalyticUpdateUserProperties(companyAnalytic models.CompanyAnalyticBSON) error {
	var (
		filter = bson.M{"_id": companyAnalytic.ID}
		update = bson.M{"$set": bson.M{
			"totalUser": companyAnalytic.TotalUser,
			"updatedAt": time.Now(),
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
