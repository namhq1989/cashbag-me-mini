package routes

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/dao"
	"cashbag-me-mini/util"
	grpcuser "cashbag-me-mini/grpc/user"
)

func companyCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			companyID =c.Get("companyID").(primitive.ObjectID)
		)
		
		company, _   := dao.CompanyFindByID(companyID)

		// check existed 
		if company.ID.IsZero() {
			return util.Response404(c, nil, "Khong tim thay company")
		}

		c.Set("company",company)
		return next(c)
	}
}

func branchCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			branchID = c.Get("branchID").(primitive.ObjectID)
		)

		branch, _ := dao.BranchFindByID(branchID)

		// check existed
		if branch.ID.IsZero() {
			return util.Response404(c, nil, "Khong tim thay branch")
		}

		c.Set("branch", branch)
		return next(c)
	}
}

func userCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			userID = c.Get("userID").(primitive.ObjectID)
			userString=userID.Hex()
		)

		user, err := grpcuser.GetUserBriefByID(userString)
		if err != nil{
			return util.Response404(c, nil, "Khong tim thay user")
		}

		c.Set("user", user)
		return next(c)
	}
	
}