package util

import (
	"bytes"
	"cashbag-me-mini/config"
	"cashbag-me-mini/modules/zookeeper"
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
	// CompanyAddress for test
	CompanyAddress = "HaiPhong"
	// CompanyName for test
	CompanyName = "Hightland"
	// CompanyBalance for test
	CompanyBalance = 10000000
	// CompanyLoyalty for test
	CompanyLoyalty = 10
	// BranchID for test
	BranchID = "5f24d45125ea51bc57a8285b"
	// BranchName for test
	BranchName = "High Quang Tri"
	// BranchAddress for test
	BranchAddress = "Dong Ha"
	// UserName for test
	UserName = "Phuc"
	// UserAddress for test
	UserAddress = "48 Nguyen Chanh"
)

// HelperConnect ...
func HelperConnect() {
	zookeeper.Connect()
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
			Name:           CompanyName,
			Address:        CompanyAddress,
			Balance:        10000000,
			LoyaltyProgram: 10,
			Active:         true,
			CreatedAt:      time.Now(),
		}
	)

	// Insert
	_, err := companyCol.InsertOne(ctx, company)
	if err != nil {
		log.Println(err)
	}

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
			Name:      BranchName,
			Address:   BranchAddress,
			Active:    true,
			CreatedAt: time.Now(),
		}
	)

	// Insert
	_, err := branchCol.InsertOne(ctx, branch)
	if err != nil {
		log.Println(err)
	}
}

// HelperUserCreateFake ..
func HelperUserCreateFake() {
	var (
		userCol      = database.UserCol()
		ctx          = context.Background()
		companyID, _ = primitive.ObjectIDFromHex(CompanyID)
		user         = models.UserBSON{
			ID:        primitive.NewObjectID(),
			CompanyID: companyID,
			Name:      UserName,
			Address:   UserAddress,
		}
	)

	//Insert
	_, err := userCol.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
}

// HelperCompanyAnalyticCreateFake ...
func HelperCompanyAnalyticCreateFake() {
	var (
		companyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
		companyID, _       = primitive.ObjectIDFromHex(CompanyID)
		companyAnalytic    = models.CompanyAnalyticBSON{
			ID:        primitive.NewObjectID(),
			CompanyID: companyID,
			UpdatedAt: time.Now(),
		}
	)

	// Insert
	_, err := companyAnalyticCol.InsertOne(ctx, companyAnalytic)
	if err != nil {
		log.Println(err)
	}
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
