package services

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/util"
)

// UserCreatePayloadToBSON ...
func userProgramCreatePayloadToBSON(body models.UserProgramCreatePayload) models.UserProgramBSON {
	var (
		companyID, _ = util.ValidationObjectID(body.CompanyID)
		silverButton  = silverProgramToSilverButton(body)
		goldenButton  = goldenProgramToGoldenButton(body)
		diamondButton = diamondProgramToDiamondButton(body)
	)

	// UserProgramBSON ...
	result := models.UserProgramBSON{
		CompanyID: companyID,
		Silver:    silverButton,
		Golden:    goldenButton,
		Diamond:   diamondButton,
	}
	return result
}

func silverProgramToSilverButton(body models.UserProgramCreatePayload) models.SilverButton {
	result := models.SilverButton{
		Spending:   body.SilverSpending,
		Commission: body.SilverCommission,
	}
	return result
}

func goldenProgramToGoldenButton(body models.UserProgramCreatePayload) models.GoldenButton {
	result := models.GoldenButton{
		Spending:   body.GoldenSpending,
		Commission: body.GoldenCommission,
	}
	return result
}
func diamondProgramToDiamondButton(body models.UserProgramCreatePayload) models.DiamondButton {
	result := models.DiamondButton{
		Spending:   body.DiamondSpending,
		Commission: body.DiamondCommission,
	}
	return result
}

// silverValidate
func silverValidate(silver models.SilverButton) bool {
	if silver.Spending <= 1000 || silver.Commission <= 1 {
		return false
	}
	return true
}

func goldenValidate(silverButton models.SilverButton, goldenButton models.GoldenButton) bool {
	if goldenButton.Spending <= silverButton.Spending || goldenButton.Commission <= silverButton.Commission {
		return false
	}
	return true
}

func diamondValidate(goldenButton models.GoldenButton, diamondButton models.DiamondButton) bool {
	if diamondButton.Spending <= goldenButton.Spending || diamondButton.Commission <= goldenButton.Commission {
		return false
	}
	return true
}
