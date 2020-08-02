package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//CompanyList ...
func CompanyList() []models.CompanyBSON {
	var (
		companyCollection = database.ConnectCol("companies")
		ctx               = context.Background()
		result            []models.CompanyBSON
	)
	cursor, err := companyCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &result)
	return result
}
