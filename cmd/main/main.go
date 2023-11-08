package main

import (
	"log"
	"net/http"

	"github.com/carfesal/go-game-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)

	//Register routes in run
	routes.RegisterRoutes(router)

	//Create the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
