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

	fmt.Println("Starting the server:")
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, r))
}

var RegisterGameRoutes = func(router *mux.Router) {
	router.HandleFunc("/games/", controller.PredictWinner).Methods("POST")
	router.HandleFunc("/games/", controller.GetGames).Methods("GET")
}
