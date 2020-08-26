package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func convertToCompanyDetail(doc models.CompanyBSON) models.CompanyDetail {
	result := models.CompanyDetail{
		ID:              doc.ID,
		Name:            doc.Name,
		Address:         doc.Address,
		Balance:         doc.Balance,
		CashbackPercent: doc.CashbackPercent,
		Active:          doc.Active,
		PaidType:        doc.PaidType,
		CreatedAt:       doc.CreatedAt,
		UpdatedAt:       doc.UpdatedAt,
	}
	return result
}

func companyCreatePayloadToBSON(body models.CompanyCreatePayload) models.CompanyBSON {
	result := models.CompanyBSON{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		Address:   body.Address,
		CreatedAt: time.Now(),
	}

	return result
}

func companyCreateCompanyAnalytic(companyID primitive.ObjectID) (err error) {
	var (
		silver       = "silver"
		gold         = "gold"
		diamond      = "diamond"
		silverMember = models.CompanyAnalyticMember{
			ID: silver,
		}
		goldMember = models.CompanyAnalyticMember{
			ID: gold,
		}
		diamondMember = models.CompanyAnalyticMember{
			ID: diamond,
		}
		members []models.CompanyAnalyticMember
	)

	members = append(members, silverMember, goldMember, diamondMember)

	companyAnalytic := models.CompanyAnalyticBSON{
		ID:        primitive.NewObjectID(),
		CompanyID: companyID,
		Members:   members,
		UpdatedAt: time.Now(),
	}

	// Create CompanyAnalytic
	err = dao.CompanyAnalyticCreate(companyAnalytic)
	if err != nil {
		err = errors.New("Khong The Tao Company Analytic")
		return
	}
	return
}

func companyUpdateBalance(companyID primitive.ObjectID, balance float64) error {
	var (
		filter = bson.M{"_id": companyID}
		update = bson.M{"$set": bson.M{
			"balance": balance,
		}}
	)

	// update
	err := dao.CompanyUpdateByID(filter, update)

	return err
}
