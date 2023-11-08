package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("games.db"), &gorm.Config{})

	if err != nil {
		panic("Oops! Something went wrong while connecting to the database!")
	}

	db = database
}

func GetDB() *gorm.DB {
	return db
}
