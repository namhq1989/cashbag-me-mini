package services

import (
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//ListBranch ...
func ListBranch() []models.BranchDetail {
	var (
		result      []models.BranchDetail
		nameCompany string
	)

	Branches := dao.ListBranch()
	for _, item := range Branches {
		branch := ConvertToBranchDetail(item)
		nameCompany = dao.GetNameCompanyById(item.CompanyId)
		branch.CompanyId.ID = item.CompanyId
		branch.CompanyId.Name = nameCompany
		result = append(result, branch)
	}
	return result
}

//ConvertToBranchDetail ...
func ConvertToBranchDetail(x models.BranchBSON) models.BranchDetail {
	result := models.BranchDetail{
		ID:       x.ID,
		Name:     x.Name,
		Address:  x.Address,
		Active:   x.Active,
		CreateAt: x.CreateAt,
		UpdateAt: x.UpdateAt,
	}
	return result
}

//CreateBranch ...
func CreateBranch(body models.PostBranch) *mongo.InsertOneResult {
	var (
		branch    models.BranchBSON
		companyId primitive.ObjectID
	)
	companyId = dao.GetIdCompanyByName(body.NameCompany)
	branch = ConvertBodyToBranchBSON(body)
	branch.CompanyId = companyId
	branch.ID = primitive.NewObjectID()
	branch.CreateAt = time.Now()
	result := dao.CreateBranch(branch)
	return result
}

// ConvertBodyToBranchBSON...
func ConvertBodyToBranchBSON(body models.PostBranch) models.BranchBSON {
	result := models.BranchBSON{
		Name:    body.Name,
		Address: body.Address,
		Active:  body.Active,
	}
	return result
}

// PatchBranch ...
func PatchBranch(idBranch interface{}) *mongo.UpdateResult {
	result := dao.PatchBranch(idBranch)
	return result
}

// PutBranch ...
func PutBranch(idBranch interface{}, body models.PutBranch) *mongo.UpdateResult {
	result := dao.PutBranch(idBranch, body)
	return result
}
