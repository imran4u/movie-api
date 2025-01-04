package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/imran4u/movie-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection
var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(config.ConnectionString)

	// connect to mondo db
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo db connection success")

	collection = client.Database(config.DbName).Collection(config.ColName)
	fmt.Println("collection Ref is ready")
}
