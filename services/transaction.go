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
		userProgram  float64
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

	// Lay cac muc bac ,kimcuong, vang
	button, err := dao.UserProgramFindByID(companyID)
	if err != nil {
		err = errors.New("Khong tim thay Chuong Trinh Tich Diem cua Cong Ty")
		return
	}

	// Lay spending
	userFind, err := dao.UserFindByID(userID)
	if err != nil {
		err = errors.New("Khong tim thay user ")
		return
	}

	var (
		silver         = button.Silver
		golden         = button.Golden
		diamond        = button.Diamond
		beforeSpending = userFind.Spending
		calLevel       = beforeSpending + body.Amount
		balance        = company.Balance
	)

	// userProgram level
	if calLevel <= silver.Spending {
		userProgram = 0
	}

	if calLevel >= silver.Spending && calLevel < golden.Spending {
		userProgram = silver.Commission
	}

	if calLevel >= golden.Spending && calLevel < diamond.Spending {
		userProgram = golden.Commission
	}

	if calLevel >= diamond.Spending {
		userProgram = diamond.Commission
	}

	// Calculation commsion
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

		// Update spending
		afterSpending := doc.Amount + beforeSpending
		err = dao.UserUpdateSpending(doc.UserID, afterSpending)
		if err != nil {
			return doc, err
		}

		// Update level
		afterLevel := checkUserLevelByID(doc.CompanyID, doc.UserID)
		err = dao.UserUpdateLevel(doc.UserID, afterLevel)
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
