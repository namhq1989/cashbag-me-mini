package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

//Connectdb ...
func Connectdb(dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://CashbagMe:Cashbag@cluster0.epe8y.gcp.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	db = client.Database(dbName)
	fmt.Println("Connected to db:", dbName)
	return db
}
