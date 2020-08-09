package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type (

	// BranchBSON ...
	BranchBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		CompanyID primitive.ObjectID `bson:"companyId"`
		Name      string             `bson:"name"`
		Address   string             `bson:"address"`
		Active    bool               `bson:"active"`
		CreatedAt  time.Time          `bson:"createAt"`
		UpdatedAt  time.Time          `bson:"updateAt"`
	}

	// BranchDetail ...
	BranchDetail struct {
		ID       primitive.ObjectID `json:"_id"`
		Company  CompanyBrief       `json:"companyId"`
		Name     string             `json:"name"`
		Address  string             `json:"address"`
		Active   bool               `json:"active"`
		CreatedAt time.Time          `json:"createAt"`
		UpdatedAt time.Time          `json:"updateAt"`
	}

	// CompanyBrief ...
	CompanyBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}

	// CreateBranchPayload ...
	BranchCreatePayload struct {
		CompanyID primitive.ObjectID `json:"CompanyID" valid:"alphanum"`
		Name      string             `json:"name" valid:"stringlength(3|30),type(string)"`
		Address   string             `json:"address" valid:"stringlength(3|100),type(string)"`
		Active    bool               `json:"active"`
	}

	// UpdateBranchPayload ...
	BranchUpdateBPayload struct {
		Name    string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address string `json:"address" valid:"stringlength(3|100),type(string)"`
		Active  bool   `json:"active"`
	}

	// BranchBrief
	BranchBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}
)
