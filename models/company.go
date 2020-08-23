package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// CompanyBSON ...
	CompanyBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		Name           string             `bson:"name"`
		Address        string             `bson:"address"`
		Balance        float64            `bson:"balance" `
		LoyaltyProgram float64            `bson:"loyaltyProgram" `
		Active         bool               `bson:"active" `
		Postpaid       string               `bson:"postpaid"`
		CreatedAt      time.Time          `bson:"createdAt"`
		UpdatedAt      time.Time          `bson:"updatedAt"`
	}

	// CompanyDetail ...
	CompanyDetail struct {
		ID             primitive.ObjectID `json:"_id"`
		Name           string             `json:"name"`
		Address        string             `json:"address"`
		Balance        float64            `json:"balance"`
		LoyaltyProgram float64            `json:"loyaltyProgram"`
		Active         bool               `json:"active"`
		Postpaid       string               `json:"postpaid"`
		CreatedAt      time.Time          `json:"createdAt"`
		UpdatedAt      time.Time          `json:"updatedAt"`
	}

	// CompanyCreatePayload ...
	CompanyCreatePayload struct {
		Name    string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address string `json:"address" `
	}

	// CompanyUpdatePayload ...
	CompanyUpdatePayload struct {
		Name           string  `json:"name" valid:"stringlength(3|30),type(string)"`
		Address        string  `json:"address" `
		Balance        float64 `json:"balance" valid:"required"`
		LoyaltyProgram float64 `json:"loyaltyProgram" valid:"required"`
		Active         bool    `json:"active"`
		Postpaid       string    `json:"postpaid"`
	}

	// CompanyBrief ...
	CompanyBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}
)
