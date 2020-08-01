package dao

import (
	"cashbag-me-mini/modules/database"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//GetCompanyById ....
func GetNameCompanyById(id interface{}) string {
	var (
		db                = database.Connectdb("CashBag")
		companyCollection = db.Collection("companies")
		ctx    = context.Background()
		result=struct{
			Name string             `bson:"name"`
		}{}
		filter = bson.M{"_id": id}
	)
	err := companyCollection.FindOne(ctx,filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Name
}