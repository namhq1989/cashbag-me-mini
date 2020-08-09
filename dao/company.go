package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// CompanyCreate ....
func CompanyCreate(doc models.CompanyBSON) (models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)

	//Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.CreatedAt = time.Now()

	_, err := companyCol.InsertOne(ctx, doc)

	return doc, err
}

// CompanyList  ...
func CompanyList() ([]models.CompanyBSON, error) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		doc        = make([]models.CompanyBSON, 0)
	)

	cursor, err := companyCol.Find(ctx, bson.M{})

	//Close cursor
	defer cursor.Close(ctx)

	//Set result
	cursor.All(ctx, &doc)

	return doc, err
}


// CompanyUpdate  to ...
func CompanyUpdate(id primitive.ObjectID, company models.CompanyBSON) (models.CompanyBSON, error) {
	var (
		filter = bson.M{"_id": id}
		update = bson.M{"$set": bson.M{
			"name":          company.Name,
			"address":        company.Address,
			"active":         company.Active,
			"balance":        company.Balance,
			"loyaltyProgram": company.LoyaltyProgram,
			"updatedAt":       time.Now(),
		}}
	)

	// Update
	err := CompanyUpdateByID(filter, update)

	// Get doc
	doc , _ := CompanyFindbyID(id)
	

	return doc, err
}

// CompanyChangeActiveStatus .... ...
func CompanyChangeActiveStatus(id primitive.ObjectID) (models.CompanyBSON, error) {
	var (
		filter     = bson.M{"_id": id}
		doc,_ 	= 	CompanyFindbyID(id)	
		update     = bson.M{"$set": bson.M{"active": !doc.Active}}
	)

	err := CompanyUpdateByID(filter, update)

	return doc, err
}
// CompanyFindbyID ...
func CompanyFindbyID(id primitive.ObjectID) (models.CompanyBSON,error){
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result   models.CompanyBSON
		filter = bson.M{"_id": id}
	)

	err := companyCol.FindOne(ctx, filter).Decode(&result)

	return result,err
}

//  GetIFCompanyByName..
func GetIFCompanyByName(NameCompany interface{}) models.IFCompany {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = models.IFCompany{}
		filter     = bson.M{"name": NameCompany}
	)
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//UpdateBalance ...
func UpdateBalance(id primitive.ObjectID, balance float64) {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
	)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"balance": balance,
	}}
	log.Println(balance)
	_, err := companyCol.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// CompanyUpdateByID ...
func CompanyUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		CompanyCol = database.CompanyCol()
		ctx       = context.Background()
	)

	_, err := CompanyCol.UpdateOne(ctx, filter, updateData)

	return err
}

// CompanyValidateID ...
func CompanyValidateID(companyID primitive.ObjectID) bool {
	var (
		companyCol = database.CompanyCol()
		ctx       = context.Background()
		filter    = bson.M{"_id": companyID}
	)

	err := companyCol.FindOne(ctx, filter)
	if err != nil{
		return false
	}
	return true
}


//GetNameCompanyById ....
func GetNameCompanyById(id interface{}) string {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = struct {
			Name string `bson:"name"`
		}{}
		filter = bson.M{"_id": id}
	)
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Name
}

// GetIdCompanyByName ...
func GetIdCompanyByName(nameCompany interface{}) primitive.ObjectID {
	var (
		companyCol = database.CompanyCol()
		ctx        = context.Background()
		result     = struct {
			ID primitive.ObjectID `bson:"_id"`
		}{}
		filter = bson.M{"name": nameCompany}
	)
	err := companyCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.ID
}