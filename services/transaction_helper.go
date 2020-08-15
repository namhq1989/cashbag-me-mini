package services

import (
	"strings"

	"cashbag-me-mini/config"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// transactionCreatePayloadToBSON ...
func transactionCreatePayloadToBSON(body models.TransactionCreatePayload) models.TransactionBSON {
	var (
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		branchID, _  = util.ValidationObjectID(body.BranchID)
	)

	result := models.TransactionBSON{
		CompanyID: companyID,
		BranchID:  branchID,
		User:      body.User,
		Amount:    body.Amount,
	}

	return result
}

// transactionValidateUser ...
func transactionValidateUser(user string) bool {
	var (
		envVars     = config.GetEnv()
		userAllowed = envVars.UserString
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

// calculateTransactionCommison ....
func calculateTransactionCommison(loyatyProgram float64, amount float64) float64 {
	var (
		commission float64
	)
	commission = (loyatyProgram / 100) * amount

	return commission
}
