package services

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// transactionCreatePayloadToBSON ...
func transactionCreatePayloadToBSON(body models.TransactionCreatePayload) models.TransactionBSON {
	var (
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		branchID, _  = util.ValidationObjectID(body.BranchID)
		userID, _    = util.ValidationObjectID(body.UserID)
	)

	result := models.TransactionBSON{
		CompanyID: companyID,
		BranchID:  branchID,
		UserID:    userID,
		Amount:    body.Amount,
	}

	return result
}

// calculateTransactionCommison ....
func calculateTransactionCommison(loyatyProgram float64, amount float64) float64 {
	var (
		commission float64
	)
	commission = (loyatyProgram / 100) * amount

	return commission
}
