package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// BranchBSON ...
	BranchBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		CompanyID primitive.ObjectID `bson:"companyID"`
		Name      string             `bson:"name"`
		Address   string             `bson:"address"`
		Active    bool               `bson:"active"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	// BranchDetail ...
	BranchDetail struct {
		ID        primitive.ObjectID `json:"_id"`
		Company   CompanyBrief       `json:"company"`
		Name      string             `json:"name"`
		Address   string             `json:"address"`
		Active    bool               `json:"active"`
		CreatedAt time.Time          `json:"createdAt"`
		UpdatedAt time.Time          `json:"updatedAt"`
	}

	// BranchCreatePayload ...
	BranchCreatePayload struct {
		CompanyID string `json:"companyID" valid:"alphanum"`
		Name      string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address   string `json:"address" valid:"stringlength(3|100),type(string)"`
		Active    bool   `json:"active"`
	}

	// BranchUpdatePayload ...
	BranchUpdatePayload struct {
		Name    string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address string `json:"address" valid:"stringlength(3|100),type(string)"`
		Active  bool   `json:"active"`
	}

	// BranchBrief ...
	BranchBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}
)
