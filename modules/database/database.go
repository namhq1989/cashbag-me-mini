package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"cashbag-me-mini/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *mongo.Database
)

//Connect ...
func Connect(dbName string) *mongo.Database {
	cfg := config.GetEnv()
	//connect
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.DatabaseURI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database ", cfg.DatabaseURI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	DB = client.Database(dbName)
	fmt.Println("Connected to db:", dbName)
	return DB
}
