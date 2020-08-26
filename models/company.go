package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// CompanyBSON ...
	CompanyBSON struct {
		ID              primitive.ObjectID `bson:"_id"`
		Name            string             `bson:"name"`
		Address         string             `bson:"address"`
		Balance         float64            `bson:"balance" `
		CashbackPercent float64            `bson:"cashbackPercent" `
		Active          bool               `bson:"active" `
		PaidType        string             `bson:"paidType"`
		CreatedAt       time.Time          `bson:"createdAt"`
		UpdatedAt       time.Time          `bson:"updatedAt"`
	}

	// CompanyDetail ...
	CompanyDetail struct {
		ID              primitive.ObjectID `json:"_id"`
		Name            string             `json:"name"`
		Address         string             `json:"address"`
		Balance         float64            `json:"balance"`
		CashbackPercent float64            `json:"cashbackPercent"`
		Active          bool               `json:"active"`
		PaidType        string             `json:"paidType"`
		CreatedAt       time.Time          `json:"createdAt"`
		UpdatedAt       time.Time          `json:"updatedAt"`
	}

	// CompanyCreatePayload ...
	CompanyCreatePayload struct {
		Name    string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address string `json:"address" `
	}

	// CompanyUpdatePayload ...
	CompanyUpdatePayload struct {
		Name            string  `json:"name" valid:"stringlength(3|30),type(string)"`
		Address         string  `json:"address" valid:"stringlength(3|30),type(string)"`
		Balance         float64 `json:"balance" valid:"required"`
		CashbackPercent float64 `json:"cashbackPercent" valid:"required"`
		Active          bool    `json:"active"`
		PaidType        string  `json:"paidType" valid:"stringlength(3|30),type(string)"`
	}

	// CompanyBrief ...
	CompanyBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}
)
