package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"fmt"
	//"log"

//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateTransaction ...
func CreateTransaction(id models.TransactionBSON) *mongo.InsertOneResult {
	var (
		collection = database.ConnectCol("transaction")
		ctx        = context.Background()
	)
	result, err := collection.InsertOne(ctx, id)
	if err != nil {
		fmt.Println(err)
	}
	//log.Println("err")
	//updateBalance()
	//log.Println("err")
	HandleTranAnalytic(id)

	return result
}

//GetCommissionByCompanyId func ...
// func GetCommissionByCompanyId(companyId interface{}) float64 {
// 	var (
// 		transactionCollection = database.ConnectCol("transaction")
// 		ctx                   = context.Background()
// 		result                = struct {
// 			Commission float64 `json:"commission"`
// 		}{}
// 		filter = bson.M{"companyId": companyId}
// 	)
// 	err := transactionCollection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("1")
// 	return result.Commission
// }
// func updateBalance() float64 {
// 	var (
// 		companyId  models.TransactionBSON
// 		name       models.CompanyBSON
// 		commission float64
// 		balance    float64
// 	)
// 	commission = GetCommissionByCompanyId(companyId.CompanyID)
// 	log.Println(commission)
// 	balance = GetBalanceByCompanyName(name.Name)
// 	balance = balance - commission
// 	return balance
// }
