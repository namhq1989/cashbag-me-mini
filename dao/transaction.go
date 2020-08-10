package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// TransactionCreate ....
func TransactionCreate(doc models.TransactionBSON, balance float64) (models.TransactionBSON, error) {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)

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

	return doc, err
}
