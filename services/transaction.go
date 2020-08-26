package services

import (
	"cashbag-me-mini/config"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/redis"
)

// TransactionCreate ...
func TransactionCreate(body models.TransactionCreatePayload, company models.CompanyBSON, branch models.BranchBSON, user models.UserBSON) (transaction models.TransactionBSON, err error) {
	var (
		prepaid    = "prepaid"
		userString = body.UserID
		companyID  = company.ID
		branchID   = branch.ID
		userID     = user.ID
		balance    = company.Balance
	)

	// Check active company & branch
	err = transactionCheckActive(company.Active, branch.Active)
	if err != nil {
		return
	}

	// Check User Request
	err = transactionCheckUserRequest(userString)
	if err != nil {
		return
	}
	redis.Set(config.RedisKeyUser, userString)

	// Get TransactionUserMilestoneAndExpense
	transactionUserMilestone, err := getTransactionUserMilestoneAndExpense(companyID, userID, body.Amount)
	if err != nil {
		return
	}
	currentUserMilestone := transactionUserMilestone.CurrentUserMilestone
	beforeUserMilestone := transactionUserMilestone.BeforeUserMilestone

	// Calculate commission
	commission := calculateTransactionCommison(company.CashbackPercent, currentUserMilestone.CashbackPercent, body.Amount)

	// Convert Transaction
	transaction = transactionCreatePayloadToBSON(body, companyID, branchID, userID)

	// Add information for Transaction
	transaction = transactionAddInformation(transaction, commission, company.CashbackPercent, currentUserMilestone.CashbackPercent, company.PaidType)

	if company.PaidType == prepaid {
		err = createPrepaidTransaction(transaction, balance)
	} else {
		err = createPostpaidTransaction(transaction)
	}
	if err != nil {
		return
	}

	// Update TransactionAnalytic
	err = transactionAnalyticUpdateAfterCreateTransaction(transaction)
	if err != nil {
		return
	}

	// Update LoyaltyProgramUserStatus
	err = loyaltyProgramUserStatusUpdateAfterCreateTransaction(transactionUserMilestone, companyID, userID)
	if err != nil {
		return
	}

	// Update CompanyAnalytic
	err = companyAnalyticUpdateAfterCreateTransaction(transaction, currentUserMilestone, beforeUserMilestone)
	if err != nil {
		return
	}

	return
}
