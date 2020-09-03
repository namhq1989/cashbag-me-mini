package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cashbag-me-mini/config"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

// Connect ...
func Connect() {
	envVars := config.GetEnv()

	// Connect
	cl, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", envVars.Database.URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = cl.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	client=cl
	db = cl.Database(envVars.Database.TransactionName)
	fmt.Println("Database Connected to", envVars.Database.TransactionName)
}

// GetClient ...
func GetClient() *mongo.Client {
	return client
}

// SetDB ...
func SetDB(dbValue *mongo.Database) {
	db = dbValue
}

// GetDB ...
func GetDB() *mongo.Database {
	return db
}
