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
		user         = body.User
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		branchID, _  = util.ValidationObjectID(body.BranchID)
		company, _   = dao.CompanyFindByID(companyID)
		branch, _    = dao.BranchFindByID(branchID)
	)

	// Check active company & branch
	if !company.Active {
		err = errors.New("Company da dung hoat dong")
		return
	}
	if !branch.Active {
		err = errors.New("Branch da dung hoat dong")
		return
	}

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

	// Convert Transaction
	transaction = transactionCreatePayloadToBSON(body)

	// Check balance && Xử lý postpaid
	if balance < commssion {
		if !company.Postpaid {
			err = errors.New("So tien hoan tra cua cong ty da het")
			return
		}
		transaction.Postpaid = true
	}

	// Add information Transaction
	transaction.Commission = commssion
	transaction.LoyaltyProgram = company.LoyaltyProgram

	// Create Transaction
	doc, err := dao.TransactionCreate(transaction)

	// Update balance & transactionAnalytic
	if err == nil {
		if !transaction.Postpaid{
			balanceCurrent := balance - doc.Commission
			dao.CompanyUpdateBalance(doc.CompanyID, balanceCurrent)
		}
		TransactionAnalyticHandle(doc)
		errTransactionAnalyticHandle := transactionAnalyticHandleForTransaction(doc)
		if errTransactionAnalyticHandle != nil {
			return doc, errTransactionAnalyticHandle
		}
	}

	return doc, err
}


