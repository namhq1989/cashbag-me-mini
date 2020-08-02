package services

import (
	
	"cashbag-me-mini/dao"
	"cashbag-me-mini/models"
)

//ListBranch ...
func ListBranch() []models.BranchDetail {
	var (
		result []models.BranchDetail
		nameCompany string
	)
	
	Branchs := dao.ListBranch()
	for _,item:= range Branchs {
			branch := ConvertToBranchDetail(item)
			nameCompany=dao.GetNameCompanyById(item.CompanyId)
			branch.CompanyId.ID=item.CompanyId
			branch.CompanyId.Name=nameCompany
			result = append(result, branch)
		}
	return result
}
//ConvertToBranchDetail ...
func ConvertToBranchDetail(x models.BranchBSON) models.BranchDetail {
	result := models.BranchDetail{
		ID:               x.ID,
		Name:			 x.Name,
		Address:		x.Address ,
		Active: 		x.Active 	,
		CreateAt:		 x.CreateAt,
		UpdateAt:		 x.UpdateAt,
	}
	return result
}
