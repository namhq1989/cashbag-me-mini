package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// BranchBSON ...
	BranchBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		CompanyId        primitive.ObjectID `bson:"companyId"`
		Name			 string 			`bson:"name"`
		Address			 string 			`bson:"address"`
		Active			 bool 				`bson:"active"`
		CreateAt		 time.Time			`bson:"createAt"`
		UpdateAt		 time.Time			`bson:"updateAt"`
	}

	// CompanyBrief
	CompanyBrief struct {
		ID   primitive.ObjectID `json:"_id"`
		Name string             `json:"name"`
	}

	//BranchDetail ...
	BranchDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		CompanyId        CompanyBrief       `json:"companyId"`
		Name			 string 			`json:"name"`
		Address			 string 			`json:"address"`
		Active			 bool 				`json:"active"`
		CreateAt		 time.Time			`json:"createAt"`
		UpdateAt		 time.Time			`json:"updateAt"`
	}
)