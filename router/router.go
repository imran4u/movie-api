package router

import (
	"github.com/gorilla/mux"
	"github.com/imran4u/movie-api/router/handler"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", handler.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", handler.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", handler.MarketWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", handler.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/delete_movies", handler.DeleteAllMovies).Methods("DELETE")

	return router
}
