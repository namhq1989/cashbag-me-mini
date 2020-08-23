package services

import (
	"errors"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

// UserProgramCreate ...
func UserProgramCreate(body models.UserProgramCreatePayload) (userProgram models.UserProgramBSON, err error) {
	var (
		silverButton  = silverProgramToSilverButton(body)
		goldenButton  = goldenProgramToGoldenButton(body)
		diamondButton = diamondProgramToDiamondButton(body)	
	)

	// validate silver
	silver := silverValidate(silverButton)
	if !silver {
		err = errors.New("Muc bac khong hop li")
		return
	}

	// validate golden
	golden := goldenValidate(silverButton, goldenButton)
	if !golden {
		err = errors.New("Muc vang khong hop li")
		return
	}

	// validate diamond
	diamond := diamondValidate(goldenButton, diamondButton)
	if !diamond {
		err = errors.New("Muc kim cuong khong hop li")
		return
	}
	
	//create userProgram

	userProgram = userProgramCreatePayloadToBSON(body)
	doc, err := dao.UserProgramCreate(userProgram)
	
	if err == nil {
		errCompanyActive := CompanyUpdateActiveTrue(doc.CompanyID)
		if errCompanyActive != nil{
			return doc,err
		}
	}

	return doc, err
}

