package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dipankarupd/betting_raja/controller"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	RegisterGameRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting the server:")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

var RegisterGameRoutes = func(router *mux.Router) {
	router.HandleFunc("/games/", controller.PredictWinner).Methods("POST")
	router.HandleFunc("/games/", controller.GetGames).Methods("GET")
}
