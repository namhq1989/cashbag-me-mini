package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"errors"
)

// UserProgramCreate ...
func UserProgramCreate(body models.UserProgramCreatePayload) (userProgram models.UserProgramBSON, err error) {

	var (
		silverButton  = silverProgramToSilverButton(body)
		goldenButton  = goldenProgramToGoldenButton(body)
		diamondButton = diamondProgramToDiamondButton(body)
	)
	silver := silverValidate(silverButton)
	if !silver {
		err = errors.New("Muc bac khong hop li")
		return
	}

	golden := goldenValidate(silverButton, goldenButton)
	if !golden {
		err = errors.New("Muc vang khong hop li")
		return
	}

	diamond := diamondValidate(goldenButton, diamondButton)
	if !diamond {
		err = errors.New("Muc kim cuong khong hop li")
		return
	}

	//create userProgram

	userProgram = userProgramCreatePayloadToBSON(body)
	doc, err := dao.UserProgramCreate(userProgram)
	if diamond {
		doc.Active = true
	} else {
		doc.Active = false
	}
	if err == nil {
		dao.CompanyUpdateActive(doc.CompanyID, doc.Active)
	}

	return doc, err
}
