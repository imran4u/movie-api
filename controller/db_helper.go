package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/imran4u/movie-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneMove(movie model.Netflix) model.Netflix {
	insert, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("One movie instered with id", insert.InsertedID)
	movie.ID = insert.InsertedID.(primitive.ObjectID)
	return movie
}

func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated one movie, ModifiedCount  =", result.ModifiedCount)

}

// delte
func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one movie, ModifiedCount  =", result.DeletedCount)

}

// delte all
func DeleteAllMovie() int64 {

	filter := bson.D{{}}

	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all movies, ModifiedCount  =", result.DeletedCount)
	return result.DeletedCount
}

// Get all movies, bit tricky

func GetAllMovies() []primitive.M {
	fmt.Println("getAllMovies called")
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
			continue
		}
		movies = append(movies, movie)
	}
	fmt.Println("Found movies", movies)
	return movies
}
