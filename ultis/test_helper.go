package ultis

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

var (
	// CompanyID for test
	CompanyID = "5f24d45125ea51bc57a8285c"
	// BranchID for test
	BranchID = "5f24d45125ea51bc57a8285b"
)

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
