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

// TransactionAnalytic ...
func TransactionAnalytic(date time.Time) ([]models.TransactionAnalyticBSON, error) {
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

// TransactionAnalyticHandle ...
func TransactionAnalyticHandle(transaction models.TransactionBSON) {
	tranAnalytic, check := TransactionAnalyticFilter(transaction)
	if check == false {
		TransactionAnalyticCreate(transaction)
	} else {
		TransactionAnalyticUpdate(tranAnalytic, transaction)
	}
}

// TransactionAnalyticFilter ...
func TransactionAnalyticFilter(transaction models.TransactionBSON) (models.TransactionAnalyticBSON, bool) {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
		startOfDate            = BeginningOfDay(transaction.CreatedAt)
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
		return tranAnalytic, false
	}

	return tranAnalytic, true
}

// TransactionAnalyticCreate ...
func TransactionAnalyticCreate(transaction models.TransactionBSON) {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
		startOfDate            = BeginningOfDay(transaction.CreatedAt)
	)

	// Set transactionAnalytic
	transactionAnalytic := models.TransactionAnalyticBSON{
		ID:               primitive.NewObjectID(),
		CompanyID:        transaction.CompanyID,
		BranchID:         transaction.BranchID,
		Date:             startOfDate,
		TotalTransaction: 1,
		TotalRevenue:     transaction.Amount,
		TotalCommission:  transaction.Commission,
		UpdateAt:         time.Now(),
	}

	// Create
	_, err := transactionAnalyticCol.InsertOne(ctx, transactionAnalytic)
	if err != nil {
		log.Println(err)
	}
}

// TransactionAnalyticUpdate ...
func TransactionAnalyticUpdate(transactionAnalytic models.TransactionAnalyticBSON, transaction models.TransactionBSON) {

	// Set for update Transaction Analytic
	transactionAnalytic.TotalTransaction++
	transactionAnalytic.TotalRevenue += transaction.Amount
	transactionAnalytic.TotalCommission += transaction.Commission

	// Update Transaction Analytic
	filter := bson.M{"_id": transactionAnalytic.ID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": transactionAnalytic.TotalTransaction,
		"totalRevenue":     transactionAnalytic.TotalRevenue,
		"totalCommission":  transactionAnalytic.TotalCommission,
		"updateAt":         time.Now(),
	}}

	// Update
	err := TransactionAnalyticUpdateByID(filter, update)
	if err != nil {
		log.Println(err)
	}
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

// BeginningOfDay ...
func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
