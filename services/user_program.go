package services

import(
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)
// UserProgramCreate ...
func UserProgramCreate(body models.UserProgramCreatePayload) ( models.UserProgramBSON, error){
	var (
		userProgram models.UserProgramBSON
		companyID, _ = primitive.ObjectIDFromHex(body.CompanyID)
		company, _   = dao.CompanyFindByID(companyID)
	)

	// Validate CompanyID
	if company.ID.IsZero() {
		return userProgram, errors.New("Khong tim thay Cong Ty")
	}

	// validate Silver
	isSilver :=silverValidate()
	if !isSilver {
		return userProgram,errors.New("Muc bac khong hop ly")
	}

	// validate Golden 
	isGolden :=goldenValidate()
	if !isGolden {
		return userProgram,errors.New("Muc vang khong hop ly")
	}

	// validate diamond
	isDiamond :=diamondValidate()
	if !isDiamond {
		return userProgram,errors.New("Muc kim cuong khong hop ly")
	}
	

	//create userProgram 
	userProgram =userCreatePayloadToBSON(body)
	doc,err :=dao.UserProgramCreate(userProgram)

	return doc,err
}


