package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carfesal/go-game-management-system/pkg/models"
	"github.com/carfesal/go-game-management-system/pkg/utils"
)

var NewGame models.Game

func GetGames(w http.ResponseWriter, r *http.Request) {
	newGames := models.GetAllGames()
	res, _ := json.Marshal(newGames)
	utils.RespondJson(w, res, http.StatusOK)
}

func GetGame(w http.ResponseWriter, r *http.Request) {

}

func CreateGame(w http.ResponseWriter, r *http.Request) {

}

func UpdateGame(w http.ResponseWriter, r *http.Request) {

}

func DeleteGame(w http.ResponseWriter, r *http.Request) {

}
