package services

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// UserCreatePayloadToBSON ...
func userCreatePayloadToBSON(body models.UserCreatePayload) models.UserBSON {
	var (
		companyID, _ = util.ValidationObjectID(body.CompanyID)
	)
	result := models.UserBSON{
		CompanyID: companyID,
		Name:      body.Name,
		Address:   body.Address,
	}
	return result
}

// UserUpdatePayloadToBSON ...
func userUpdatePayloadToBSON(body models.UserUpdatePayload) models.UserBSON {
	result := models.UserBSON{
		Name:    body.Name,
		Address: body.Address,
	}
	return result
}

func checkUserLevelByID(companyID primitive.ObjectID, userID primitive.ObjectID) (level string) {
	button, err := dao.UserProgramFindByID(companyID)
	if err != nil {
		log.Println(err)
	}

	// Lay spending
	userFind, err := dao.UserFindByID(userID)
	if err != nil {
		log.Println(err)
	}

	var (
		silver  = button.Silver
		golden  = button.Golden
		diamond = button.Diamond
		spending = userFind.Spending
	)
	
	// so sanh spending voi cac muc cua userProgram
	if spending >= silver.Spending && spending < golden.Spending {
		level = "Muc bac"
	}

	if spending >= golden.Spending && spending < diamond.Spending {
		level = "Muc vang"
	}

	if spending >= diamond.Spending {
		level = "Muc kim cuong"
	}
	
	return
}
