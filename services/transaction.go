package services

import (
	"errors"

	"cashbag-me-mini/config"
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/redis"
	"cashbag-me-mini/util"
)

// TransactionCreate ...
func TransactionCreate(body models.TransactionCreatePayload) (transaction models.TransactionBSON, err error) {
	var (
		user         = body.UserID
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		branchID, _  = util.ValidationObjectID(body.BranchID)
		userID, _    = util.ValidationObjectID(body.UserID)
		company, _   = dao.CompanyFindByID(companyID)
		branch, _    = dao.BranchFindByID(branchID)
		balance        = company.Balance
	
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

	// Tinh spending hien tai
	userProgram,currentUserSpending,level,err :=calculateCurrentUserProgram(companyID,userID,body.Amount)
	if err !=nil{
		err =errors.New("Khong the tinh duoc muc userProgram hien tai")
		return 
	}
	commission := calculateTransactionCommison(company.LoyaltyProgram, userProgram, body.Amount)
	
	// Convert Transaction
	transaction = transactionCreatePayloadToBSON(body)

	// Check balance && Xử lý postpaid
	if balance < commission {
		if !company.Postpaid {
			err = errors.New("So tien hoan tra cua cong ty da het")
			return
		}
		transaction.Postpaid = true
	}

	// Add information Transaction
	transaction.Commission = commission
	transaction.LoyaltyProgram = company.LoyaltyProgram
	transaction.UserProgram = userProgram

	// Create Transaction
	doc, err := dao.TransactionCreate(transaction)

	// Update balance & transactionAnalytic
	if err == nil {
		if !transaction.Postpaid {
			balanceCurrent := balance - doc.Commission
			dao.CompanyUpdateBalance(doc.CompanyID, balanceCurrent)
		}

		// Update spending && level
		err = dao.UserUpdateSpendingAndLevel(doc.UserID,level,currentUserSpending)
		if err != nil {
			return doc, err
		}

		TransactionAnalyticHandle(doc)
		errTransactionAnalyticHandle := transactionAnalyticHandleForTransaction(doc)
		if errTransactionAnalyticHandle != nil {
			return doc, errTransactionAnalyticHandle
		}
	}

	return doc, err
}
