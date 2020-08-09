package services

import (
	
	"log"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/redis"
)

//TransactionCreate .....
func TransactionCreate(body models.TransactionCreatePayload) (models.TransactionBSON, error) {
	company :=dao.CompanyFindbyID(body.CompanyID)
	if company.ID.isZero() 	{
		return errors.New("Khong tim thay cong ty")
	}
	branch := dao.BranchFindbyID(body.BranchID)
	if branch.ID.isZero(){
		//return error.New("khong tim thay Cua hang ")
	}
	value := CheckValueRedis(body.User)
		if value == true {
			//return error.New("User dang thuc hien giao dich")
		} 
	//check uservalid
	isValidUser :=TransactionValidUser(body.User)
	if (!isValidUser) {
		//return error.New("User khong nam trong danh sach hoan tien ")
	}
	redis.SetValueRedis("user", body.User)	
	//calculation commsion 
	commssion := calculateTransactionCommison(company.LoyatyProgram,body.Amount)
	balance :=company.Balance
	//balance vs commission
	if balance < commssion {
		return error.New("so tien hoan tra cua cty da het")
	}
	doc,err = dao.TransactionCreate(body,balance)
	return doc,err
	
}

//ConvertBodyToTransactionBSON func ...
func ConvertBodyToTransactionBSON(body models.TransactionCreatePayload) models.TransactionBSON {
	result := models.TransactionBSON{
		User:   body.User,
		Amount: body.Amount,
	}
	return result
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