package models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type (
	CompanyBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		Name           string             `bson:"name"`
		Address        string             `bson:"address"`
		Balance        float64            `bson:"balance" `
		LoyaltyProgram float64            `bson:"loyaltyProgram" `
		Active         bool               `bson:"active" `
		CreateAt       time.Time          `bson:"createAt"`
		UpdateAt       time.Time          `bson:"updateAt"`
	}
	CompanyDetail struct {
		ID             primitive.ObjectID `json:"_id"`
		Name           string             `json:"name"`
		Address        string             `json:"address"`
		Balance        float64            `json:"balance"`
		LoyaltyProgram float64            `json:"loyaltyProgram"`
		Active         bool               `json:"active"`
		CreateAt       time.Time          `json:"createAt"`
		UpdateAt       time.Time          `json:"updateAt"`
	}
	PostCompany struct {
		Name        string `json:"name" valid:"stringlength(3|30),type(string)"`
		Address     string `json:"address" valid:"stringlength(3|100),type(string)"`
		Active      bool   `json:"active"`
	}
)
