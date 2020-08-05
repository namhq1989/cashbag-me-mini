package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TranAnalytic ...
func TranAnalytic(date interface{}) []models.TranAnalyticBSON {
	var (
		TranAnalyticColl = database.ConnectCol("transactionAnalytic")
		ctx              = context.Background()
		result           []models.TranAnalyticBSON
		filter           = bson.M{
			"date": date,
		}
	)
	cursor, err := TranAnalyticColl.Find(ctx, filter)
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &result)
	return result
}

//HandleTranAnalytic ...
func HandleTranAnalytic(transaction models.TransactionBSON) {
	tranAnalytic, check := FindTranAnalytic(transaction)
	if check == false {
		CreateTranAnalytic(transaction)
	} else {
		UpdateTranAnalytic(tranAnalytic, transaction)
	}

}

//FindTranAnalytic ...
func FindTranAnalytic(transaction models.TransactionBSON) (models.TranAnalyticBSON, bool) {
	var (
		tranAnalyticColl = database.ConnectCol("transactionAnalytic")
		ctx              = context.Background()
		startOfDate      = BeginningOfDay(transaction.CreateAT)
		filter           = bson.M{
			"companyId": transaction.CompanyId,
			"branchId":  transaction.BranchId,
			"date":      startOfDate,
		}
		tranAnalytic models.TranAnalyticBSON
	)
	err := tranAnalyticColl.FindOne(ctx, filter).Decode(&tranAnalytic)
	if err != nil {
		return tranAnalytic, false
	}
	return tranAnalytic, true
}

//CreateTranAnalytic ...
func CreateTranAnalytic(transaction models.TransactionBSON) {
	var (
		branchCollection = database.ConnectCol("transactionAnalytic")
		ctx              = context.Background()
		startOfDate      = BeginningOfDay(transaction.CreateAT)
		tranAnalytic     = models.TranAnalyticBSON{
			ID:               primitive.NewObjectID(),
			CompanyId:        transaction.CompanyId,
			BranchId:         transaction.BranchId,
			Date:             startOfDate,
			TotalTransaction: 1,
			TotalRevenue:     transaction.Amount,
			TotalCommission:  transaction.Commission,
			UpdateAt:         time.Now(),
		}
	)
	_, err := branchCollection.InsertOne(ctx, tranAnalytic)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateTranAnalytic ...
func UpdateTranAnalytic(tranAnalytic models.TranAnalyticBSON, transaction models.TransactionBSON) {
	var (
		tranAnalyticColl = database.ConnectCol("transactionAnalytic")
		ctx              = context.Background()
	)
	tranAnalytic.TotalTransaction++
	tranAnalytic.TotalRevenue += transaction.Amount
	tranAnalytic.TotalCommission += transaction.Commission
	filter := bson.M{"_id": tranAnalytic.ID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": tranAnalytic.TotalTransaction,
		"totalRevenue":     tranAnalytic.TotalRevenue,
		"totalCommission":  tranAnalytic.TotalCommission,
		"updateAt":         time.Now(),
	}}
	_, err := tranAnalyticColl.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

//BeginningOfDay ...
func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
