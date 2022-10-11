package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	db.AutoMigrate(&Task{})
}
