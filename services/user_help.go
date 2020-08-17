package services

import (
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
		Name:     body.Name,
		Address:  body.Address,
	}
	return result
}

// UserUpdatePayloadToBSON ...
func userUpdatePayloadToBSON(body models.UserUpdatePayload) models.UserBSON {
	result := models.UserBSON{
		Name:     body.Name,
		Address:  body.Address,
		Level:    body.Level,
		Spending: body.Spending,
	}
	return result
}



