package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/imran4u/movie-api/router"
)

func main() {

	r := router.Router()

	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Server has started and listening on port 5000 ....")
}
