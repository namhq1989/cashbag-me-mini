package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

func transactionAnalyticUpdateAfterCreateTransaction(transaction models.TransactionBSON) (err error) {
	tranAnalytic, err := dao.TransactionAnalyticFilterByDate(transaction)
	if err != nil {
		err = transactionAnalyticCreate(transaction)
		return err
	}
	err = transactionAnalyticUpdate(tranAnalytic, transaction)
	return err
}

func transactionAnalyticCreate(transaction models.TransactionBSON) error {
	var (
		startOfDate = util.BeginningOfDay(transaction.CreatedAt)
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

	// Create transactionAnalytic
	err := dao.TransactionAnalyticCreate(transactionAnalytic)

	return err
}

func transactionAnalyticUpdate(transactionAnalytic models.TransactionAnalyticBSON, transaction models.TransactionBSON) error {

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
	err := dao.TransactionAnalyticUpdateByID(filter, update)
	return err
}

func convertToTransactionAnalyticDetail(doc models.TransactionAnalyticBSON) models.TransactionAnalyticDetail {
	var (
		company, _   = dao.CompanyFindByID(doc.CompanyID)
		nameCompany  = company.Name
		companyBrief models.CompanyBrief
		branch, _    = dao.BranchFindByID(doc.BranchID)
		nameBranch   = branch.Name
		brannchBrief models.BranchBrief
	)

	// Add information companyBrief
	companyBrief.ID = doc.CompanyID
	companyBrief.Name = nameCompany

	// Add information branchBrief
	brannchBrief.ID = doc.BranchID
	brannchBrief.Name = nameBranch

	// Transaction Analytic Detail
	result := models.TransactionAnalyticDetail{
		ID:               doc.ID,
		Company:          companyBrief,
		Branch:           brannchBrief,
		Date:             doc.Date,
		TotalTransaction: doc.TotalTransaction,
		TotalRevenue:     doc.TotalRevenue,
		TotalCommission:  doc.TotalCommission,
		UpdateAt:         doc.UpdateAt,
	}
	return result
}
