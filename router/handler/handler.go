package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imran4u/movie-api/controller"
	"github.com/imran4u/movie-api/model"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GEt All movies called")
	w.Header().Set("Content-type", "application/json")
	movies := controller.GetAllMovies()
	json.NewEncoder(w).Encode(movies)
}

// func GetAMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	params := mux.Vars(r)

// 	json.NewEncoder(w).Encode(movies)
// }

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "POST") // it will allow only post method
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	movie = controller.InsertOneMove(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarketWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "PUT") // it will allow only put method
	params := mux.Vars(r)
	controller.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(fmt.Sprintf("updated watched for id = %s", params["id"]))
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "DELETE") // it will allow only DElete method
	params := mux.Vars(r)
	controller.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(fmt.Sprintf("Deleted the movie with id = %s", params["id"]))
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("allow-control-allow-methods", "DELETE") // it will allow only DElete method
	count := controller.DeleteAllMovie()
	json.NewEncoder(w).Encode(fmt.Sprintf("Deleted the total movies, count is = %d", count))
}
