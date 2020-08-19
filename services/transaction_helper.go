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

func calculateCurrentUserProgram(companyID primitive.ObjectID, userID primitive.ObjectID, amount float64) (userProgram float64, currentUserSpending float64, level string, err error) {
	// Lay cac muc bac ,kimcuong, vang
	button, err := dao.UserProgramFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong tim thay chuong trinh tich diem cua cong ty")
		return
	}

	// Lay spending
	userFind, err := dao.UserFindByID(userID)
	if err != nil {
		err = errors.New("Khong tim thay nguoi dung")
		return
	}

	var (
		silver              = button.Silver
		golden              = button.Golden
		diamond             = button.Diamond
		beforeUserSpending  = userFind.Spending
	)
	currentUserSpending = beforeUserSpending + amount

	// userProgram level
	if currentUserSpending <= silver.Spending {
		userProgram = 0
	}

	if currentUserSpending >= silver.Spending && currentUserSpending < golden.Spending {
		userProgram = silver.Commission
	}

	if currentUserSpending >= golden.Spending && currentUserSpending < diamond.Spending {
		userProgram = golden.Commission
	}

	if currentUserSpending >= diamond.Spending {
		userProgram = diamond.Commission
	}

	// so sanh spending voi cac muc cua userProgram
	if beforeUserSpending >= silver.Spending && beforeUserSpending < golden.Spending {
		level = "Muc bac"
	}

	if beforeUserSpending >= golden.Spending && beforeUserSpending < diamond.Spending {
		level = "Muc vang"
	}

	if beforeUserSpending >= diamond.Spending {
		level = "Muc kim cuong"
	}
	return
}
