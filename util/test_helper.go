package util

import (
	"bytes"
	"cashbag-me-mini/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

var (
	// CompanyID for test
	CompanyID = "5f24d45125ea51bc57a8285c"
	// BranchID for test
	BranchID = "5f24d45125ea51bc57a8285b"
)

// HelperConnect ...
func HelperConnect() {
	envVars := config.GetEnv()

	// Connect
	client, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", envVars.Database.URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	db := client.Database(envVars.Database.Name)
	fmt.Println("Database Connected to", envVars.Database.TestName)
	database.SetDB(db)

}

// HelperToIOReader ...
func HelperToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// HelperCompanyCreateFake ...
func HelperCompanyCreateFake() {
	var (
		companyCol   = database.CompanyCol()
		ctx          = context.Background()
		companyID, _ = primitive.ObjectIDFromHex(CompanyID)
		company      = models.CompanyBSON{
			ID:             companyID,
			Name:           "Hightland",
			Address:        "HaiPhong",
			Balance:        10000000,
			LoyaltyProgram: 10,
			Active:         false,
			CreatedAt:      time.Now(),
		}
	)

	// Insert
	companyCol.InsertOne(ctx, company)
}

// HelperBranchCreateFake ...
func HelperBranchCreateFake() {
	var (
		branchCol    = database.BranchCol()
		ctx          = context.Background()
		companyID, _ = primitive.ObjectIDFromHex(CompanyID)
		branchID, _  = primitive.ObjectIDFromHex(BranchID)
		branch       = models.BranchBSON{
			ID:        branchID,
			CompanyID: companyID,
			Name:      "Hight QuangTri",
			Address:   "120 QuangTri",
			Active:    false,
			CreatedAt: time.Now(),
		}
	)

	// Insert
	branchCol.InsertOne(ctx, branch)
}

// HelperTransactionAnalyticFindByID ...
func HelperTransactionAnalyticFindByID() models.TransactionAnalyticBSON {
	var (
		transactionAnalyticCol = database.TransactionAnalyticCol()
		ctx                    = context.Background()
		branchID, _            = primitive.ObjectIDFromHex(BranchID)
		companyID, _           = primitive.ObjectIDFromHex(CompanyID)
		filter                 = bson.M{
			"companyID": companyID,
			"branchID":  branchID,
		}
		result models.TransactionAnalyticBSON
	)

	// Find
	transactionAnalyticCol.FindOne(ctx, filter).Decode(&result)

	return result
}
