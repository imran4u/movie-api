package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// should be in config file
const connectionString = "mongodb+srv://erimran4u:imran123@movieclustor0.eumdv.mongodb.net/?retryWrites=true&w=majority&appName=movieclustor0"
const dbName = "netflix"
const colName = "watchlist"

// collection
var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mondo db
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo db connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection Ref is ready")
}
