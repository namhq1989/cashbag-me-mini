package services

import (
	"cashbag-me-mini/models"
)

var (
	silverSpendingLevel    = 10000
	silverCommissionLevel  = 5
	goldenSpendingLevel    = 15000
	goldenCommissionLevel  = 10
	diamondSpendingLevel   = 20000
	diamondCommissionLevel = 15
)

// UserCreatePayloadToBSON ...
func userProgramCreatePayloadToBSON(body models.UserProgramCreatePayload) models.UserProgramBSON {
	result := models.UserProgramBSON{
		CompanyID: body.CompanyID,
		Silver:    body.SilverSpending,
		Silver:    body.SilverCommission,
		Golden:    body.GoldenSpending,
		Golden:    body.GoldenCommission,
		Diamond:   body.DiamondSpending,
		Diamond:   body.DiamondCommission,
	}
	return result
}

// silverValidate ...
func silverValidate() bool {
	var (
		button models.UserProgramCreatePayload
	)
	if button.SilverSpending == silverSpendingLevel {
		return true
	}
	if button.SilverCommission == silverCommissionLevel {
		return true
	}

}

// goldenValidate ...
func goldenValidate() bool {
	var (
		button models.UserProgramCreatePayload
	)
	if button.GoldenSpending == goldenSpendingLevel {
		return true
	}
	if button.GoldenCommission == goldenSpendingLevel {
		return true
	}

}

// diamondValidate ...
func diamondValidate() bool {
	var (
		button models.UserProgramCreatePayload
	)
	if button.DiamondSpending == diamondSpendingLevel {
		return true
	}
	if button.DiamondCommission == diamondSpendingLevel {
		return true
	}

}
