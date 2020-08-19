package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
func calculateTransactionCommison(loyatyProgram float64, userProgram float64, amount float64) float64 {
	var (
		commission float64
	)
	commission = ((loyatyProgram + userProgram) / 100) * amount

	return commission
}

func getUserInformation(companyID primitive.ObjectID, userID primitive.ObjectID, amount float64) (userInformation models.UserInformation, err error) {
	var(
		silverLevel = "Silver"
		goldenLevel ="Golden"
		diamondLevel ="Diamond"
	)

	// Find UserProgram
	userProgramDoc, err := dao.UserProgramFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong tim thay chuong trinh tich diem cua cong ty")
		return
	}

	// Find User
	userDoc, err := dao.UserFindByID(userID)
	if err != nil {
		err = errors.New("Khong tim thay nguoi dung")
		return
	}

	// Get silver, golden, diamond
	silver              := userProgramDoc.Silver
	golden              := userProgramDoc.Golden
	diamond             := userProgramDoc.Diamond
	beforeUserSpending  := userDoc.Spending
	currentUserSpending := beforeUserSpending + amount

	// Set userInformation
	userInformation.CurrentUserSpending = currentUserSpending
	userInformation.BeforeUserLevel=userDoc.Level

	// userProgram level
	if currentUserSpending <= silver.Spending {
		userInformation.UserProgram = 0
	}

	if currentUserSpending >= silver.Spending && currentUserSpending < golden.Spending {
		userInformation.UserProgram = silver.Commission
		userInformation.CurrentUserLevel= silverLevel
	}

	if currentUserSpending >= golden.Spending && currentUserSpending < diamond.Spending {
		userInformation.UserProgram = golden.Commission
		userInformation.CurrentUserLevel= goldenLevel
	}

	if currentUserSpending >= diamond.Spending {
		userInformation.UserProgram = diamond.Commission
		userInformation.CurrentUserLevel= diamondLevel
	}
	
	return
}
