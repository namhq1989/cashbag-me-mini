package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateTransaction func ....
func CreateTransaction(body models.PostTransaction) *mongo.InsertOneResult {
	var (
		transaction    models.TransactionBSON
		companyID      primitive.ObjectID
		branchID       primitive.ObjectID
		commission     float64
		loyaltyProgram float64
		amount         float64
	)

	companyID = dao.GetIdCompanyByName(body.NameCompany)
	branchID = dao.GetIdBranchByName(body.NameBranch)
	loyaltyProgram = dao.GetLoyaltyProgramByCompany(body.NameCompany)
	transaction = ConvertBodyToTransactionBSON(body)
	amount = body.Amount
	commission = (loyaltyProgram / 100) * amount
	transaction.LoyaltyProgram = loyaltyProgram
	transaction.Commission = commission
	transaction.CompanyID = companyID
	transaction.BranchID = branchID
	transaction.ID = primitive.NewObjectID()
	transaction.CreateAt = time.Now()
	result := dao.CreateTransaction(transaction)
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

// //CalCulateCommission func ...
// func CalCulateCommission(body models.PostTransaction) float64{
// 	var ( transaction    models.TransactionBSON
// 		loyaltyProgram float64
// 		)
// 	CreateTransaction(*body)
// 	loyaltyProgram = dao.GetLoyaltyProgramByCompany(body.NameCompany)
// 	commission = loyaltyProgram * (transaction.Amount)
// 	transaction.Commission = commission

// }
