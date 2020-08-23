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
		balance      = company.Balance
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

	// Get userInformation
	userInformation, err := getUserInformation(companyID, userID, body.Amount)
	if err != nil {
		return
	}
	userProgram := userInformation.UserProgram
	currentUserSpending := userInformation.CurrentUserSpending
	currentUserLevel := userInformation.CurrentUserLevel
	beforeUserLevel := userInformation.BeforeUserLevel

	// Calculate commission
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
	transaction, err = dao.TransactionCreate(transaction)

	// Update balance & transactionAnalytic
	if err == nil {
		if !transaction.Postpaid {
			balanceCurrent := balance - transaction.Commission
			CompanyUpdateBalance(transaction.CompanyID, balanceCurrent)
		}

		// Handle TransactionAnalytic
		TransactionAnalyticHandle(transaction)

		// Update spending && level for User
		err = UserUpdateSpendingAndLevel(transaction.UserID, currentUserLevel, currentUserSpending)
		if err != nil {
			return
		}

		// Update CompanyAnalytic
		err = companyAnalyticHandleForTransaction(transaction, beforeUserLevel, currentUserLevel)
		if err != nil {
			return
		}
	}

	return
}
