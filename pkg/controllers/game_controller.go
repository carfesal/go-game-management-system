package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/carfesal/go-game-management-system/pkg/models"
	"github.com/carfesal/go-game-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

var NewGame models.Game

func GetGames(w http.ResponseWriter, r *http.Request) {
	newGames := models.GetAllGames()
	res, _ := json.Marshal(newGames)
	utils.ResponseJson(w, res, http.StatusOK)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["id"]
	ID, err := strconv.ParseInt(gameId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing the ID")
		utils.ResponseJson(w, []byte("Invalid ID"), http.StatusNotFound)
		return
	}

	game, _ := models.GetGameById(ID)
	res, _ := json.Marshal(game)
	utils.ResponseJson(w, res, http.StatusOK)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	createGame := &models.Game{}
	utils.ParseBody(r, createGame)
	b := createGame.CreateGame()
	res, _ := json.Marshal(b)
	utils.ResponseJson(w, res, http.StatusCreated)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["id"]
	ID, err := strconv.ParseInt(gameId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing the ID")
		utils.ResponseJson(w, []byte("Invalid ID"), http.StatusNotFound)
		return
	}
	var updateGame = &models.Game{}
	utils.ParseBody(r, updateGame)
	game, db := models.GetGameById(ID)

	if updateGame.Name != "" {
		game.Name = updateGame.Name
	}

	if updateGame.Developer != "" {
		game.Developer = updateGame.Developer
	}

	if updateGame.Platform != "" {
		game.Platform = updateGame.Platform
	}

	if updateGame.ReleaseDate != "" {
		game.ReleaseDate = updateGame.ReleaseDate
	}

	db.Save(&game)
	res, _ := json.Marshal(game)
	utils.ResponseJson(w, res, http.StatusOK)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["id"]
	ID, err := strconv.ParseInt(gameId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing the ID")
		utils.ResponseJson(w, []byte("Invalid ID"), http.StatusNotFound)
		return
	}

	deleteGame := models.DeleteGameById(ID)
	res, _ := json.Marshal(deleteGame)
	json.Marshal(res)
	utils.ResponseJson(w, res, http.StatusOK)
}
