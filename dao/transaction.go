package dao

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TransactionCreate ....
func TransactionCreate(doc models.TransactionBSON, balance float64) (models.TransactionBSON, error) {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)

	//Add update information
	if doc.ID.IsZero() {
		doc.ID = primitive.NewObjectID()
	}
	doc.CreatedAt = time.Now()
	balanceCurrent := balance - doc.Commission
	_, err := collection.InsertOne(ctx, doc)

	UpdateBalance(doc.ID, balanceCurrent)
	HandleTranAnalytic(doc)

	return doc, err
}
