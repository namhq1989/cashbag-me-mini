package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateTransaction func ....
func CreateTransaction(body models.PostTransaction) *mongo.InsertOneResult {
	var (
		transaction models.TransactionBSON
		companyID   primitive.ObjectID
		branchID    primitive.ObjectID
	)
	companyID = dao.GetIdCompanyByName(body.NameCompany)
	branchID = GetIdBranchByName(body.NameBranch)
	transaction = ConvertBodyToTransactionBSON(body)
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


//GetIdBranchByName .....
func GetIdBranchByName(NameBranch interface{}) primitive.ObjectID {

	var (
		branchCollection = database.ConnectCol("branches")
		ctx              = context.Background()
		result           = struct {
			ID primitive.ObjectID `bson:"_id"`
		}{}
		filter = bson.M{"name": NameBranch}
	)
	err := branchCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.ID
}
