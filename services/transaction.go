package services

import (
	"log"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/redis"
	"cashbag-me-mini/modules/zookeeper"
)

// TransactionCreate .....
func TransactionCreate(body models.TransactionCreatePayload) (models.TransactionBSON, error) {
	var (
		user        = body.User
		transaction models.TransactionBSON
	)
	// Validate User
	isUserValid := TransactionValidateUser(user)
	if !isUserValid {
		return transaction, errors.New("User khong nam trong danh sach hoan tien")
	}

	// Validate request
	userReq := redis.GetUser()
	log.Println(userReq)
	if userReq == user {
		return transaction, errors.New("User Dang Thuc hien giao dich")
	}
	redis.SetUser(body.User)

	// Validate branch id & company id
	companyID, _ := primitive.ObjectIDFromHex(body.CompanyID)
	branchID, _ := primitive.ObjectIDFromHex(body.BranchID)
	company, _ := dao.CompanyFindByID(companyID)
	branch, _ := dao.BranchFindByID(branchID)
	if company.ID.IsZero() {
		return transaction, errors.New("Khong tim thay Cong Ty ")
	}
	if branch.ID.IsZero() {
		return transaction, errors.New("Khong tim thay Chi Nhanh")
	}

	// Calculation commsion
	commssion := CalculateTransactionCommison(company.LoyaltyProgram, body.Amount)
	balance := company.Balance

	// Check balance
	if balance < commssion {
		return transaction, errors.New("So tien hoan tra cua cong ty da het")
	}

	// Convert & add information Transaction
	transaction = TransactionCreatePayloadToBSON(body)
	transaction.Commission = commssion
	transaction.LoyaltyProgram = company.LoyaltyProgram

	// Create Transaction
	doc, err := dao.TransactionCreate(transaction, balance)

	return doc, err
}

// TransactionCreatePayloadToBSON ...
func TransactionCreatePayloadToBSON(body models.TransactionCreatePayload) models.TransactionBSON {
	var (
		companyID, _ = primitive.ObjectIDFromHex(body.CompanyID)
		branchID, _  = primitive.ObjectIDFromHex(body.BranchID)
	)

	result := models.TransactionBSON{
		CompanyID: companyID,
		BranchID:  branchID,
		User:      body.User,
		Amount:    body.Amount,
	}

	return result
}

// TransactionValidateUser ...
func TransactionValidateUser(user string) bool {
	var (
		userAllowed = zookeeper.GetUser()
		users       = strings.Split(userAllowed, ",")
	)

	// Validate User
	for _, item := range users {
		if item == user {
			return true
		}
	}

	return false
}

// CalculateTransactionCommison ....
func CalculateTransactionCommison(loyatyProgram float64, amount float64) float64 {
	var(
		commission float64
	)
	commission = (loyatyProgram / 100) * amount
	
	return commission
}
//CheckValueRedis ...
func CheckValueRedis(string) bool {
	var body models.TransactionCreatePayload
	userReq := redis.GetValueRedis("user")
	if userReq == doc.User {
		return true
	} else {
		return false
	}
}
//TransactionValidUser ...
func TransactionValidUser(string) bool{
	var body models.TransactionCreatePayload
	userZoo := zookeeper.GetValueFromZoo("/Users")
	users := strings.Split(userZoo, ",")
	check := 0 
	for _, user := range users {
	if user == body.User {
		return true
	}
	}
	if check == 0 {
		return false
	}
}
// calculateTransactionCommison ....
func calculateTransactionCommison(loyatyProgram float64,amount float64) float64{
	var commission float64
	commission = (loyatyProgram/100) *amount
	return commission
}