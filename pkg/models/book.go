package models

import (
	"github.com/carfesal/go-game-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Game struct {
	gorm.Model
	Name        string `json:"name"`
	Developer   string `json:"developer"`
	Platform    string `json:"platform"`
	ReleaseDate string `json:"release_date"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB() // returns a *gorm.DB object
	db.AutoMigrate(&Game{})
}

func (g *Game) CreateGame() *Game {
	db.Create(&g)
	return g
}

func GetAllGames() []Game {
	var games []Game
	db.Find(&games)
	return games
}

func GetGameById(Id int64) (*Game, *gorm.DB) {
	var game Game
	db := db.Where("ID = ?", Id).Find(&game) // find game with id
	return &game, db
}

func DeleteGameById(Id int64) Game {
	var game Game
	db.Where("ID = ?", Id).Delete(game)
	return game
}
