package routes

import (
	"github.com/carfesal/go-game-management-system/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/games", controllers.GetGames).Methods("GET")
	router.HandleFunc("/api/games/{id}", controllers.GetGame).Methods("GET")
	router.HandleFunc("/api/games", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/api/games/{id}", controllers.UpdateGame).Methods("PUT")
	router.HandleFunc("/api/games/{id}", controllers.DeleteGame).Methods("DELETE")
}
