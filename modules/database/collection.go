package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// ConnectCol ...
func ConnectCol(nameCol string) *mongo.Collection {
	return db.Collection(nameCol)
}
