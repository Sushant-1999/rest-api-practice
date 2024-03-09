package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Roll is model for sushi
type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

// Init rolls var as a slice
// slice will act as a Database
var rolls []Roll

func main() {
	fmt.Println("Rest API")
	//Defining our routes using Mux Router

	//initialization of Mux Router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
