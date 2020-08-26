package dao

import (
	"cashbag-me-mini/util"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// TransactionAnalyticList ...
func TransactionAnalyticList(date time.Time) ([]models.TransactionAnalyticBSON, error) {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
		result                 = make([]models.TransactionAnalyticBSON, 0)
		filter                 = bson.M{
			"date": date,
		}
	)

	//Find
	cursor, err := transactionAnalyticCol.Find(ctx, filter)

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &result)

	return result, err
}

// TransactionAnalyticFilterByDate ...
func TransactionAnalyticFilterByDate(transaction models.TransactionBSON) (models.TransactionAnalyticBSON, error) {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
		startOfDate            = util.BeginningOfDay(transaction.CreatedAt)
		tranAnalytic           models.TransactionAnalyticBSON
		filter                 = bson.M{
			"companyID": transaction.CompanyID,
			"branchID":  transaction.BranchID,
			"date":      startOfDate,
		}
	)

	// Find
	err := transactionAnalyticCol.FindOne(ctx, filter).Decode(&tranAnalytic)
	if err != nil {
		return tranAnalytic, err
	}

	return tranAnalytic, err
}

// TransactionAnalyticCreate ...
func TransactionAnalyticCreate(transactionAnalytic models.TransactionAnalyticBSON) error {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
	)

	// Create
	_, err := transactionAnalyticCol.InsertOne(ctx, transactionAnalytic)
	return err
}

// TransactionAnalyticUpdateByID ...
func TransactionAnalyticUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
	)

	_, err := transactionAnalyticCol.UpdateOne(ctx, filter, updateData)

	return err
}
