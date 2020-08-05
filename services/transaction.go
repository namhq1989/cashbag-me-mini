package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateTransaction func ....
func CreateTransaction(body models.PostTransaction) *mongo.InsertOneResult {
	var (
		result         *mongo.InsertOneResult
	)
	branchID := dao.GetIdBranchByName(body.NameBranch)
	ifCompany :=dao.GetIFCompanyByName(body.NameCompany)
	transaction := ConvertBodyToTransactionBSON(body)
	commission := (ifCompany.LoyaltyProgram / 100) * body.Amount
	transaction.LoyaltyProgram = ifCompany.LoyaltyProgram
	transaction.Commission = commission
	transaction.CompanyID = ifCompany.ID
	transaction.BranchID = branchID
	transaction.ID = primitive.NewObjectID()
	transaction.CreateAt = time.Now()
	if ifCompany.Balance >= commission {
		result = dao.CreateTransaction(transaction,ifCompany.Balance)
	} else {
		log.Fatal("So tien hoan tra cua cong ty da het")
	}
	return result
}


//ConvertBodyToTransactionBSON func ...
func ConvertBodyToTransactionBSON(body models.PostTransaction) models.TransactionBSON {
	result := models.TransactionBSON{
		User:   body.User,
		Amount: body.Amount,
	}
	return result
}

