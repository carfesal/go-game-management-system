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
