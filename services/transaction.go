package services

import (
	"cashbag-me-mini/util"
	"errors"

	"cashbag-me-mini/config"
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/redis"
)

// TransactionCreate ...
func TransactionCreate(body models.TransactionCreatePayload) (transaction models.TransactionBSON, err error) {
	var (
		user         = body.UserID
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		company, _   = dao.CompanyFindByID(companyID)
	)

	// Validate request
	userReq := redis.Get(config.RedisKeyUser)
	if userReq == user {
		err = errors.New("User Dang Thuc hien giao dich")
		return
	}
	redis.Set(config.RedisKeyUser, body.UserID)

	// Calculation commsion
	commssion := calculateTransactionCommison(company.LoyaltyProgram, body.Amount)
	balance := company.Balance

	// Check balance
	if balance < commssion {
		err = errors.New("So tien hoan tra cua cong ty da het")
		return
	}

	// Convert & add information Transaction
	transaction = transactionCreatePayloadToBSON(body)
	transaction.Commission = commssion
	transaction.LoyaltyProgram = company.LoyaltyProgram

	// Create Transaction
	doc, err := dao.TransactionCreate(transaction)

	// Update balance & transactionAnalytic
	if err == nil {
		balanceCurrent := balance - doc.Commission
		dao.CompanyUpdateBalance(doc.CompanyID, balanceCurrent)
		TransactionAnalyticHandle(doc)
		errTransactionAnalyticHandle := transactionAnalyticHandleForTransaction(doc)
		if errTransactionAnalyticHandle != nil {
			return doc, errTransactionAnalyticHandle
		}
	}

	return doc, err
}


