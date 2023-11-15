package main

import (
	"fmt"
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
	log.Fatal(http.ListenAndServe("localhost:9010", router))
	fmt.Println("Server running on port 9010")
}
