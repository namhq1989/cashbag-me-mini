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
	// CompanyString for test
	CompanyString = "5f24d45125ea51bc57a8285c"

	// CompanyID for test
	CompanyID, _ = primitive.ObjectIDFromHex(CompanyString)

	// Company for test
	Company    = models.CompanyBSON{
		ID:              CompanyID,
		Name:            "Hightland",
		Address:         "HaiPhong",
		Balance:         10000000,
		CashbackPercent: 10,
		Active:          true,
		PaidType:        "prepaid",
		CreatedAt:       time.Now(),
	}

	// BranchString for test
	BranchString = "5f24d45125ea51bc57a8285b"

	// BranchID for test
	BranchID, _ = primitive.ObjectIDFromHex(BranchString)

	// Branch for test
	Branch    = models.BranchBSON{
		ID:        BranchID,
		CompanyID: CompanyID,
		Name:      "High Quang Tri",
		Address:   "Dong Ha",
		Active:    true,
		CreatedAt: time.Now(),
	}

	// UserString for test
	UserString = "5f24d45125ea51bc57a8285a"

	// UserID for test
	UserID, _ = primitive.ObjectIDFromHex(UserString)

	// User for test
	User    = models.UserBSON{
		ID:      UserID,
		Name:    "Phuc",
		Address: "48 Nguyen Chanh",
	}
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
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	// Insert
	_, err := companyCol.InsertOne(ctx, Company)
	if err != nil {
		log.Println(err)
	}

}

// HelperBranchCreateFake ...
func HelperBranchCreateFake() {
	var (
		branchCol = database.BranchCol()
		ctx       = context.Background()
	)

	// Insert
	_, err := branchCol.InsertOne(ctx, Branch)
	if err != nil {
		log.Println(err)
	}
}

// HelperUserCreateFake ..
func HelperUserCreateFake() {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	//Insert
	_, err := userCol.InsertOne(ctx, User)
	if err != nil {
		log.Println(err)
	}
}

// HelperCompanyAnalyticCreateFake ...
func HelperCompanyAnalyticCreateFake() {
	var (
		companyAnalyticCol = database.CompanyAnalyticCol()
		ctx                = context.Background()
		companyAnalytic    = models.CompanyAnalyticBSON{
			ID:        primitive.NewObjectID(),
			CompanyID: CompanyID,
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
		filter                 = bson.M{
			"companyID": CompanyID,
			"branchID":  BranchID,
		}
		result models.TransactionAnalyticBSON
	)

	// Find
	transactionAnalyticCol.FindOne(ctx, filter).Decode(&result)

	return result
}
