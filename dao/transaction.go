package dao

import (
	"context"

	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/database"
)

// TransactionCreate ...
func TransactionCreate(doc models.TransactionBSON) (models.TransactionBSON, error) {
	var (
		collection = database.TransactionCol()
		ctx        = context.Background()
	)

	// Insert
	_, err := collection.InsertOne(ctx, doc)

	return doc, err
}
