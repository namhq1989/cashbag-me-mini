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
		Name:      body.Name,
		Address:   body.Address,
	}
	return result
}
