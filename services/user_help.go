package services

import (
	"cashbag-me-mini/models"
)

// UserCreatePayloadToBSON ...
func userCreatePayloadToBSON(body models.UserCreatePayload) models.UserBSON {
	result := models.UserBSON{
		Name:     body.Name,
		Address:  body.Address,
		Spending: body.Spending,
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
