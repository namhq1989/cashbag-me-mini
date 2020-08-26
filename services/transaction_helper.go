package services

import (
	"cashbag-me-mini/config"
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/redis"
	"cashbag-me-mini/util"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func transactionCreatePayloadToBSON(body models.TransactionCreatePayload, companyID, branchID, userID primitive.ObjectID) models.TransactionBSON {

	result := models.TransactionBSON{
		CompanyID: companyID,
		BranchID:  branchID,
		UserID:    userID,
		Amount:    body.Amount,
	}

	return result
}

func transactionCheckActive(companyActive bool, branchActive bool) (err error) {
	if !companyActive {
		err = errors.New("Company da dung hoat dong")
		return
	}
	if !branchActive {
		err = errors.New("Branch da dung hoat dong")
		return
	}
	return
}

func transactionCheckUserRequest(userString string) (err error) {
	userReq := redis.Get(config.RedisKeyUser)
	if userReq == userString {
		err = errors.New("User Dang Thuc hien giao dich")
		return
	}
	return
}

func calculateTransactionCommison(CompanyCashbackPercent, MilestoneCashbackPercent, amount float64) float64 {
	var (
		commission float64
	)
	commission = ((CompanyCashbackPercent + MilestoneCashbackPercent) / 100) * amount

	return commission
}

func getTransactionUserMilestoneAndExpense(companyID, userID primitive.ObjectID, amount float64) (transactionUserMilestone models.TransactionUserMilestoneAndExpense, err error) {

	// Get loyaltyProgramUserStatus
	loyaltyProgramUserStatus, err := dao.LoyaltyProgramUserStatusFindByCompanyIDAndUserID(companyID, userID)
	transactionUserMilestone.BeforeUserExpense = loyaltyProgramUserStatus.CurrentExpense
	transactionUserMilestone.CurrentUserExpense = loyaltyProgramUserStatus.CurrentExpense + amount
	transactionUserMilestone.BeforeUserMilestone = loyaltyProgramUserStatus.Milestone

	// Get loyaltyProgram
	loyaltyProgram, err := dao.LoyaltyProgramFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong tim thay chuong trinh tich diem cua cong ty")
		return
	}
	milestones := loyaltyProgram.Milestones

	// Reverse milestones
	milestones = util.HelperReverseArrayMilestone(milestones)

	// Check currentUserMilestone
	for _, milestone := range milestones {
		if transactionUserMilestone.CurrentUserExpense >= milestone.Expense {
			transactionUserMilestone.CurrentUserMilestone = milestone
			break
		}
	}

	return
}

func transactionAddInformation(transaction models.TransactionBSON, commission, companyCashbackPercent, milestoneCashbackPercent float64, paidType string) models.TransactionBSON {
	transaction.Commission = commission
	transaction.CompanyCashbackPercent = companyCashbackPercent
	transaction.MilestoneCashbackPercent = milestoneCashbackPercent
	transaction.PaidType = paidType
	transaction.ID = primitive.NewObjectID()
	transaction.CreatedAt = time.Now()
	return transaction
}

func createPrepaidTransaction(transaction models.TransactionBSON, balance float64) (err error) {
	if balance < transaction.Commission {
		err = errors.New("So tien hoan tra cua cong ty da het")
		return
	}

	// Create Transaction
	_, err = dao.TransactionCreate(transaction)
	if err != nil {
		return
	}

	// Update Balance
	balanceCurrent := balance - transaction.Commission
	err = companyUpdateBalance(transaction.CompanyID, balanceCurrent)
	if err != nil {
		err = errors.New("Update Balance That Bai")
	}
	return
}

func createPostpaidTransaction(transaction models.TransactionBSON) (err error) {

	// Create Transaction
	_, err = dao.TransactionCreate(transaction)
	return
}

func loyaltyProgramUserStatusUpdateAfterCreateTransaction(transactionUserMilestone models.TransactionUserMilestoneAndExpense, companyID, userID primitive.ObjectID) (err error) {
	var (
		beforeUserMilestone  = transactionUserMilestone.BeforeUserMilestone
		currentUserMilestone = transactionUserMilestone.CurrentUserMilestone
		beforeUserExpense    = transactionUserMilestone.BeforeUserExpense
		currentUserExpense   = transactionUserMilestone.CurrentUserExpense
	)

	// No upgrade Milestone for case nil Milestone
	if beforeUserExpense == 0 {
		err = createLoyaltyProgramUserStatus(currentUserMilestone, currentUserExpense, companyID, userID)
		return
	}

	// No upgrade Milestone
	if currentUserMilestone.ID == beforeUserMilestone.ID {
		err = updateLoyaltyProgramUserStatus(beforeUserMilestone, currentUserExpense, userID)
		return
	}

	// Upgrade Milestone
	err = createLoyaltyProgramUserStatus(currentUserMilestone, currentUserExpense, companyID, userID)
	return
}

func companyAnalyticUpdateAfterCreateTransaction(transaction models.TransactionBSON, currentUserMilestone, beforeUserMilestone models.LoyaltyProgramMilestone) (err error) {
	var (
		postpaid      = "postpaid"
		companyID     = transaction.CompanyID
		countPostpaid int
	)

	// Find CompanyAnalytic
	companyAnalytic, err := dao.CompanyAnalyticFindByCompanyID(companyID)
	if err != nil {
		err = errors.New("Khong Tim Thay CompanyAnalytic")
		return
	}

	// Set data Update CompanyAnalytic
	companyAnalytic.TotalTransaction++
	companyAnalytic.TotalRevenue += transaction.Amount
	companyAnalytic.TotalCommission += transaction.Commission
	if transaction.PaidType == postpaid {
		companyAnalytic.TotalDebt += transaction.Commission
		companyAnalytic.CountPostpaid++
		countPostpaid = companyAnalytic.CountPostpaid
	}

	// PostpaidLog
	err = postpaidLog(countPostpaid, companyID)
	if err != nil {
		err = errors.New("Deactive company trong postpaidLog that bai")
		return
	}

	// Update CompanyAnalytic case no upgrade milestone
	if currentUserMilestone.ID == beforeUserMilestone.ID {
		err = companyAnalyticUpdateNoUpgradeMilestone(companyAnalytic)
	} else {

		//  Update CompanyAnalytic case upgrade milestone
		err = companyAnalyticUpdateUpgradeMilestone(companyAnalytic, currentUserMilestone, beforeUserMilestone)
	}

	if err != nil {
		err = errors.New("Update CompanyAnalytic That Bai")
		return
	}
	return
}

func companyAnalyticUpdateNoUpgradeMilestone(companyAnalytic models.CompanyAnalyticBSON) (err error) {
	var (
		filter = bson.M{"_id": companyAnalytic.ID}
		update = bson.M{"$set": bson.M{
			"totalTransaction": companyAnalytic.TotalTransaction,
			"totalRevenue":     companyAnalytic.TotalRevenue,
			"totalCommission":  companyAnalytic.TotalCommission,
			"totalDebt":        companyAnalytic.TotalDebt,
			"countPostpaid":    companyAnalytic.CountPostpaid,
			"updatedAt":        time.Now(),
		}}
	)

	// Update
	err = dao.CompanyAnalyticUpdateByID(filter, update)

	return
}

func companyAnalyticUpdateUpgradeMilestone(companyAnalytic models.CompanyAnalyticBSON, currentUserMilestone, beforeUserMilestone models.LoyaltyProgramMilestone) (err error) {
	var (
		members []models.CompanyAnalyticMember
	)

	for _, member := range companyAnalytic.Members {
		if member.ID == currentUserMilestone.ID {
			member.Total++
		}
		if member.ID == beforeUserMilestone.ID {
			member.Total--
		}
		members = append(members, member)
	}

	filter := bson.M{"_id": companyAnalytic.ID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": companyAnalytic.TotalTransaction,
		"totalRevenue":     companyAnalytic.TotalRevenue,
		"totalCommission":  companyAnalytic.TotalCommission,
		"totalDebt":        companyAnalytic.TotalDebt,
		"countPostpaid":    companyAnalytic.CountPostpaid,
		"members":          members,
		"updatedAt":        time.Now(),
	}}

	// Update
	err = dao.CompanyAnalyticUpdateByID(filter, update)

	return
}

func postpaidLog(countPostpaid int, companyID primitive.ObjectID) (err error) {
	if countPostpaid > 3 {
		log.Println("Số đơn hàng trả sau vượt mức cho phép")
	}
	if countPostpaid > 6 {
		log.Println("Số đơn hàng trả sau vượt mức cho phép và deactive company")
		_, err = CompanyChangeActiveStatus(companyID, false)
		return
	}
	return
}

func convertToTransactionDetail(transaction models.TransactionBSON) models.TransactionDetail {
	var (
		company, _  = dao.CompanyFindByID(transaction.CompanyID)
		branch, _   = dao.BranchFindByID(transaction.BranchID)
		user, _     = dao.UserFindByID(transaction.UserID)
		companyName = company.Name
		branchName  = branch.Name
		userName    = user.Name
	)

	// TransactionDetail
	result := models.TransactionDetail{
		ID:                       transaction.ID,
		Company:                  companyName,
		Branch:                   branchName,
		User:                     userName,
		Amount:                   transaction.Amount,
		Commission:               transaction.Commission,
		CompanyCashbackPercent:   transaction.CompanyCashbackPercent,
		MilestoneCashbackPercent: transaction.MilestoneCashbackPercent,
		PaidType:                 transaction.PaidType,
		CreatedAt:                transaction.CreatedAt,
	}

	return result
}
