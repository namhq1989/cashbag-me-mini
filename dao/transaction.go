package dao

import (
	"context"
	"time"
<<<<<<< HEAD

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
=======

	"go.mongodb.org/mongo-driver/bson/primitive"
>>>>>>> 35bc62c3407eecf25d9e630ee81dc51e89f7e4bb
)

// TransactionCreate ....
func TransactionCreate(doc models.TransactionBSON, balance float64) (models.TransactionBSON, error) {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)

<<<<<<< HEAD
	// Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.CreatedAt = time.Now()

	// Insert
	_, err := collection.InsertOne(ctx, doc)

	if err == nil {
		balanceCurrent := balance - doc.Commission
		CompanyUpdateBalance(doc.CompanyID, balanceCurrent)
		TransactionAnalyticHandle(doc)
	}
=======
	//Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.CreatedAt = time.Now()
	balanceCurrent := balance - doc.Commission
	_, err := collection.InsertOne(ctx, doc)

	UpdateBalance(doc.ID, balanceCurrent)
	HandleTranAnalytic(doc)
>>>>>>> 35bc62c3407eecf25d9e630ee81dc51e89f7e4bb

	return doc, err
}
